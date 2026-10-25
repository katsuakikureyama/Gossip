package gohp_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
func TestOriginalWordPressArrayEdgeBoundary(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	manifest := read(t, "array-edge-boundary.json")
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
	if pin.Schema != "gossip.array-edge-boundary/v1" || len(pin.Sources) != 1 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid array-edge boundary pin")
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
		"array_last":        {"wp-includes/compat.php", 679, 685},
		"array_last_guard":  {"wp-includes/compat.php", 668, 686},
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
		t.Fatal("missing original edge fragments")
	}
	definitions := []byte("<?php\n")
	for _, name := range []string{"array_first", "array_last"} {
		if !bytes.Contains(fragments[name+"_guard"], fragments[name]) {
			t.Fatal("incomplete original guarded polyfill")
		}
		definitions = append(definitions, fragments[name+"_guard"]...)
	}
	casesBytes := read(t, "array-edge-cases.json")
	expected := arrayEdgeExpected(t, casesBytes)
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
	observer := read(t, "array-edge-observer.php")
	transportOriginal := read(t, "numeric-array-observer.php")
	start, end := []byte("function numeric_array_fixture_type("), []byte("if (! is_array($cases) || ! array_is_list($cases)) {")
	if bytes.Count(transportOriginal, start) != 1 || bytes.Count(transportOriginal, end) != 1 {
		t.Fatal("transport boundary drift")
	}
	transport := append([]byte("<?php\n"), transportOriginal[bytes.Index(transportOriginal, start):bytes.Index(transportOriginal, end)]...)
	write(t, filepath.Join(input, "transport.php"), transport)
	write(t, filepath.Join(input, "helpers.php"), definitions)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "observer.php"), observer)
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	oracle, oracleErr := runPHPOracle(t, "array_edge_reference", pin.PHPImage, input, "/input/observer.php")
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
	buildOut, buildErr := container(t, "array_edge_compiler_build", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp")...)
	if len(buildErr) != 0 {
		t.Fatalf("compiler build stderr: %s", buildErr)
	}
	var rejections []map[string]any
	for _, probe := range []struct {
		name, diagnostic, role string
		body                   []byte
	}{
		{"array_first", "PHP004 line 3: unsupported function empty", "original_declaration", fragments["array_first"]},
		{"array_last", "PHP004 line 3: unsupported function empty", "original_declaration", fragments["array_last"]},
		{"gossip_probe_array_key_last", "PHP004 line 2: unsupported function array_key_last", "owned_dependency_probe", read(t, "array-edge-key-probe.php.txt")},
	} {
		php := append([]byte("<?php\n"), probe.body...)
		write(t, filepath.Join(input, probe.name+".php"), php)
		// Names are hard-coded above, not sourced from fixture/manifest input.
		// Capture the actual nonzero CLI result without treating Docker startup
		// failure or a shell error as an expected compiler rejection.
		script := fmt.Sprintf(`set +e
/candidate/gohp -in /input/%[1]s.php -out /candidate/%[1]s.go -php-version 8.3 -report /candidate/%[1]s.report.json > /candidate/%[1]s.stdout 2> /candidate/%[1]s.stderr
result=$?
set -e
printf '%%s\n' "$result" > /candidate/%[1]s.exit
`, probe.name)
		out, stderr := container(t, "array_edge_rejection", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c", script)...)
		if len(out) != 0 || len(stderr) != 0 {
			t.Fatalf("probe wrapper output: %s %s", out, stderr)
		}
		exitCode := strings.TrimSpace(string(read(t, filepath.Join(candidate, probe.name+".exit"))))
		stdout := read(t, filepath.Join(candidate, probe.name+".stdout"))
		diagnostic := read(t, filepath.Join(candidate, probe.name+".stderr"))
		if exitCode != "1" || len(stdout) != 0 || strings.TrimSpace(string(diagnostic)) != probe.diagnostic {
			t.Fatalf("unexpected candidate result: %s exit=%s stdout=%s stderr=%s", probe.name, exitCode, stdout, diagnostic)
		}
		for _, suffix := range []string{".go", ".report.json"} {
			if _, err := os.Lstat(filepath.Join(candidate, probe.name+suffix)); !os.IsNotExist(err) {
				t.Fatalf("rejected candidate published artifact %s: %v", suffix, err)
			}
		}
		rejections = append(rejections, map[string]any{"name": probe.name, "role": probe.role, "source_sha256": digest(php), "status": "rejected", "exit_code": 1,
			"stdout": string(stdout), "stderr": string(diagnostic), "generated_present": false, "report_present": false})
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "native_reference_pass_candidate_rejected", "differential_status": "not-run",
		"manifest_sha256": digest(manifest), "wordpress_commit": pin.WordPressCommit, "fragments": selections,
		"native_bundle_sha256": digest(definitions), "cases_sha256": digest(casesBytes), "case_count": len(expected), "original_declaration_count": 2, "owned_dependency_probe_count": 1,
		"php_observer_sha256": digest(observer), "transport_source_sha256": digest(transportOriginal), "transport_selected_sha256": digest(transport), "php_version": phpVersion, "php_minimum": phporacle.Minimum,
		"php_image": pin.PHPImage, "php_target": pin.PHPTarget, "php_lc_ctype": "C", "native_stderr": "", "native_exit_code": 0,
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "typed_native_observations": observed, "compiler_rejections": rejections,
		"scope": "native reference only for two unchanged original array-edge guarded polyfills;24 flat scalar cases; original declarations rejected at empty and owned dependency probe rejected at array_key_last; no generated/report artifacts or candidate binary; no Go/PHP differential success",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}

func arrayEdgeExpected(t *testing.T, data []byte) []map[string]any {
	t.Helper()
	var cases []struct {
		Name      string `json:"name"`
		Operation string `json:"operation"`
		Input     any    `json:"input"`
		Expected  struct {
			Type  string `json:"type"`
			Value any    `json:"value"`
		} `json:"expected"`
		InputAfter any `json:"input_after"`
	}
	numericBundleDecode(t, data, &cases)
	if len(cases) != 24 {
		t.Fatal("expected twenty-four edge cases")
	}
	seen, counts := map[string]bool{}, map[string]int{}
	var observations []map[string]any
	for _, c := range cases {
		if c.Name == "" || seen[c.Name] {
			t.Fatal("invalid edge case name")
		}
		seen[c.Name] = true
		switch c.Operation {
		case "array_first", "array_last":
			counts[c.Operation]++
		default:
			t.Fatal("unknown edge operation")
		}
		validateNumericArrayTransport(t, "array", c.Input)
		validateNumericArrayTransport(t, "array", c.InputAfter)
		switch c.Expected.Type {
		case "null", "bool", "int", "string":
		default:
			t.Fatal("expected result outside scalar contract")
		}
		validateNumericArrayTransport(t, c.Expected.Type, c.Expected.Value)
		observations = append(observations, map[string]any{"name": c.Name, "type": c.Expected.Type, "value": c.Expected.Value, "input_after": c.InputAfter})
	}
	if counts["array_first"] != 12 || counts["array_last"] != 12 {
		t.Fatal("expected twelve cases per original")
	}
	return observations
}
