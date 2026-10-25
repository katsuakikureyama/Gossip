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

// This unit keeps the mixed-input Boolean helper separate from earlier
// constant-return and slash-path observations, including its compiler pin.
func TestOriginalWordPressBooleanHelper(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	var pin unitPin
	if err := json.Unmarshal(read(t, "boolean-helper.json"), &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.boolean-observation/v1" || pin.SourcePath != "wp-includes/functions.php" || len(pin.Fragments) != 1 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) ||
		phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid Boolean unit pin")
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
		t.Fatal("original WordPress source drift; review extraction ranges")
	}
	part := pin.Fragments[0]
	lines := bytes.SplitAfter(original, []byte("\n"))
	if part.Name != "wp_validate_boolean" || part.First < 1 || part.Last < part.First || part.Last > len(lines) {
		t.Fatal("invalid Boolean fragment selection")
	}
	excerpt := bytes.Join(lines[part.First-1:part.Last], nil)
	if !bytes.HasPrefix(excerpt, []byte("function wp_validate_boolean( $value ) {\n")) || !bytes.HasSuffix(excerpt, []byte("}\n")) {
		t.Fatal("fragment must contain the complete original one-argument Boolean helper")
	}
	definitions := append([]byte("<?php\n"), excerpt...)

	casesBytes := read(t, "boolean-cases.json")
	var cases []struct {
		Name      string `json:"name"`
		InputType string `json:"input_type"`
		Expected  bool   `json:"expected"`
	}
	if err := json.Unmarshal(casesBytes, &cases); err != nil {
		t.Fatal(err)
	}
	if len(cases) != 23 {
		t.Fatal("expected twenty-three owned mixed-input Boolean cases")
	}
	expected := make([]map[string]any, 0, len(cases))
	seen := map[string]bool{}
	for _, testCase := range cases {
		if testCase.Name == "" || seen[testCase.Name] {
			t.Fatal("missing or duplicate Boolean case ID")
		}
		seen[testCase.Name] = true
		switch testCase.InputType {
		case "null", "bool", "int", "string", "array":
		default:
			t.Fatalf("unsupported owned input type %q", testCase.InputType)
		}
		expected = append(expected, map[string]any{"name": testCase.Name, "input_type": testCase.InputType, "type": "bool", "value": testCase.Expected})
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
	phpObserver, goObserver := read(t, "boolean-observer.php"), read(t, "boolean-observer.go.txt")
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
	buildOut, buildErr := container(t, "boolean_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
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
	_, compileErr := container(t, "boolean_build", append(buildArgs, "--entrypoint=go", pin.GoImage,
		"build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("unexpected build stderr: %s", compileErr)
	}
	t.Log("status=compiled; same pinned runtime; not yet compatibility evidence")
	actual, actualErr := container(t, "boolean_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "boolean_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("status=differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": "gossip.boolean-observation/v1", "status": "differential_pass",
		"wordpress_commit": pin.WordPressCommit, "source_path": pin.SourcePath, "source_sha256": pin.SourceHash,
		"fragment":      map[string]any{"name": part.Name, "original_first_line": part.First, "original_last_line": part.Last, "bundle_first_line": 2, "sha256": digest(excerpt)},
		"bundle_sha256": digest(definitions), "cases_sha256": digest(casesBytes), "case_count": len(cases), "typed_call_count": len(cases),
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "php_image": pin.PHPImage, "php_target": pin.PHPTarget, "go_version": strings.TrimSpace(string(buildOut)),
		"php_version": phpVersion, "php_minimum": phporacle.Minimum,
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "binary_sha256": digest(read(t, filepath.Join(candidate, "program"))),
		"typed_observations": observed, "stderr": "", "exit_code": 0, "php_lc_ctype": "C",
		"scope": "one unchanged WordPress Boolean helper, twenty-three JSON-transported inputs limited to null, bool, signed integer, string and array values; PHP executes with LC_ALL=C and asserts LC_CTYPE=C; no floats, objects/maps, whole-file/core, HTTP or filesystem behavior",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
