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

// A passing boundary test means the native reference matched and translation
// was REJECTED. It must never be counted as a Go/PHP differential success.
func TestOriginalWordPressNumericArrayBoundary(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	manifest := read(t, "numeric-array-boundary.json")
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
	if pin.Schema != "gossip.numeric-array-boundary/v1" || len(pin.Sources) != 2 || pin.PHPTarget != "8.3" ||
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
	// The declaration-only compiler probe also preserves its original indentation.
	definitions := append([]byte("<?php\n"), fragments["array_all_guard"]...)
	definitions = append(definitions, fragments["wp_is_numeric_array"]...)
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
	observer := read(t, "numeric-array-observer.php")
	write(t, filepath.Join(input, "helpers.php"), definitions)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "observer.php"), observer)
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	oracle, oracleErr := runPHPOracle(t, "numeric_array_reference", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, oracle)
	if len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("native_reference_failed: PHP=%s stderr=%s expected=%s", oracle, oracleErr, expectedBytes)
	}
	t.Log("native_reference_pass; NOT Go/PHP differential evidence")
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456",
		"--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local",
		"--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "numeric_array_compiler_build", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp")...)
	if len(buildErr) != 0 {
		t.Fatalf("compiler build stderr: %s", buildErr)
	}
	var rejections []map[string]any
	for _, probe := range []struct{ name, diagnostic string }{
		{"wp_is_numeric_array", "only array return declarations are supported"},
		{"array_all", "expected var"},
	} {
		php := append([]byte("<?php\n"), fragments[probe.name]...)
		write(t, filepath.Join(input, probe.name+".php"), php)
		// Names are hard-coded above, not sourced from fixture/manifest input.
		// Capture the actual nonzero CLI result without treating Docker startup
		// failure or a shell error as an expected parser rejection.
		script := fmt.Sprintf(`set +e
/candidate/gohp -in /input/%[1]s.php -out /candidate/%[1]s.go -php-version 8.3 -report /candidate/%[1]s.report.json > /candidate/%[1]s.stdout 2> /candidate/%[1]s.stderr
result=$?
set -e
printf '%%s\n' "$result" > /candidate/%[1]s.exit
`, probe.name)
		out, stderr := container(t, "numeric_array_rejection", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c", script)...)
		if len(out) != 0 || len(stderr) != 0 {
			t.Fatalf("probe wrapper output: %s %s", out, stderr)
		}
		exitCode := strings.TrimSpace(string(read(t, filepath.Join(candidate, probe.name+".exit"))))
		stdout := read(t, filepath.Join(candidate, probe.name+".stdout"))
		diagnostic := read(t, filepath.Join(candidate, probe.name+".stderr"))
		if exitCode != "1" || len(stdout) != 0 || !strings.HasPrefix(string(diagnostic), "PHP002 line 2:") || !strings.Contains(string(diagnostic), probe.diagnostic) {
			t.Fatalf("unexpected candidate result: %s exit=%s stdout=%s stderr=%s", probe.name, exitCode, stdout, diagnostic)
		}
		for _, suffix := range []string{".go", ".report.json"} {
			if _, err := os.Lstat(filepath.Join(candidate, probe.name+suffix)); !os.IsNotExist(err) {
				t.Fatalf("rejected candidate published artifact %s: %v", suffix, err)
			}
		}
		rejections = append(rejections, map[string]any{"name": probe.name, "source_sha256": digest(php), "status": "rejected", "exit_code": 1,
			"stdout": string(stdout), "stderr": string(diagnostic), "generated_present": false, "report_present": false})
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "native_reference_pass_candidate_rejected", "differential_status": "not-run",
		"manifest_sha256": digest(manifest), "wordpress_commit": pin.WordPressCommit, "fragments": selections,
		"native_bundle_sha256": digest(definitions), "cases_sha256": digest(casesBytes), "case_count": len(cases),
		"php_observer_sha256": digest(observer), "php_version": phpVersion, "php_minimum": phporacle.Minimum,
		"php_image": pin.PHPImage, "php_target": pin.PHPTarget, "php_lc_ctype": "C", "native_stderr": "", "native_exit_code": 0,
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "typed_native_observations": observed, "compiler_rejections": rejections,
		"scope": "native reference only using complete unchanged WordPress helper and guarded array_all polyfill; flat scalar-valued arrays and scalars; isolated compiler declarations rejected without generated/report artifacts; no generated candidate binary or Go/PHP behavior comparison",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}

func validateNumericArrayTransport(t *testing.T, kind string, value any) {
	t.Helper()
	if kind != "array" {
		if numericArrayScalarType(t, value) != kind {
			t.Fatal("input type mismatch")
		}
		return
	}
	pairs, ok := value.([]any)
	if !ok {
		t.Fatal("array transport requires ordered pairs")
	}
	for _, item := range pairs {
		pair, ok := item.([]any)
		if !ok || len(pair) != 2 {
			t.Fatal("array entry requires key and scalar value")
		}
		keyType := numericArrayScalarType(t, pair[0])
		if keyType != "int" && keyType != "string" {
			t.Fatal("array key must be int/string")
		}
		numericArrayScalarType(t, pair[1])
	}
}

func numericArrayScalarType(t *testing.T, value any) string {
	t.Helper()
	switch v := value.(type) {
	case nil:
		return "null"
	case bool:
		return "bool"
	case string:
		return "string"
	case json.Number:
		n, err := v.Int64()
		if err != nil || n < -1000000 || n > 1000000 {
			t.Fatal("integer transport outside bounded domain")
		}
		return "int"
	default:
		t.Fatalf("unsupported transport value %T", value)
		return ""
	}
}
