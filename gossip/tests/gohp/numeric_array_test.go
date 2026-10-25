package gohp_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

// Separate successful-candidate unit; historical rejection evidence is unchanged.
func TestOriginalWordPressNumericArrayHelper(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	manifest := read(t, "numeric-array-helper.json")
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
	if pin.Schema != "gossip.numeric-array-observation/v1" || len(pin.Sources) != 2 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid numeric-array boundary pin")
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
		"wp_is_numeric_array": {"wp-includes/functions.php", 5448, 5454},
		"array_all_guard":     {"wp-includes/compat.php", 623, 644},
		"array_all":           {"wp-includes/compat.php", 635, 643},
	}
	fragments := map[string][]byte{}
	var selections []map[string]any
	for _, source := range pin.Sources {
		if source.Path != "wp-includes/functions.php" && source.Path != "wp-includes/compat.php" {
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
	if len(fragments) != len(wanted) || !bytes.Contains(fragments["array_all_guard"], fragments["array_all"]) {
		t.Fatal("missing original helper or full guarded polyfill")
	}
	// Preserve the complete original guard and body in the PHP 8.3 reference.
	// The compiled unit explicitly extracts the declaration, without rewriting it.
	nativeDefinitions := append([]byte("<?php\n"), fragments["array_all_guard"]...)
	nativeDefinitions = append(nativeDefinitions, fragments["wp_is_numeric_array"]...)
	drivers := read(t, "numeric-array-drivers.php.txt")
	nativeDefinitions = append(nativeDefinitions, drivers...)
	definitions := append([]byte("<?php\n"), fragments["array_all"]...)
	definitions = append(definitions, fragments["wp_is_numeric_array"]...)
	definitions = append(definitions, drivers...)
	casesBytes := read(t, "numeric-array-cases.json")
	var cases []struct {
		Name           string `json:"name"`
		InputType      string `json:"input_type"`
		Input          any    `json:"input"`
		Expected       *bool  `json:"expected"`
		ExpectedIsList *bool  `json:"expected_is_list"`
		InputAfter     any    `json:"input_after"`
	}
	decoder := json.NewDecoder(bytes.NewReader(casesBytes))
	decoder.UseNumber()
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cases); err != nil {
		t.Fatal(err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		t.Fatal("trailing corpus data")
	}
	if len(cases) != 24 {
		t.Fatal("expected twenty-four native reference cases")
	}
	expected := make([]map[string]any, 0, len(cases))
	seen := map[string]bool{}
	for _, c := range cases {
		if c.Name == "" || seen[c.Name] || c.Expected == nil || (c.InputType == "array") != (c.ExpectedIsList != nil) {
			t.Fatal("invalid native reference expectation")
		}
		seen[c.Name] = true
		validateNumericArrayTransport(t, c.InputType, c.Input)
		validateNumericArrayTransport(t, c.InputType, c.InputAfter)
		expected = append(expected, map[string]any{"name": c.Name, "input_type": c.InputType, "type": "bool", "value": *c.Expected,
			"is_list": c.ExpectedIsList, "input_after": map[string]any{"type": c.InputType, "value": c.InputAfter}})
	}
	polyfillCases := read(t, "numeric-array-polyfill-cases.json")
	polyfillExpected := numericPolyfillExpected(t, polyfillCases)
	expectedBytes, err := json.Marshal(map[string]any{"predicate": expected, "polyfill": polyfillExpected})
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
	reference := read(t, "numeric-array-observer.php")
	phpObserver := read(t, "numeric-array-differential-observer.php")
	goObserver := read(t, "numeric-array-observer.go.txt")
	write(t, filepath.Join(input, "helpers.php"), nativeDefinitions)
	write(t, filepath.Join(input, "unit.php"), definitions)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "polyfill-cases.json"), polyfillCases)
	write(t, filepath.Join(input, "reference.php"), reference)
	write(t, filepath.Join(input, "observer.php"), phpObserver)
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456",
		"--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local",
		"--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "numeric_array_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
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
	_, compileErr := container(t, "numeric_array_build", append(buildArgs, "--entrypoint=go", pin.GoImage,
		"build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("build stderr: %s", compileErr)
	}
	t.Log("status=compiled; no compatibility claim before observations")
	actual, actualErr := container(t, "numeric_array_candidate",
		"--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "numeric_array_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "manifest_sha256": digest(manifest),
		"wordpress_commit": pin.WordPressCommit, "fragments": selections, "original_declaration_count": 2, "owned_driver_count": 5,
		"compiled_bundle_sha256": digest(definitions), "native_bundle_sha256": digest(nativeDefinitions), "drivers_sha256": digest(drivers),
		"loading_guard": "native executes unchanged guard and verifies user-defined polyfill; compiler extracts unchanged declarations only",
		"cases_sha256":  digest(casesBytes), "polyfill_cases_sha256": digest(polyfillCases), "predicate_case_count": len(cases), "polyfill_case_count": len(polyfillExpected), "case_count": len(cases) + len(polyfillExpected),
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "php_image": pin.PHPImage, "php_version": phpVersion, "php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "reference_observer_sha256": digest(reference),
		"binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "two unchanged original declarations;24 predicate and7 direct-polyfill cases; actual bool/list/ordered typed caller input; flat scalars only; no bootstrap/conditional declarations/general callable/full-core compatibility; not a new combined17-original integration",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}

func numericPolyfillExpected(t *testing.T, data []byte) []map[string]any {
	t.Helper()
	var cases []struct {
		Name       string `json:"name"`
		Operation  string `json:"operation"`
		Input      any    `json:"input"`
		Expected   *bool  `json:"expected"`
		InputAfter any    `json:"input_after"`
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cases); err != nil {
		t.Fatal(err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		t.Fatal("trailing polyfill corpus")
	}
	if len(cases) != 7 {
		t.Fatal("expected seven direct polyfill cases")
	}
	var expected []map[string]any
	seen := map[string]bool{}
	for _, c := range cases {
		if c.Name == "" || seen[c.Name] || c.Expected == nil {
			t.Fatal("invalid polyfill case")
		}
		seen[c.Name] = true
		switch c.Operation {
		case "gossip_array_all_key_value", "gossip_array_all_short_circuit", "gossip_array_all_empty", "gossip_array_all_mutation":
		default:
			t.Fatal("unexpected polyfill operation")
		}
		validateNumericArrayTransport(t, "array", c.Input)
		validateNumericArrayTransport(t, "array", c.InputAfter)
		isList := true
		for i, raw := range c.InputAfter.([]any) {
			if raw.([]any)[0] != json.Number(strconv.Itoa(i)) {
				isList = false
			}
		}
		expected = append(expected, map[string]any{"name": c.Name, "input_type": "array", "type": "bool", "value": *c.Expected, "is_list": isList, "input_after": map[string]any{"type": "array", "value": c.InputAfter}})
	}
	return expected
}
