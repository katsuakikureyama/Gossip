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

// Compare the unchanged declaration, including its default argument and both
// isset/continue paths. The old rejection manifest remains independently pinned.
func TestOriginalWordPressArrayGetUnit(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE; cached pinned Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run as a non-root user")
	}
	manifest := read(t, "array-get-unit.json")
	var pin struct {
		unitPin
		CasesHash string `json:"cases_sha256"`
	}
	if err := json.Unmarshal(manifest, &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.array-get-unit/v1" || pin.SourcePath != "wp-includes/functions.php" || len(pin.Fragments) != 1 ||
		pin.PHPTarget != "8.3" || !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid array-get unit pin")
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
	original := read(t, filepath.Join("../../..", pin.SourcePath))
	if digest(original) != pin.SourceHash {
		t.Fatal("WordPress source drift")
	}
	part := pin.Fragments[0]
	lines := bytes.SplitAfter(original, []byte("\n"))
	if part.Name != "_wp_array_get" || part.First != 5233 || part.Last != 5272 || len(lines) < part.Last {
		t.Fatal("unexpected fragment")
	}
	declaration := bytes.Join(lines[part.First-1:part.Last], nil)
	if !bytes.HasPrefix(declaration, []byte("function _wp_array_get(")) || !bytes.HasSuffix(declaration, []byte("}\n")) {
		t.Fatal("incomplete declaration")
	}
	definitions := append([]byte("<?php\n"), declaration...)
	casesBytes := read(t, "array-get-cases.json")
	if digest(casesBytes) != pin.CasesHash {
		t.Fatal("native-reference case corpus drift")
	}
	var cases []struct {
		Name        string `json:"name"`
		Input       any    `json:"input"`
		Path        any    `json:"path"`
		Default     any    `json:"default"`
		Expected    any    `json:"expected"`
		OmitDefault *bool  `json:"omit_default"`
	}
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 24 {
		t.Fatal("expected24 retained cases")
	}
	workspace := t.TempDir()
	source, input, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "input"), filepath.Join(workspace, "candidate")
	for _, dir := range []string{source, input, candidate} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}
	archiveHash := exportGoHP(t, repository, pin.GoHPCommit, source)
	phpObserver := read(t, "array-get-observer.php")
	goObserver := read(t, "array-get-observer.go.txt")
	write(t, filepath.Join(input, "helpers.php"), definitions)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "observer.php"), phpObserver)
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456", "--env=GOCACHE=/candidate/cache",
		"--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local", "--env=GOPROXY=off", "--env=GOSUMDB=off",
		"--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "array_get_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
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
	if report.Schema != "gohp.build/v0.1" || report.Status != "generated" || report.Target != pin.PHPTarget ||
		report.SourceHash != digest(definitions) || report.GeneratedHash != digest(generated) {
		t.Fatal("report does not match exact source/output bytes")
	}
	t.Log("status=generated; original declaration unchanged, including null default")
	_, compileErr := container(t, "array_get_build", append(buildArgs, "--entrypoint=go", pin.GoImage,
		"build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("build stderr: %s", compileErr)
	}
	t.Log("status=compiled; no compatibility claim before observations")
	actual, actualErr := container(t, "array_get_candidate",
		"--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "array_get_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) {
		t.Fatalf("differential_failed: Go=%s stderr=%s PHP=%s stderr=%s", actual, actualErr, oracle, oracleErr)
	}
	var records []map[string]any
	numericBundleDecode(t, actual, &records)
	if len(records) != len(cases) {
		t.Fatal("observation count mismatch")
	}
	seen := map[string]bool{}
	for i, c := range cases {
		if c.Name == "" || seen[c.Name] || c.OmitDefault == nil || records[i]["name"] != c.Name || !reflect.DeepEqual(records[i]["value"], c.Expected) {
			t.Fatalf("case/expectation mismatch: %s", c.Name)
		}
		seen[c.Name] = true
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "manifest_sha256": digest(manifest),
		"wordpress_commit": pin.WordPressCommit, "source_path": pin.SourcePath, "source_sha256": pin.SourceHash,
		"fragment":                   map[string]any{"name": part.Name, "first_line": part.First, "last_line": part.Last, "sha256": digest(declaration)},
		"original_declaration_count": 1, "compiled_bundle_sha256": digest(definitions), "native_bundle_sha256": digest(definitions),
		"cases_sha256": digest(casesBytes), "case_count": len(cases), "gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash,
		"compiler_version": report.CompilerVersion, "compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))),
		"generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "php_image": pin.PHPImage, "php_version": phpVersion,
		"php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver),
		"binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "one unchanged original _wp_array_get;24 retained cases; measured scalar/nested array returns and caller argument state; separate from24-original262-case integration; full WordPress/callers not run",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
