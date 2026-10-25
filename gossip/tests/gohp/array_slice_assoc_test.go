package gohp_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

type arraySliceAssocCase struct {
	Name       string `json:"name"`
	Input      []any  `json:"input"`
	Keys       []any  `json:"keys"`
	Expected   []any  `json:"expected"`
	InputAfter []any  `json:"input_after"`
	KeysAfter  []any  `json:"keys_after"`
}

func TestOriginalWordPressArraySliceAssocHelper(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	var pin unitPin
	if err := json.Unmarshal(read(t, "array-slice-assoc-helper.json"), &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.array-slice-assoc-observation/v1" || pin.SourcePath != "wp-includes/functions.php" || len(pin.Fragments) != 1 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) ||
		phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid wp_array_slice_assoc unit pin")
	}
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil {
		t.Fatal(err)
	}
	if baseline.Commit != pin.WordPressCommit {
		t.Fatal("unit and fork baseline disagree")
	}
	original := read(t, filepath.Join("../../..", pin.SourcePath))
	if digest(original) != pin.SourceHash {
		t.Fatal("original WordPress source drift; review extraction range")
	}
	part := pin.Fragments[0]
	lines := bytes.SplitAfter(original, []byte("\n"))
	if part.Name != "wp_array_slice_assoc" || part.First != 5173 || part.Last != 5183 || part.Last > len(lines) {
		t.Fatal("invalid wp_array_slice_assoc fragment selection")
	}
	excerpt := bytes.Join(lines[part.First-1:part.Last], nil)
	if !bytes.Equal(excerpt, []byte("function wp_array_slice_assoc( $input_array, $keys ): array {\n\t$slice = array();\n\n\tforeach ( $keys as $key ) {\n\t\tif ( isset( $input_array[ $key ] ) ) {\n\t\t\t$slice[ $key ] = $input_array[ $key ];\n\t\t}\n\t}\n\n\treturn $slice;\n}\n")) {
		t.Fatal("fragment must contain the complete unchanged original helper, including its array return declaration")
	}
	definitions := append([]byte("<?php\n"), excerpt...)

	casesBytes := read(t, "array-slice-assoc-cases.json")
	decoder := json.NewDecoder(bytes.NewReader(casesBytes))
	decoder.UseNumber()
	var cases []arraySliceAssocCase
	if err := decoder.Decode(&cases); err != nil {
		t.Fatal(err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		t.Fatal("array-slice-assoc corpus has trailing JSON values")
	}
	if len(cases) != 14 {
		t.Fatal("expected fourteen owned wp_array_slice_assoc cases")
	}
	expected := make([]map[string]any, 0, len(cases))
	seen := map[string]bool{}
	for _, testCase := range cases {
		if testCase.Name == "" || seen[testCase.Name] {
			t.Fatal("missing or duplicate wp_array_slice_assoc case ID")
		}
		seen[testCase.Name] = true
		input := normalizeArraySliceAssocPairs(t, testCase.Input, "input")
		keys := normalizeArraySliceAssocSelectors(t, testCase.Keys, "keys")
		expectedEntries := normalizeArraySliceAssocPairs(t, testCase.Expected, "expected")
		inputAfter := normalizeArraySliceAssocPairs(t, testCase.InputAfter, "input_after")
		keysAfter := normalizeArraySliceAssocSelectors(t, testCase.KeysAfter, "keys_after")
		if !reflect.DeepEqual(keys, keysAfter) {
			t.Fatalf("case %q must explicitly require unchanged selectors", testCase.Name)
		}
		expected = append(expected, map[string]any{
			"name": testCase.Name, "type": "array", "entries": expectedEntries,
			"input_after": map[string]any{"type": "array", "entries": inputAfter},
			"keys_after":  map[string]any{"type": "array", "value": keysAfter},
		})
		_ = input // Input is separately transported to both independent observers.
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
	phpObserver, goObserver := read(t, "array-slice-assoc-observer.php"), read(t, "array-slice-assoc-observer.go.txt")
	write(t, filepath.Join(input, "helpers.php"), definitions)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "observer.php"), phpObserver)
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456",
		"--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local",
		"--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "array_slice_assoc_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/helpers.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
	if len(buildErr) != 0 {
		t.Fatalf("unexpected compiler stderr: %s", buildErr)
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
		t.Fatal("status=rejected: report does not match exact source/output bytes")
	}
	t.Log("status=generated; hashes verified; not yet compatibility evidence")
	_, compileErr := container(t, "array_slice_assoc_build", append(buildArgs, "--entrypoint=go", pin.GoImage,
		"build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("unexpected build stderr: %s", compileErr)
	}
	t.Log("status=compiled; same pinned runtime; not yet compatibility evidence")
	actual, actualErr := container(t, "array_slice_assoc_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "array_slice_assoc_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("status=differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": "gossip.array-slice-assoc-observation/v1", "status": "differential_pass",
		"wordpress_commit": pin.WordPressCommit, "source_path": pin.SourcePath, "source_sha256": pin.SourceHash,
		"fragment":      map[string]any{"name": part.Name, "original_first_line": part.First, "original_last_line": part.Last, "bundle_first_line": 2, "sha256": digest(excerpt)},
		"bundle_sha256": digest(definitions), "cases_sha256": digest(casesBytes), "case_count": len(cases), "typed_call_count": len(cases),
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "php_image": pin.PHPImage, "php_target": pin.PHPTarget, "go_version": strings.TrimSpace(string(buildOut)),
		"php_version": phpVersion, "php_minimum": phporacle.Minimum,
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "binary_sha256": digest(read(t, filepath.Join(candidate, "program"))),
		"typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "one unchanged WordPress associative array-slice helper; fourteen JSON-transported ordered key/value inputs and selector lists limited to bounded integer/string keys plus null, bool, modest signed integer, string and nested sequential list values; records ordered return entries and caller input/selectors after invocation; no floats, JSON objects, nested associative values, references, cycles, whole-file/core, HTTP or filesystem behavior",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}

func normalizeArraySliceAssocKey(t *testing.T, value any, field string) any {
	t.Helper()
	switch typed := value.(type) {
	case string:
		return typed
	case json.Number:
		integer, err := typed.Int64()
		if err != nil || integer < -1_000_000 || integer > 1_000_000 {
			t.Fatalf("%s must be a bounded signed integer or string key: %v", field, value)
		}
		return integer
	default:
		t.Fatalf("%s must be an integer or string key, got %T", field, value)
		return nil
	}
}

func normalizeArraySliceAssocValue(t *testing.T, value any, depth int, field string) any {
	t.Helper()
	if depth > 8 {
		t.Fatalf("%s exceeds depth eight", field)
	}
	switch typed := value.(type) {
	case nil, bool, string:
		return typed
	case json.Number:
		integer, err := typed.Int64()
		if err != nil || integer < -1_000_000 || integer > 1_000_000 {
			t.Fatalf("%s must be a modest signed integer: %v", field, value)
		}
		return integer
	case []any:
		out := make([]any, len(typed))
		for index, item := range typed {
			out[index] = normalizeArraySliceAssocValue(t, item, depth+1, fmt.Sprintf("%s[%d]", field, index))
		}
		return out
	case map[string]any:
		t.Fatalf("%s must not contain JSON maps/objects", field)
		return nil
	default:
		t.Fatalf("%s has unsupported JSON value type %T", field, value)
		return nil
	}
}

func normalizeArraySliceAssocPairs(t *testing.T, pairs []any, field string) [][]any {
	t.Helper()
	if pairs == nil {
		t.Fatalf("%s must be an explicit list, not absent/null", field)
	}
	out := make([][]any, 0, len(pairs))
	for index, rawPair := range pairs {
		pair, ok := rawPair.([]any)
		if !ok || len(pair) != 2 {
			t.Fatalf("%s pair %d must contain exactly key and value", field, index)
		}
		out = append(out, []any{
			normalizeArraySliceAssocKey(t, pair[0], fmt.Sprintf("%s pair %d key", field, index)),
			normalizeArraySliceAssocValue(t, pair[1], 0, fmt.Sprintf("%s pair %d value", field, index)),
		})
	}
	return out
}

func normalizeArraySliceAssocSelectors(t *testing.T, selectors []any, field string) []any {
	t.Helper()
	if selectors == nil {
		t.Fatalf("%s must be an explicit list, not absent/null", field)
	}
	out := make([]any, len(selectors))
	for index, selector := range selectors {
		out[index] = normalizeArraySliceAssocKey(t, selector, fmt.Sprintf("%s selector %d", field, index))
	}
	return out
}
