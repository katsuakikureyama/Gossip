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

// This separate configuration constructs the input object and assigns supports
// inside translated PHP, before invoking the unchanged WordPress helpers.
func TestOriginalWordPressElementsSourceInput(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE; cached pinned Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run as a non-root user")
	}
	manifest := read(t, "elements-source-input.json")
	var pin struct {
		unitPin
		CasesHash string `json:"cases_sha256"`
		Sources   []struct {
			Path      string     `json:"source_path"`
			Hash      string     `json:"source_sha256"`
			Fragments []fragment `json:"fragments"`
		} `json:"sources"`
	}
	if err := json.Unmarshal(manifest, &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.elements-source-input/v1" || len(pin.Sources) != 3 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid serialization pin")
	}
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil {
		t.Fatal(err)
	}
	if baseline.Commit != pin.WordPressCommit {
		t.Fatal("baseline mismatch")
	}
	type selection struct {
		path        string
		first, last int
	}
	wanted := map[string]selection{
		"_wp_array_get": {"wp-includes/functions.php", 5233, 5272},
		"wp_should_skip_block_supports_serialization": {"wp-includes/block-supports/utils.php", 22, 35},
		"wp_should_add_elements_class_name":           {"wp-includes/block-supports/elements.php", 33, 97},
	}
	definitions := []byte("<?php\n")
	seen := map[string]bool{}
	var selections []map[string]any
	for _, source := range pin.Sources {
		if len(source.Fragments) != 1 {
			t.Fatal("expected one declaration per source")
		}
		part := source.Fragments[0]
		if want, ok := wanted[part.Name]; !ok || want != (selection{source.Path, part.First, part.Last}) || seen[part.Name] {
			t.Fatal("unexpected/duplicate original fragment")
		}
		seen[part.Name] = true
		original := read(t, filepath.Join("../../..", source.Path))
		if digest(original) != source.Hash {
			t.Fatal("WordPress source drift")
		}
		lines := bytes.SplitAfter(original, []byte("\n"))
		if part.Last > len(lines) {
			t.Fatal("invalid original range")
		}
		declaration := bytes.Join(lines[part.First-1:part.Last], nil)
		if !bytes.HasPrefix(declaration, []byte("function "+part.Name+"(")) || !bytes.HasSuffix(declaration, []byte("}\n")) {
			t.Fatal("incomplete original declaration")
		}
		definitions = append(definitions, declaration...)
		selections = append(selections, map[string]any{"name": part.Name, "source_path": source.Path, "source_sha256": source.Hash,
			"first_line": part.First, "last_line": part.Last, "sha256": digest(declaration)})
	}
	if len(seen) != len(wanted) {
		t.Fatal("missing declaration")
	}
	drivers := read(t, "elements-source-input-drivers.php.txt")
	definitions = append(definitions, drivers...)
	casesBytes := read(t, "elements-serialization-cases.json")
	if digest(casesBytes) != pin.CasesHash {
		t.Fatal("case corpus drift")
	}
	var cases []struct {
		Name       string `json:"name"`
		Mode       string `json:"mode"`
		Object     *bool  `json:"object"`
		Input      any    `json:"input"`
		FeatureSet any    `json:"feature_set"`
		Feature    any    `json:"feature"`
		Block      any    `json:"block"`
		Expected   any    `json:"expected"`
	}
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 24 {
		t.Fatal("expected24 serialization cases")
	}
	var expected []map[string]any
	seenCases, counts := map[string]bool{}, map[string]int{}
	for _, c := range cases {
		if c.Name == "" || seenCases[c.Name] || c.Object == nil || (c.Mode != "direct" && c.Mode != "composition") {
			t.Fatal("invalid case")
		}
		seenCases[c.Name] = true
		counts[c.Mode]++
		var kind string
		switch c.Expected.(type) {
		case nil:
			kind = "null"
		case bool:
			kind = "bool"
		case json.Number:
			kind = "int"
		case string:
			kind = "string"
		default:
			t.Fatal("expected scalar return")
		}
		expected = append(expected, map[string]any{"name": c.Name, "type": kind, "value": c.Expected, "input_after": c.Input, "source_input_after": c.Input,
			"feature_set_after": c.FeatureSet, "feature_after": c.Feature, "block_after": c.Block})
	}
	if counts["direct"] != 16 || counts["composition"] != 8 {
		t.Fatal("case mode counts")
	}
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
	// Reuse the bounded node transport, not the old test entrypoint. Each cut
	// delimiter and import replacement must occur exactly once; record both hashes.
	goTransportOriginal := read(t, "array-get-observer.go.txt")
	phpTransportOriginal := read(t, "array-get-observer.php")
	goCut, phpStart, phpEnd := "// init is compiled beside GoHP's generated main.go.", "const ARRAY_GET_MAX_DEPTH", "$cases = json_decode("
	if bytes.Count(goTransportOriginal, []byte(goCut)) != 1 || bytes.Count(phpTransportOriginal, []byte(phpStart)) != 1 ||
		bytes.Count(phpTransportOriginal, []byte(phpEnd)) != 1 {
		t.Fatal("transport adapter drift")
	}
	goTransport := append([]byte(nil), goTransportOriginal[:bytes.Index(goTransportOriginal, []byte(goCut))]...)
	if bytes.Count(goTransport, []byte("\n\t\"reflect\"")) != 1 {
		t.Fatal("Go transport import drift")
	}
	goTransport = bytes.Replace(goTransport, []byte("\n\t\"reflect\""), nil, 1)
	phpTransport := append([]byte("<?php\n"), phpTransportOriginal[bytes.Index(phpTransportOriginal, []byte(phpStart)):bytes.Index(phpTransportOriginal, []byte(phpEnd))]...)
	phpObserver, goObserver := read(t, "elements-source-input-observer.php"), read(t, "elements-source-input-observer.go.txt")
	for name, data := range map[string][]byte{"helpers.php": definitions, "cases.json": casesBytes, "observer.php": phpObserver, "transport.php": phpTransport, "wordpress-license.txt": read(t, "../../../license.txt")} {
		write(t, filepath.Join(input, name), data)
	}
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	write(t, filepath.Join(candidate, "transport.go"), goTransport)
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456", "--env=GOCACHE=/candidate/cache",
		"--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local", "--env=GOPROXY=off", "--env=GOSUMDB=off",
		"--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "elements_source_input_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/helpers.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
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
	if report.Schema != "gohp.build/v0.1" || report.Status != "generated" || report.Target != pin.PHPTarget || report.SourceHash != digest(definitions) || report.GeneratedHash != digest(generated) {
		t.Fatal("report/source/output mismatch")
	}
	t.Log("status=generated; three complete original declarations unchanged")
	_, compileErr := container(t, "elements_source_input_build", append(buildArgs, "--entrypoint=go", pin.GoImage, "build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/transport.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("build stderr: %s", compileErr)
	}
	t.Log("status=compiled; no compatibility claim before observations")
	actual, actualErr := container(t, "elements_source_input_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "elements_source_input_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "manifest_sha256": digest(manifest), "wordpress_commit": pin.WordPressCommit,
		"fragments": selections, "original_declaration_count": 3, "owned_driver_count": 2, "owned_input_class_count": 1,
		"compiled_bundle_sha256": digest(definitions), "native_bundle_sha256": digest(definitions), "drivers_sha256": digest(drivers),
		"cases_sha256": digest(casesBytes), "case_count": len(cases), "direct_case_count": counts["direct"], "composition_case_count": counts["composition"],
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "php_image": pin.PHPImage, "php_version": phpVersion, "php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver),
		"go_transport":  map[string]string{"original_sha256": digest(goTransportOriginal), "adapted_sha256": digest(goTransport)},
		"php_transport": map[string]string{"original_sha256": digest(phpTransportOriginal), "adapted_sha256": digest(phpTransport)},
		"binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "three unchanged originals, two owned functions and one input-data class; object construction/assignment in translated PHP, no host field injection;16 direct and8 composed retained cases; old24/334 configurations unchanged/not rerun; no original WP class/registry/render/CSS/HTML/bootstrap compatibility claim",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
