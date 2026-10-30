package gohp_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// A native-only, separately named low-budget reference. The generated profile
// fixes match100000 and has no compatible match1 entry; never label this a Go PASS.
func TestOriginalWordPressParserMatch1Reference(t *testing.T) {
	if os.Getenv("GOSSIP_PHP_PARSER_MATCH1_REFERENCE") == "" {
		t.Skip("set GOSSIP_PHP_PARSER_MATCH1_REFERENCE=1, fixed image and fresh evidence output")
	}
	output := os.Getenv("GOSSIP_PHP_PARSER_MATCH1_EVIDENCE_OUT")
	if os.Getenv("GOSSIP_PHP_PARSER_MATCH1_REFERENCE") != "1" || !filepath.IsAbs(output) || os.Getuid() == 0 {
		t.Fatal("requires opt-in1, fresh absolute evidence path and non-root user")
	}
	if _, err := os.Lstat(output); !os.IsNotExist(err) {
		t.Fatal("evidence path must be unused")
	}
	pinBytes := read(t, "parser-resource-match1.json")
	var pin struct {
		Schema   string `json:"schema"`
		Profile  string `json:"profile"`
		Manifest string `json:"reference_manifest_sha256"`
		Cases    string `json:"cases_sha256"`
		Observer string `json:"php_observer_sha256"`
	}
	numericBundleDecode(t, pinBytes, &pin)
	manifest := read(t, "parser-full-reference.json")
	cases, observer := read(t, "parser-resource-boundaries-cases.json"), read(t, "parser-resource-match1-observer.php")
	if pin.Schema != "gossip.parser-resource-match1/v1" || pin.Profile != "php83-original-parser-match1-v1" || digest(manifest) != pin.Manifest || digest(cases) != pin.Cases || digest(observer) != pin.Observer {
		t.Fatal("low-budget reference identities drifted")
	}
	var nativePin struct {
		Schema    string            `json:"schema"`
		PHPImage  string            `json:"php_image"`
		WordPress string            `json:"wordpress_commit"`
		PHPTarget string            `json:"php_target"`
		Cases     string            `json:"cases_sha256"`
		Observer  string            `json:"observer_sha256"`
		Sources   map[string]string `json:"sources"`
	}
	numericBundleDecode(t, manifest, &nativePin)
	if nativePin.Schema != "gossip.parser-full-reference/v1" || nativePin.PHPTarget != "8.3" || os.Getenv("GOSSIP_PHP_RUNNER_IMAGE") != nativePin.PHPImage || nativePin.WordPress != "54106f6c545fc8b881f6b39b7963002bece6c28c" || len(nativePin.Sources) != 5 {
		t.Fatal("requires fixed original source and exact cached image")
	}
	input := t.TempDir()
	for path, hash := range nativePin.Sources {
		data := read(t, filepath.Join("../../..", path))
		if digest(data) != hash {
			t.Fatal("original source drift: " + path)
		}
		write(t, filepath.Join(input, path), data)
	}
	write(t, filepath.Join(input, "cases.json"), cases)
	write(t, filepath.Join(input, "observer.php"), observer)
	evidence := map[string]any{
		"schema": "gossip.parser-resource-match1-evidence/v1", "status": "failed", "profile": pin.Profile,
		"application_acceptance": "not-established", "generated_comparison": "not-run-unsupported-budget",
		"generated_exclusion": "pcre2-php83-v1 generated entry fixes match100000 and exposes no match1 host setting; not changed or compared under different limits",
		"pin_sha256":          digest(pinBytes), "reference_manifest_sha256": digest(manifest), "cases_sha256": digest(cases), "php_observer_sha256": digest(observer), "harness_sha256": digest(read(t, "parser_resource_match1_test.go")),
		"wordpress_commit": nativePin.WordPress, "wordpress_sources": nativePin.Sources, "php_image": nativePin.PHPImage,
		"phases": map[string]string{"native_execute": "not-run", "native_reference": "not-run", "generated_execute": "not-run-unsupported-budget"},
	}
	defer func() {
		if t.Failed() {
			evidence["status"] = "failed"
		}
		encoded, err := json.Marshal(evidence)
		if err != nil {
			t.Error(err)
			return
		}
		file, err := os.OpenFile(output, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
		if err != nil {
			t.Error(err)
			return
		}
		_, writeErr := file.Write(append(encoded, '\n'))
		closeErr := file.Close()
		if writeErr != nil || closeErr != nil {
			t.Errorf("evidence write: %v %v", writeErr, closeErr)
		}
		t.Logf("PARSER_MATCH1_REFERENCE status=%v cases=%v calls=%v sha256=%s", evidence["status"], evidence["case_count"], evidence["parse_call_count"], digest(append(encoded, '\n')))
	}()
	stdout, stderr := runPHPOracle(t, "parser_resource_match1_native", nativePin.PHPImage, input, "/input/observer.php")
	evidence["native_stdout_sha256"], evidence["native_stdout_bytes"], evidence["native_stderr_bytes"] = digest(stdout), len(stdout), len(stderr)
	if len(stderr) != 0 {
		t.Fatal("native runtime diagnostics are not quiet")
	}
	evidence["native_exit"] = 0
	phases := evidence["phases"].(map[string]string)
	phases["native_execute"] = "pass"
	var native map[string]any
	numericBundleDecode(t, stdout, &native)
	evidence["native_observation"] = native
	expectedProfile := map[string]any{"pcre.jit": "0", "pcre.backtrack_limit": "1", "pcre.recursion_limit": "10000", "memory_limit": "128M"}
	if native["schema"] != "gossip.parser-full-native-observation/v1" || native["php_version"] != "8.3.36" || native["pcre_version"] != "10.42 2022-12-12" || native["int_size"] != json.Number("8") || native["lc_ctype"] != "C" || native["php_os"] != "Linux" || native["php_arch"] != "aarch64" || !reflect.DeepEqual(native["profile"], expectedProfile) {
		t.Fatal("native low-budget environment/profile mismatch")
	}
	var inputs []struct {
		Name     string `json:"name"`
		Document string `json:"document_base64"`
	}
	numericBundleDecode(t, cases, &inputs)
	rows, ok := native["cases"].([]any)
	if !ok || len(rows) != 12 || len(inputs) != len(rows) {
		t.Fatal("native reference case count mismatch")
	}
	codes, witnesses := map[string]int{}, []string{}
	for i, item := range rows {
		row := parserJSONBoundaryMap(t, item)
		in := parserJSONBoundaryMap(t, row["input"])
		if row["name"] != inputs[i].Name || in["document_base64"] != inputs[i].Document || row["caller_unchanged"] != true || row["reuse_reset"] != true {
			t.Fatal("native identity/caller/reuse mismatch")
		}
		failed := false
		for _, pass := range []string{"first_parse", "reuse_parse"} {
			call := parserJSONBoundaryMap(t, row[pass])
			code, ok := call["pcre_error_code"].(json.Number)
			if !ok || (code != "0" && code != "2") {
				t.Fatal("unexpected native resource outcome")
			}
			codes[string(code)]++
			failed = failed || code == "2"
		}
		if failed {
			witnesses = append(witnesses, inputs[i].Name)
		}
	}
	if len(witnesses) == 0 {
		t.Fatal("no actual original-pattern resource error observed")
	}
	canonical, _ := json.Marshal(native["cases"])
	evidence["native_typed_cases_sha256"], evidence["native_pcre_error_counts"], evidence["resource_error_case_names"] = digest(canonical), codes, witnesses
	evidence["case_count"], evidence["parse_call_count"], evidence["status"] = len(rows), 2*len(rows), "native-resource-reference-recorded"
	phases["native_reference"] = "pass"
}
