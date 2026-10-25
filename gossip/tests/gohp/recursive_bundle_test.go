package gohp_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

// This is a new, independent recursive-helper binary. It preserves the
// earlier standalone and integrated-bundle observations without relabeling or
// rerunning them.
func TestOriginalWordPressRecursiveHelperBundle(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	manifest := read(t, "recursive-helper-bundle.json")
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
	if pin.Schema != "gossip.recursive-helper-bundle/v1" || len(pin.Sources) != 2 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid recursive helper bundle pin")
	}
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil {
		t.Fatal(err)
	}
	if baseline.Commit != pin.WordPressCommit {
		t.Fatal("bundle and fork baseline disagree")
	}

	wanted := map[string]string{
		"__return_true": "wp-includes/functions.php", "__return_false": "wp-includes/functions.php",
		"__return_zero": "wp-includes/functions.php", "__return_empty_array": "wp-includes/functions.php",
		"__return_null": "wp-includes/functions.php", "__return_empty_string": "wp-includes/functions.php",
		"wp_validate_boolean": "wp-includes/functions.php", "_autop_newline_preservation_helper": "wp-includes/formatting.php",
		"zeroise": "wp-includes/formatting.php", "trailingslashit": "wp-includes/formatting.php",
		"untrailingslashit": "wp-includes/formatting.php", "stripslashes_from_strings_only": "wp-includes/formatting.php",
		"wp_slash": "wp-includes/formatting.php",
	}
	var bundle bytes.Buffer
	bundle.WriteString("<?php\n")
	var selections []map[string]any
	seenSources, seenFunctions := map[string]bool{}, map[string]bool{}
	for _, sourcePin := range pin.Sources {
		if (sourcePin.Path != "wp-includes/functions.php" && sourcePin.Path != "wp-includes/formatting.php") || seenSources[sourcePin.Path] {
			t.Fatal("unexpected or duplicate original source")
		}
		seenSources[sourcePin.Path] = true
		original := read(t, filepath.Join("../../..", sourcePin.Path))
		if digest(original) != sourcePin.Hash {
			t.Fatal("original source drift")
		}
		lines := bytes.SplitAfter(original, []byte("\n"))
		for _, part := range sourcePin.Fragments {
			if wanted[part.Name] != sourcePin.Path || seenFunctions[part.Name] || part.First < 1 || part.Last < part.First || part.Last > len(lines) {
				t.Fatal("unexpected, duplicate or invalid fragment")
			}
			seenFunctions[part.Name] = true
			excerpt := bytes.Join(lines[part.First-1:part.Last], nil)
			if !bytes.HasPrefix(excerpt, []byte("function "+part.Name+"(")) || !bytes.HasSuffix(excerpt, []byte("}\n")) {
				t.Fatal("incomplete original function")
			}
			selections = append(selections, map[string]any{
				"name": part.Name, "source_path": sourcePin.Path, "source_sha256": sourcePin.Hash,
				"original_first_line": part.First, "original_last_line": part.Last,
				"bundle_first_line": bytes.Count(bundle.Bytes(), []byte("\n")) + 1, "sha256": digest(excerpt),
			})
			bundle.Write(excerpt)
			bundle.WriteByte('\n')
		}
	}
	if len(seenSources) != 2 || len(seenFunctions) != 13 {
		t.Fatal("bundle must contain exactly thirteen unchanged original functions")
	}
	drivers := read(t, "recursive-bundle-drivers.php.txt")
	driverFirstLine := bytes.Count(bundle.Bytes(), []byte("\n")) + 1
	bundle.Write(drivers) // Owned composition calls; never substitutes original bodies.
	definitions := bundle.Bytes()

	casesBytes := read(t, "recursive-bundle-cases.json")
	decoder := json.NewDecoder(bytes.NewReader(casesBytes))
	decoder.UseNumber()
	var cases []struct {
		Name      string `json:"name"`
		Operation string `json:"operation"`
		Args      []any  `json:"args"`
		Expected  struct {
			Type  string `json:"type"`
			Value any    `json:"value"`
		} `json:"expected"`
	}
	if err := decoder.Decode(&cases); err != nil {
		t.Fatal(err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		t.Fatal("recursive bundle fixture has trailing JSON values")
	}
	if len(cases) != 33 {
		t.Fatal("expected thirty-three owned recursive bundle cases")
	}
	expected := make([]map[string]any, 0, len(cases))
	seenCases, operations := map[string]bool{}, map[string]int{}
	for _, c := range cases {
		if c.Name == "" || seenCases[c.Name] {
			t.Fatal("missing or duplicate case name")
		}
		seenCases[c.Name] = true
		if err := validateIntegratedBundleJSON(c.Args); err != nil {
			t.Fatalf("invalid case arguments %s: %v", c.Name, err)
		}
		if err := validateIntegratedBundleJSON(c.Expected.Value); err != nil {
			t.Fatalf("invalid expected value %s: %v", c.Name, err)
		}
		if c.Expected.Type != integratedBundleJSONType(c.Expected.Value) {
			t.Fatalf("expected type mismatch for %s", c.Name)
		}
		operations[c.Operation]++
		expected = append(expected, map[string]any{
			"name": c.Name, "type": c.Expected.Type, "value": c.Expected.Value, "args_after": c.Args,
		})
	}
	for name := range wanted {
		if operations[name] != 1 {
			t.Fatalf("missing direct original probe: %s", name)
		}
	}
	for _, name := range []string{"gossip_zero_pad", "gossip_false_boolean", "gossip_empty_string_path", "gossip_path_pad"} {
		if operations[name] != 1 {
			t.Fatalf("missing scalar composition probe: %s", name)
		}
	}
	if operations["gossip_capture_newline_unslash"] != 4 || operations["gossip_unslash_array_keys"] != 2 || operations["gossip_unslash_array_shape"] != 2 ||
		operations["gossip_slash_scalar_roundtrip"] != 3 || operations["gossip_slash_list_shallow_roundtrip"] != 3 ||
		operations["gossip_slash_capture_newline_unslash"] != 1 || operations["gossip_slash_array_shape"] != 1 || len(operations) != 24 {
		t.Fatal("unexpected recursive composition coverage")
	}
	expectedBytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	phpVersion := requirePHPOracle(t, pin.PHPImage)
	workspace := t.TempDir()
	source, input, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "input"), filepath.Join(workspace, "candidate")
	for _, dir := range []string{source, input, candidate} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}
	archiveHash := exportGoHP(t, repository, pin.GoHPCommit, source)
	phpObserver, goObserver := read(t, "recursive-bundle-observer.php"), read(t, "recursive-bundle-observer.go.txt")
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
	buildOut, buildErr := container(t, "recursive_bundle_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
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
	if report.Schema != "gohp.build/v0.1" || report.Status != "generated" || report.Target != pin.PHPTarget || report.SourceHash != digest(definitions) || report.GeneratedHash != digest(generated) {
		t.Fatal("generation report does not match bundle bytes")
	}
	t.Log("status=generated; recursive source/report hashes verified")
	_, compileErr := container(t, "recursive_bundle_build", append(buildArgs, "--entrypoint=go", pin.GoImage, "build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("unexpected build stderr: %s", compileErr)
	}
	actual, actualErr := container(t, "recursive_bundle_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly", "--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "recursive_bundle_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("status=differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "wordpress_commit": pin.WordPressCommit,
		"manifest_sha256": digest(manifest), "selections": selections, "original_function_count": 13, "owned_driver_count": 11,
		"driver_first_line": driverFirstLine, "driver_sha256": digest(drivers), "bundle_sha256": digest(definitions),
		"cases_sha256": digest(casesBytes), "case_count": len(cases), "direct_call_count": 13, "composed_call_count": 20,
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "php_image": pin.PHPImage, "php_target": pin.PHPTarget, "go_version": strings.TrimSpace(string(buildOut)),
		"php_version": phpVersion, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "binary_sha256": digest(read(t, filepath.Join(candidate, "program"))),
		"typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "thirteen unchanged WordPress helpers and eleven owned composition drivers translated together in one binary; 33 bounded typed observations include finite nested lists, actual returned values, and caller arguments observed after each call; no claim of recursive unescaping, full-file/core/startup, HTTP, DB, theme or plugin compatibility",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
