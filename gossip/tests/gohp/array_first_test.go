package gohp_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

// Separate array_first comparison on the integrated empty-capable compiler.
func TestOriginalWordPressArrayFirst(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	manifest := read(t, "array-first-unit.json")
	var pin struct {
		unitPin
		Sources []struct {
			Path      string     `json:"source_path"`
			Hash      string     `json:"source_sha256"`
			Fragments []fragment `json:"fragments"`
		} `json:"sources"`
	}
	if err := json.Unmarshal(manifest, &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.array-first-unit/v1" || len(pin.Sources) != 1 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid array-first pin")
	}
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil {
		t.Fatal(err)
	}
	if baseline.Commit != pin.WordPressCommit {
		t.Fatal("boundary and WordPress baseline disagree")
	}
	type selection struct {
		path        string
		first, last int
	}
	wanted := map[string]selection{
		"array_first":       {"wp-includes/compat.php", 657, 665},
		"array_first_guard": {"wp-includes/compat.php", 646, 666},
	}
	fragments := map[string][]byte{}
	var selections []map[string]any
	for _, source := range pin.Sources {
		if source.Path != "wp-includes/compat.php" {
			t.Fatal("unexpected original source")
		}
		original := read(t, filepath.Join("../../..", source.Path))
		if digest(original) != source.Hash {
			t.Fatal("WordPress source drift; review original extraction")
		}
		lines := bytes.SplitAfter(original, []byte("\n"))
		for _, part := range source.Fragments {
			if expected, ok := wanted[part.Name]; !ok || expected != (selection{source.Path, part.First, part.Last}) || fragments[part.Name] != nil || part.Last > len(lines) {
				t.Fatal("unexpected/duplicate original fragment")
			}
			fragments[part.Name] = bytes.Join(lines[part.First-1:part.Last], nil)
			selections = append(selections, map[string]any{"name": part.Name, "source_path": source.Path, "source_sha256": source.Hash,
				"first_line": part.First, "last_line": part.Last, "sha256": digest(fragments[part.Name])})
		}
	}
	if len(fragments) != len(wanted) {
		t.Fatal("missing original array_first fragment")
	}
	definitions, nativeDefinitions := []byte("<?php\n"), []byte("<?php\n")
	for _, name := range []string{"array_first"} {
		if !bytes.Contains(fragments[name+"_guard"], fragments[name]) {
			t.Fatal("incomplete original guard")
		}
		definitions = append(definitions, fragments[name]...)
		nativeDefinitions = append(nativeDefinitions, fragments[name+"_guard"]...)
	}
	casesBytes := read(t, "array-first-cases.json")
	expected := arrayFirstExpected(t, casesBytes)
	expectedBytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}
	workspace := t.TempDir()
	source, input, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "input"), filepath.Join(workspace, "candidate")
	for _, dir := range []string{source, input, candidate} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}
	archiveHash := exportGoHP(t, repository, pin.GoHPCommit, source)
	// Reuse only bounded transport functions, never old observers or generated calls.
	goTransportOriginal := read(t, "numeric-array-observer.go.txt")
	goBoundary := []byte("func numericArrayPredicateObservation(")
	if bytes.Count(goTransportOriginal, goBoundary) != 1 {
		t.Fatal("Go transport boundary drift")
	}
	goTransport := goTransportOriginal[:bytes.Index(goTransportOriginal, goBoundary)]
	phpTransportOriginal := read(t, "numeric-array-observer.php")
	phpStart, phpEnd := []byte("function numeric_array_fixture_type("), []byte("if (! is_array($cases) || ! array_is_list($cases)) {")
	if bytes.Count(phpTransportOriginal, phpStart) != 1 || bytes.Count(phpTransportOriginal, phpEnd) != 1 {
		t.Fatal("PHP transport boundary drift")
	}
	phpTransport := append([]byte("<?php\n"), phpTransportOriginal[bytes.Index(phpTransportOriginal, phpStart):bytes.Index(phpTransportOriginal, phpEnd)]...)
	phpObserver := read(t, "array-first-observer.php")
	goObserver := read(t, "array-first-observer.go.txt")
	write(t, filepath.Join(input, "helpers.php"), nativeDefinitions)
	write(t, filepath.Join(input, "unit.php"), definitions)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "transport.php"), phpTransport)
	write(t, filepath.Join(input, "observer.php"), phpObserver)
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	write(t, filepath.Join(candidate, "transport.go"), goTransport)
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456",
		"--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local",
		"--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "array_first_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/unit.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
	if len(buildErr) != 0 {
		t.Fatalf("compiler stderr: %s", buildErr)
	}
	generated, reportBytes := read(t, filepath.Join(candidate, "main.go")), read(t, filepath.Join(candidate, "report.json"))
	var report struct {
		Schema          string `json:"schema"`
		Status          string `json:"status"`
		Target          string `json:"php_target"`
		SourceHash      string `json:"source_sha256"`
		GeneratedHash   string `json:"generated_sha256"`
		CompilerVersion string `json:"compiler_version"`
	}
	if err := json.Unmarshal(reportBytes, &report); err != nil {
		t.Fatal(err)
	}
	if report.Schema != "gohp.build/v0.1" || report.Status != "generated" || report.Target != pin.PHPTarget ||
		report.SourceHash != digest(definitions) || report.GeneratedHash != digest(generated) {
		t.Fatal("report does not match exact source/output bytes")
	}
	t.Log("status=generated; original declarations unchanged; top-level loading guard is not compiled")
	_, compileErr := container(t, "array_first_build", append(buildArgs, "--entrypoint=go", pin.GoImage,
		"build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/transport.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("build stderr: %s", compileErr)
	}
	t.Log("status=compiled; no compatibility claim before observations")
	actual, actualErr := container(t, "array_first_candidate",
		"--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "array_first_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "manifest_sha256": digest(manifest),
		"wordpress_commit": pin.WordPressCommit, "fragments": selections, "original_declaration_count": 1, "owned_driver_count": 0,
		"compiled_bundle_sha256": digest(definitions), "native_bundle_sha256": digest(nativeDefinitions),
		"loading_guard": "native executes unchanged guard and verifies user-defined polyfill; compiler extracts unchanged declarations only",
		"cases_sha256":  digest(casesBytes), "case_count": len(expected),
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "php_image": pin.PHPImage, "php_version": phpVersion, "php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "transport_hashes": map[string]string{"go_source": digest(goTransportOriginal), "go_selected": digest(goTransport), "php_source": digest(phpTransportOriginal), "php_selected": digest(phpTransport)},
		"binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "one unchanged original array_first polyfill;12 retained cases; actual scalar return types/values and ordered caller inputs; flat scalar arrays; separate from22-original230-case integration; array_last remains unsupported",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}

func arrayFirstExpected(t *testing.T, data []byte) []map[string]any {
	t.Helper()
	// Carry the native-reference cases forward unchanged, including expectations.
	original := read(t, "array-edge-cases.json")
	var allCases, selected []map[string]any
	numericBundleDecode(t, original, &allCases)
	numericBundleDecode(t, data, &selected)
	allObservations := arrayEdgeExpected(t, original)
	var wantedCases, wantedObservations []map[string]any
	for i, c := range allCases {
		if c["operation"] == "array_first" {
			wantedCases = append(wantedCases, c)
			wantedObservations = append(wantedObservations, allObservations[i])
		}
	}
	if len(selected) != 12 || !reflect.DeepEqual(selected, wantedCases) {
		t.Fatal("array_first cases drifted from the twelve native-reference cases")
	}
	return wantedObservations
}
