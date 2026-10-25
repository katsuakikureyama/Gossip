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

type utf8CharsetCase struct {
	Name      string `json:"name"`
	InputType string `json:"input_type"`
	Input     any    `json:"input"`
	Expected  bool   `json:"expected"`
}

// This independent unit deliberately covers just the unchanged compat.php
// charset predicate. It neither admits the compat file nor claims arbitrary
// case-insensitive string operations beyond this exact original helper.
func TestOriginalWordPressUTF8CharsetHelper(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	var pin unitPin
	manifestBytes := read(t, "utf8-charset-helper.json")
	if err := json.Unmarshal(manifestBytes, &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.utf8-charset-observation/v1" || pin.SourcePath != "wp-includes/compat.php" || len(pin.Fragments) != 1 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) ||
		phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid UTF-8 charset unit pin")
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
	if part.Name != "_is_utf8_charset" || part.First != 102 || part.Last != 111 || part.Last > len(lines) {
		t.Fatal("invalid _is_utf8_charset fragment selection")
	}
	excerpt := bytes.Join(lines[part.First-1:part.Last], nil)
	if !bytes.Equal(excerpt, []byte("function _is_utf8_charset( $charset_slug ) {\n\tif ( ! is_string( $charset_slug ) ) {\n\t\treturn false;\n\t}\n\n\treturn (\n\t\t0 === strcasecmp( 'UTF-8', $charset_slug ) ||\n\t\t0 === strcasecmp( 'UTF8', $charset_slug )\n\t);\n}\n")) {
		t.Fatal("fragment must contain the complete byte-unchanged original UTF-8 charset helper")
	}
	definitions := append([]byte("<?php\n"), excerpt...)

	casesBytes := read(t, "utf8-charset-cases.json")
	decoder := json.NewDecoder(bytes.NewReader(casesBytes))
	decoder.UseNumber()
	var cases []utf8CharsetCase
	if err := decoder.Decode(&cases); err != nil {
		t.Fatal(err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		t.Fatal("UTF-8 charset corpus has trailing JSON values")
	}
	if len(cases) != 28 {
		t.Fatal("expected twenty-eight owned UTF-8 charset cases")
	}
	expected := make([]map[string]any, 0, len(cases))
	seen := map[string]bool{}
	for _, testCase := range cases {
		if testCase.Name == "" || seen[testCase.Name] {
			t.Fatal("missing or duplicate UTF-8 charset case ID")
		}
		seen[testCase.Name] = true
		input := normalizeUTF8CharsetInput(t, testCase.Input, 0, "input")
		inputType := utf8CharsetInputType(t, input)
		if testCase.InputType != inputType {
			t.Fatalf("case %q input_type %q does not match transported input type %q", testCase.Name, testCase.InputType, inputType)
		}
		expected = append(expected, map[string]any{
			"name": testCase.Name, "input_type": inputType, "type": "bool", "value": testCase.Expected,
			"input_after": map[string]any{"type": inputType, "value": input},
		})
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
	phpObserver, goObserver := read(t, "utf8-charset-observer.php"), read(t, "utf8-charset-observer.go.txt")
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
	buildOut, buildErr := container(t, "utf8_charset_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
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
	_, compileErr := container(t, "utf8_charset_build", append(buildArgs, "--entrypoint=go", pin.GoImage,
		"build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("unexpected build stderr: %s", compileErr)
	}
	t.Log("status=compiled; same pinned runtime; not yet compatibility evidence")
	actual, actualErr := container(t, "utf8_charset_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "utf8_charset_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("status=differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": "gossip.utf8-charset-observation/v1", "status": "differential_pass",
		"manifest_sha256": digest(manifestBytes), "wordpress_commit": pin.WordPressCommit, "source_path": pin.SourcePath, "source_sha256": pin.SourceHash,
		"fragment":      map[string]any{"name": part.Name, "original_first_line": part.First, "original_last_line": part.Last, "bundle_first_line": 2, "sha256": digest(excerpt)},
		"bundle_sha256": digest(definitions), "cases_sha256": digest(casesBytes), "case_count": len(cases), "typed_call_count": len(cases), "input_after_call_count": len(cases),
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "php_image": pin.PHPImage, "php_target": pin.PHPTarget, "go_version": strings.TrimSpace(string(buildOut)),
		"php_version": phpVersion, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "binary_sha256": digest(read(t, filepath.Join(candidate, "program"))),
		"typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "one byte-unchanged WordPress compat.php UTF-8 charset predicate; twenty-eight JSON-transported inputs limited to null, bool, modest signed integer, string and nested sequential lists; observes actual bool returns and unchanged caller input after every call; PHP executes with LC_ALL=C and asserts LC_CTYPE=C; no floats, JSON objects, references, cycles, whole compat file/core, HTTP or filesystem behavior",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}

func normalizeUTF8CharsetInput(t *testing.T, value any, depth int, field string) any {
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
			out[index] = normalizeUTF8CharsetInput(t, item, depth+1, fmt.Sprintf("%s[%d]", field, index))
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

func utf8CharsetInputType(t *testing.T, value any) string {
	t.Helper()
	switch value.(type) {
	case nil:
		return "null"
	case bool:
		return "bool"
	case int64:
		return "int"
	case string:
		return "string"
	case []any:
		return "array"
	default:
		t.Fatalf("unexpected normalized UTF-8 charset input type %T", value)
		return ""
	}
}
