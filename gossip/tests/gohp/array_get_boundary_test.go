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

// PASS here records native behavior and explicit compiler rejection, not a
// successful Go/PHP comparison of the original.
func TestOriginalWordPressArrayGetBoundary(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE; cached pinned Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run as a non-root user")
	}
	manifest := read(t, "array-get-boundary.json")
	var pin unitPin
	if err := json.Unmarshal(manifest, &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.array-get-boundary/v1" || pin.SourcePath != "wp-includes/functions.php" || len(pin.Fragments) != 1 ||
		pin.PHPTarget != "8.3" || !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid array-get pin")
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
		t.Fatal("expected24 cases")
	}
	workspace := t.TempDir()
	source, input, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "input"), filepath.Join(workspace, "candidate")
	for _, dir := range []string{source, input, candidate} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}
	archiveHash := exportGoHP(t, repository, pin.GoHPCommit, source)
	observer := read(t, "array-get-observer.php")
	write(t, filepath.Join(input, "helpers.php"), definitions)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "observer.php"), observer)
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	oracle, oracleErr := runPHPOracle(t, "array_get_reference", pin.PHPImage, input, "/input/observer.php")
	if len(oracleErr) != 0 {
		t.Fatalf("native stderr: %s", oracleErr)
	}
	var observations []map[string]any
	numericBundleDecode(t, oracle, &observations)
	if len(observations) != len(cases) {
		t.Fatal("native count mismatch")
	}
	seen := map[string]bool{}
	for i, c := range cases {
		if c.Name == "" || seen[c.Name] || c.OmitDefault == nil {
			t.Fatal("invalid case")
		}
		seen[c.Name] = true
		kind := ""
		switch c.Expected.(type) {
		case nil:
			kind = "null"
		case bool:
			kind = "bool"
		case json.Number:
			kind = "int"
		case string:
			kind = "string"
		case map[string]any:
			kind = "array"
		default:
			t.Fatal("unsupported expected node")
		}
		o := observations[i]
		if o["name"] != c.Name || o["type"] != kind || !reflect.DeepEqual(o["value"], c.Expected) {
			t.Fatalf("native expectation mismatch %s: %#v", c.Name, o)
		}
		for _, field := range []string{"input_after", "path_after", "default_after"} {
			if _, ok := o[field]; !ok {
				t.Fatal("missing measured argument state")
			}
		}
	}
	t.Log("native_reference_pass; Go/PHP differential not-run")
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456", "--env=GOCACHE=/candidate/cache",
		"--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local", "--env=GOPROXY=off", "--env=GOSUMDB=off",
		"--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "array_get_compiler_build", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp")...)
	if len(buildErr) != 0 {
		t.Fatalf("compiler build stderr: %s", buildErr)
	}
	probes := []struct {
		name, diagnostic, role string
		body                   []byte
	}{
		{"_wp_array_get", `PHP002 line 2: expected , (at "=")`, "original_declaration", declaration},
		{"gossip_probe_default", `PHP002 line 2: expected , (at "=")`, "owned_default_parameter", []byte("function gossip_probe_default($value = null) { return $value; }\n")},
		{"gossip_probe_isset_var", `PHP002 line 2: expected [ (at ")")`, "owned_bare_isset", []byte("function gossip_probe_isset_var($value) { return isset($value); }\n")},
		{"gossip_probe_isset_multi", `PHP002 line 2: expected [ (at ",")`, "owned_multi_isset", []byte("function gossip_probe_isset_multi($items, $key) { return isset($key, $items[$key]); }\n")},
		{"gossip_probe_continue", `PHP002 line 2: unsupported statement (at "continue")`, "owned_continue", []byte("function gossip_probe_continue(array $items) { foreach ($items as $item) { continue; } return null; }\n")},
	}
	var rejections []map[string]any
	for _, probe := range probes {
		php := append([]byte("<?php\n"), probe.body...)
		write(t, filepath.Join(input, probe.name+".php"), php)
		script := fmt.Sprintf(`set +e
/candidate/gohp -in /input/%[1]s.php -out /candidate/%[1]s.go -php-version 8.3 -report /candidate/%[1]s.report.json > /candidate/%[1]s.stdout 2> /candidate/%[1]s.stderr
result=$?
set -e
printf '%%s\n' "$result" > /candidate/%[1]s.exit
`, probe.name)
		out, stderr := container(t, "array_get_rejection", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c", script)...)
		if len(out) != 0 || len(stderr) != 0 {
			t.Fatalf("probe wrapper output %s %s", out, stderr)
		}
		exitCode := strings.TrimSpace(string(read(t, filepath.Join(candidate, probe.name+".exit"))))
		stdout := read(t, filepath.Join(candidate, probe.name+".stdout"))
		diagnostic := read(t, filepath.Join(candidate, probe.name+".stderr"))
		if exitCode != "1" || len(stdout) != 0 || strings.TrimSpace(string(diagnostic)) != probe.diagnostic {
			t.Fatalf("unexpected rejection %s: exit=%s stdout=%s stderr=%s", probe.name, exitCode, stdout, diagnostic)
		}
		for _, suffix := range []string{".go", ".report.json"} {
			if _, err := os.Lstat(filepath.Join(candidate, probe.name+suffix)); !os.IsNotExist(err) {
				t.Fatal("rejected candidate published an artifact")
			}
		}
		rejections = append(rejections, map[string]any{"name": probe.name, "role": probe.role, "status": "rejected", "exit_code": 1,
			"source_sha256": digest(php), "stdout": string(stdout), "stderr": string(diagnostic), "generated_present": false, "report_present": false})
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "native_reference_pass_candidate_rejected", "differential_status": "not-run",
		"wordpress_commit": pin.WordPressCommit, "source_path": pin.SourcePath, "source_sha256": pin.SourceHash,
		"fragment":        map[string]any{"name": part.Name, "first_line": part.First, "last_line": part.Last, "sha256": digest(declaration)},
		"manifest_sha256": digest(manifest), "native_bundle_sha256": digest(definitions), "cases_sha256": digest(casesBytes),
		"case_count": len(cases), "original_declaration_count": 1, "owned_dependency_probe_count": 4, "php_observer_sha256": digest(observer),
		"php_image": pin.PHPImage, "php_version": phpVersion, "php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"native_stderr": "", "native_exit_code": 0, "typed_native_observations": observations,
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "compiler_rejections": rejections,
		"scope": "native reference for unchanged _wp_array_get and explicit default/isset/continue rejections; no Go/PHP differential success",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
