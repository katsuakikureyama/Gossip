package gohp_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

// This is native reference evidence only: no compiler or alternate regex engine.
func TestOriginalWordPressParserTokenReference(t *testing.T) {
	image := os.Getenv("GOSSIP_PHP_RUNNER_IMAGE")
	if image == "" {
		t.Skip("set GOSSIP_PHP_RUNNER_IMAGE to the admitted cached PHP image; native reference only")
	}
	if err := phporacle.ValidateImage(image); err != nil {
		t.Fatal(err)
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	manifest := read(t, "parser-token-reference.json")
	var pin struct {
		Schema          string            `json:"schema"`
		WordPressCommit string            `json:"wordpress_commit"`
		PHPImage        string            `json:"php_image"`
		PHPTarget       string            `json:"php_target"`
		CasesSHA256     string            `json:"cases_sha256"`
		PatternSHA256   string            `json:"pattern_sha256"`
		Sources         map[string]string `json:"sources"`
	}
	numericBundleDecode(t, manifest, &pin)
	wanted := map[string]string{
		"wp-includes/class-wp-block-parser.php":       "8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b",
		"wp-includes/class-wp-block-parser-block.php": "7e21f483cf0c9deddd2144e026842bf906c72983908f7bd163c0716e8f6f325e",
		"wp-includes/class-wp-block-parser-frame.php": "e5883d82dff099a4b1a23786e9a53b7e29aa73654db83e7ddc444ba17df5e105",
		"license.txt":             "49e97f629ce763abb2f8f1f13da74cb5a8a9a72a0e72a7477d1540c98cef5141",
		"wp-includes/version.php": "6fbc8cf7e4667ba118b92186eea0c9ec35154476790c218d55ead45eb2a1f1d1",
	}
	if pin.Schema != "gossip.parser-token-reference/v1" || pin.WordPressCommit != "54106f6c545fc8b881f6b39b7963002bece6c28c" || pin.PHPTarget != "8.3" || pin.PHPImage != image || !reflect.DeepEqual(pin.Sources, wanted) {
		t.Fatal("reference pin or complete original source selection drift")
	}
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil || baseline.Commit != pin.WordPressCommit {
		t.Fatal("reference differs from retained WordPress baseline")
	}
	verified := map[string][]byte{}
	for path, hash := range pin.Sources {
		data := read(t, filepath.Join("../../..", path))
		if digest(data) != hash {
			t.Fatalf("source drift: %s", path)
		}
		verified[path] = data
	}
	// Select the exact single-quoted pattern literal, not a handwritten regex.
	// This pinned literal has no PHP escaped quotes or double backslashes;
	// therefore removing its quotes leaves the runtime string bytes unchanged.
	lines := bytes.Split(verified["wp-includes/class-wp-block-parser.php"], []byte("\n"))
	literal := bytes.TrimSpace(lines[250])
	if !bytes.HasPrefix(literal, []byte("'")) || !bytes.HasSuffix(literal, []byte("',")) {
		t.Fatal("pattern line shape drift")
	}
	pattern := literal[1 : len(literal)-2]
	if bytes.Contains(pattern, []byte("'")) || bytes.Contains(pattern, []byte(`\\`)) || digest(pattern) != pin.PatternSHA256 {
		t.Fatal("pattern literal requires a new reviewed extraction")
	}
	casesBytes := read(t, "parser-token-reference-cases.json")
	if digest(casesBytes) != pin.CasesSHA256 {
		t.Fatal("case corpus drift")
	}
	var cases []struct {
		Name     string `json:"name"`
		Document string `json:"document_base64"`
		Offset   int    `json:"offset"`
		Limit    int    `json:"backtrack_limit"`
		Match    any    `json:"expected_match"`
		Error    int    `json:"expected_pcre_error"`
		Token    []any  `json:"expected_token"`
	}
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 24 {
		t.Fatal("requires 24 predeclared cases")
	}
	seen := map[string]bool{}
	for _, c := range cases {
		doc, err := base64.StdEncoding.Strict().DecodeString(c.Document)
		if err != nil || base64.StdEncoding.EncodeToString(doc) != c.Document || len(doc) > 4096 || c.Offset < -len(doc) || c.Offset > len(doc) || c.Name == "" || seen[c.Name] || (c.Limit != 1 && c.Limit != 100000) || len(c.Token) != 5 {
			t.Fatalf("case outside fixed domain: %s", c.Name)
		}
		seen[c.Name] = true
	}
	input := t.TempDir()
	for path, data := range verified {
		write(t, filepath.Join(input, path), data)
	}
	observer := read(t, "parser-token-reference-observer.php")
	write(t, filepath.Join(input, "observer.php"), observer)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "pattern.txt"), pattern)
	out, stderr := runPHPOracle(t, "parser_token_native_reference", image, input, "/input/observer.php")
	if len(stderr) != 0 {
		t.Fatalf("unexpected native diagnostics: %s", stderr)
	}
	var observed struct {
		Schema     string `json:"schema"`
		Status     string `json:"status"`
		PHPVersion string `json:"php_version"`
		Pattern    string `json:"pattern_sha256"`
		Cases      []struct {
			Name            string `json:"name"`
			CallerUnchanged bool   `json:"caller_unchanged"`
		} `json:"cases"`
	}
	if err := json.Unmarshal(out, &observed); err != nil {
		t.Fatal(err)
	}
	if phporacle.ValidateVersion(observed.PHPVersion) != nil || observed.Schema != "gossip.parser-token-native-observation/v1" || observed.Status != "native-reference-pass" || observed.Pattern != pin.PatternSHA256 || len(observed.Cases) != len(cases) {
		t.Fatal("incomplete native observation")
	}
	for i, c := range observed.Cases {
		if c.Name != cases[i].Name || !c.CallerUnchanged {
			t.Fatal("case identity or mutation mismatch")
		}
	}
	evidence := map[string]any{
		"schema": "gossip.parser-token-reference-evidence/v1", "status": "native-reference-pass", "differential_status": "not-run",
		"compiler_status": "not-run", "backend_comparison": "not-run", "php_target": pin.PHPTarget, "php_image": image,
		"wordpress_commit": pin.WordPressCommit, "sources": pin.Sources, "manifest_sha256": digest(manifest), "cases_sha256": digest(casesBytes),
		"observer_sha256": digest(observer), "harness_sha256": digest(read(t, "parser_token_reference_test.go")),
		"runner_policy_sha256": digest(read(t, "../../internal/phporacle/policy.go")), "container_helper_sha256": digest(read(t, "helpers_test.go")),
		"oracle_helper_sha256":   digest(read(t, "php_oracle_test.go")),
		"pattern_literal_sha256": digest(literal[:len(literal)-1]), "pattern_sha256": digest(pattern), "pattern_line": 251,
		"go_host_version": runtime.Version(), "go_host_os": runtime.GOOS, "go_host_arch": runtime.GOARCH,
		"native_exit_code": 0, "native_stderr_bytes": 0, "native_stdout_bytes": len(out), "native_stdout_sha256": digest(out), "native_observation": jsonObservation(t, out),
		"scope": "24 fixed cases; complete original PHP next_token plus separate direct preg_match; native PHP8.3 reference only; no parse loop, generated Go, alternate backend or PHP8.4 evidence",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TOKEN_REFERENCE_EVIDENCE " + string(encoded))
}
