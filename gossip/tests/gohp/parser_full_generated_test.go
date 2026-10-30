package gohp_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

type fullParserGeneratedPin struct {
	Schema                  string `json:"schema"`
	GoHPCommit              string `json:"gohp_commit"`
	GenerationCommit        string `json:"reusable_generation_commit"`
	GeneratedSHA256         string `json:"generated_sha256"`
	ReusableReportSHA256    string `json:"reusable_report_sha256"`
	ReferenceSHA256         string `json:"reference_sha256"`
	ReferenceManifestSHA256 string `json:"reference_manifest_sha256"`
	ClosureSHA256           string `json:"closure_manifest_sha256"`
	ObserverSHA256          string `json:"observer_sha256"`
}

// This is an opt-in execution of the complete generated original, not another
// native-reference capture or an empty-main smoke test. Retained mode never
// starts Docker; fresh mode separately requires the fixed native oracle.
func TestOriginalWordPressFullParserGenerated(t *testing.T) {
	mode := os.Getenv("GOSSIP_GOHP_FULL_PARSER_MODE")
	if mode == "" {
		t.Skip("set GOSSIP_GOHP_FULL_PARSER_MODE=retained or fresh, GOSSIP_GOHP_FULL_PARSER_SOURCE and GOSSIP_PCRE2_PREFIX")
	}
	if mode != "retained" && mode != "fresh" {
		t.Fatal("full parser mode must be retained or fresh")
	}
	repository, prefix := os.Getenv("GOSSIP_GOHP_FULL_PARSER_SOURCE"), os.Getenv("GOSSIP_PCRE2_PREFIX")
	if !filepath.IsAbs(repository) || !filepath.IsAbs(prefix) || os.Getuid() == 0 {
		t.Fatal("requires explicit absolute source/prefix paths and a non-root local user")
	}
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" || runtime.Version() != "go1.23.0" {
		t.Fatal("this measured native build profile is Darwin ARM64 Go1.23.0 only")
	}
	pinBytes := read(t, "parser-full-generated.json")
	var pin fullParserGeneratedPin
	numericBundleDecode(t, pinBytes, &pin)
	if pin.Schema != "gossip.parser-full-generated/v1" || pin.GoHPCommit != "b2d1482268a7710c515c07bf919af3d40b09d17c" || pin.GenerationCommit != "35524a636073f4abbc8bb69b9e36263e71ac62d8" {
		t.Fatal("full generated parser series identity drift")
	}
	manifest := read(t, "parser-full-reference.json")
	var nativePin struct {
		Schema          string            `json:"schema"`
		WordPressCommit string            `json:"wordpress_commit"`
		PHPImage        string            `json:"php_image"`
		PHPTarget       string            `json:"php_target"`
		CasesSHA256     string            `json:"cases_sha256"`
		ObserverSHA256  string            `json:"observer_sha256"`
		Sources         map[string]string `json:"sources"`
	}
	// The native pin keeps its original field names and identity; do not repin it.
	numericBundleDecode(t, manifest, &nativePin)
	if digest(manifest) != pin.ReferenceManifestSHA256 || nativePin.Schema != "gossip.parser-full-reference/v1" || nativePin.WordPressCommit != "54106f6c545fc8b881f6b39b7963002bece6c28c" || nativePin.PHPTarget != "8.3" {
		t.Fatal("native full-parser manifest changed")
	}
	verified := map[string][]byte{}
	for path, hash := range nativePin.Sources {
		data := read(t, filepath.Join("../../..", path))
		if digest(data) != hash {
			t.Fatalf("original source drift: %s", path)
		}
		verified[path] = data
	}
	if len(verified) != 5 {
		t.Fatal("requires complete three-class/notice/provenance closure")
	}
	casesBytes, phpObserver := read(t, "parser-full-reference-cases.json"), read(t, "parser-full-reference-observer.php")
	observer := read(t, "parser-full-generated-observer.go.txt")
	closure := read(t, "parser-source-closure.json")
	if digest(casesBytes) != nativePin.CasesSHA256 || digest(phpObserver) != nativePin.ObserverSHA256 || digest(observer) != pin.ObserverSHA256 || digest(closure) != pin.ClosureSHA256 {
		t.Fatal("fixed corpus/observers/closure drift")
	}
	var cases []struct {
		Name     string `json:"name"`
		Document string `json:"document_base64"`
	}
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 12 {
		t.Fatal("requires all twelve original cases, without filtering")
	}
	seen := map[string]bool{}
	for _, item := range cases {
		document, err := base64.StdEncoding.Strict().DecodeString(item.Document)
		if err != nil || base64.StdEncoding.EncodeToString(document) != item.Document || len(document) > 4096 || item.Name == "" || seen[item.Name] {
			t.Fatal("invalid fixed case")
		}
		seen[item.Name] = true
	}
	referenceBytes := read(t, "../../../docs/evidence/parser-full-reference-2026-10-27.json")
	if digest(referenceBytes) != pin.ReferenceSHA256 {
		t.Fatal("retained native evidence changed")
	}
	var reference map[string]any
	numericBundleDecode(t, referenceBytes, &reference)
	native, ok := reference["native_observation"].(map[string]any)
	if !ok || reference["status"] != "native-reference-recorded" || reference["native_exit_code"] != json.Number("0") || reference["native_stderr_bytes"] != json.Number("0") || reference["cases_sha256"] != nativePin.CasesSHA256 || reference["observer_sha256"] != nativePin.ObserverSHA256 || reference["manifest_sha256"] != digest(manifest) || reference["php_image"] != nativePin.PHPImage {
		t.Fatal("retained record does not bind the current original corpus/profile")
	}
	var recordedSources map[string]string
	encodedSources, _ := json.Marshal(reference["sources"])
	if err := json.Unmarshal(encodedSources, &recordedSources); err != nil || !reflect.DeepEqual(recordedSources, nativePin.Sources) {
		t.Fatal("retained source identity mismatch")
	}
	if native["schema"] != "gossip.parser-full-native-observation/v1" || native["php_version"] != "8.3.36" || native["pcre_version"] != "10.42 2022-12-12" || native["lc_ctype"] != "C" || native["int_size"] != json.Number("8") {
		t.Fatal("retained native environment/profile mismatch")
	}
	expected, ok := native["cases"].([]any)
	if !ok || len(expected) != len(cases) {
		t.Fatal("retained case count mismatch")
	}
	for i, item := range expected {
		row, ok := item.(map[string]any)
		if !ok || row["name"] != cases[i].Name || row["caller_unchanged"] != true || row["reuse_reset"] != true {
			t.Fatal("retained case order/invariants mismatch")
		}
	}
	if mode == "fresh" && os.Getenv("GOSSIP_PHP_RUNNER_IMAGE") != nativePin.PHPImage {
		t.Fatal("fresh mode requires the exact retained PHP image; no implicit setup")
	}
	header, archive := filepath.Join(prefix, "include", "pcre2.h"), filepath.Join(prefix, "lib", "libpcre2-8.a")
	if digest(read(t, header)) != preg83ConsumerHeaderSHA || digest(read(t, archive)) != preg83ConsumerArchiveSHA {
		t.Fatal("PCRE2 differs from the measured static profile")
	}
	cc, sdk := preg83ConsumerToolchain(t)
	sdkSettings := read(t, filepath.Join(sdk, "SDKSettings.json"))
	var sdkInfo struct {
		Version string `json:"Version"`
	}
	if json.Unmarshal(sdkSettings, &sdkInfo) != nil || sdkInfo.Version != "15.5" {
		t.Fatal("requires measured SDK15.5")
	}
	goTool, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}

	evidence := map[string]any{
		"schema": "gossip.parser-full-generated-evidence/v1", "status": "failed", "application_acceptance": "not-established", "fresh_native_php": "not-run",
		"mode": mode, "gohp_commit": pin.GoHPCommit, "native_reference_sha256": digest(referenceBytes), "native_reference_reused": true,
		"native_reference_manifest_sha256": digest(manifest), "cases_sha256": digest(casesBytes), "observer_sha256": digest(observer), "harness_sha256": digest(read(t, "parser_full_generated_test.go")), "pin_sha256": digest(pinBytes), "closure_manifest_sha256": digest(closure), "wordpress_sources": nativePin.Sources,
		"php_image": nativePin.PHPImage, "go_version": runtime.Version(), "goos": runtime.GOOS, "goarch": runtime.GOARCH, "go_tool_sha256": digest(read(t, goTool)), "header_sha256": digest(read(t, header)), "static_archive_sha256": digest(read(t, archive)), "clang_sha256": digest(read(t, cc)), "sdk_version": sdkInfo.Version, "sdk_settings_sha256": digest(sdkSettings),
		"profile": map[string]any{"cgo": true, "tags": []string{"gohp_pcre2"}, "pcre2": "10.47_1", "limits": "match100000/depth10000/heap131072KiB/subject4096bytes", "jit": false, "utf_ucp": false},
		"phases":  map[string]string{"generation": "not-run", "candidate_build": "not-run", "candidate_execute": "not-run", "retained_comparison": "not-run", "fresh_native_comparison": "not-run"},
	}
	nativeCasesJSON, _ := json.Marshal(expected)
	evidence["native_typed_cases_sha256"] = digest(nativeCasesJSON)
	defer func() {
		if t.Failed() {
			evidence["status"] = "failed"
		}
		encoded, err := json.Marshal(evidence)
		if err != nil {
			t.Error(err)
			return
		}
		if output := os.Getenv("GOSSIP_GOHP_FULL_PARSER_EVIDENCE_OUT"); output != "" {
			if !filepath.IsAbs(output) {
				t.Error("evidence output requires an absolute fresh path")
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
				return
			}
		}
		if os.Getenv("GOSSIP_GOHP_FULL_PARSER_EVIDENCE_OUT") == "" {
			t.Log("FULL_PARSER_GENERATED_EVIDENCE " + string(encoded))
		} else {
			t.Logf("FULL_PARSER_GENERATED_EVIDENCE status=%v cases=%v calls=%v sha256=%s", evidence["status"], evidence["case_count"], evidence["parse_call_count"], digest(append(encoded, '\n')))
		}
	}()
	phases := evidence["phases"].(map[string]string)
	workspace := t.TempDir()
	source, input, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "input"), filepath.Join(workspace, "candidate")
	for _, path := range []string{source, input, candidate} {
		if err := os.MkdirAll(path, 0755); err != nil {
			t.Fatal(err)
		}
	}
	evidence["gohp_archive_sha256"] = exportGoHP(t, repository, pin.GoHPCommit, source)
	for path, data := range verified {
		write(t, filepath.Join(input, path), data)
	}
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "closure.json"), closure)
	generatedPath, reportPath := filepath.Join(candidate, "parser.go"), filepath.Join(candidate, "report.json")
	baseEnv := preg83ConsumerEnv(map[string]string{"GOTOOLCHAIN": "local", "GOPROXY": "off", "GOSUMDB": "off", "GOWORK": "off", "GOFLAGS": ""})
	if reused := os.Getenv("GOSSIP_GOHP_FULL_PARSER_GENERATED_DIR"); reused != "" {
		if !filepath.IsAbs(reused) {
			t.Fatal("reused generation directory must be absolute")
		}
		generated, report := read(t, filepath.Join(reused, "parser.go")), read(t, filepath.Join(reused, "report.json"))
		if digest(generated) != pin.GeneratedSHA256 || digest(report) != pin.ReusableReportSHA256 {
			t.Fatal("reused generation artifacts differ from recorded proof")
		}
		check := exec.Command("git", "-C", repository, "diff", "--exit-code", pin.GenerationCommit, pin.GoHPCommit, "--", "cmd", "internal", "runtime", "go.mod", "go.sum")
		if out, err := check.CombinedOutput(); err != nil {
			t.Fatalf("generation/runtime implementation changed: %v %s", err, out)
		}
		write(t, generatedPath, generated)
		write(t, reportPath, report)
		phases["generation"] = "reused-checked-artifact"
		evidence["generation_commit"] = pin.GenerationCommit
	} else {
		result, err := preg83ConsumerRun(t, "full_generation", 60*time.Second, source, preg83ConsumerWithEnv(baseEnv, map[string]string{"CGO_ENABLED": "0"}), goTool, "run", "./cmd/gohp", "-manifest", filepath.Join(input, "closure.json"), "-source-root", input, "-entry", "wp-includes/class-wp-block-parser.php", "-php-version", "8.3", "-experimental-pcre2", "-out", generatedPath, "-report", reportPath)
		if err != nil || len(result.stdout) != 0 || len(result.stderr) != 0 {
			t.Fatalf("full generation failed: %v stdout=%s stderr=%s", err, result.stdout, result.stderr)
		}
		phases["generation"] = "pass"
		evidence["generation_commit"] = pin.GoHPCommit
	}
	generated, reportBytes := read(t, generatedPath), read(t, reportPath)
	var report map[string]any
	numericBundleDecode(t, reportBytes, &report)
	if digest(generated) != pin.GeneratedSHA256 || report["schema"] != "gohp.build/v0.2" || report["status"] != "generated" || report["php_target"] != "8.3" || report["regex_profile"] != "pcre2-php83-v1" || report["manifest_sha256"] != pin.ClosureSHA256 || report["generated_sha256"] != digest(generated) || report["entry"] != "wp-includes/class-wp-block-parser.php" {
		t.Fatal("complete generated source/report identity mismatch")
	}
	evidence["generated_sha256"], evidence["report_sha256"] = digest(generated), digest(reportBytes)
	observerPath, binary := filepath.Join(candidate, "observer.go"), filepath.Join(candidate, "program")
	write(t, observerPath, observer)
	cgoEnv := preg83ConsumerWithEnv(baseEnv, map[string]string{"CGO_ENABLED": "1", "CC": cc, "SDKROOT": sdk, "CGO_CFLAGS": "-I" + filepath.Join(prefix, "include") + " -O2 -Wall -Wextra -Werror", "CGO_LDFLAGS": archive})
	build, err := preg83ConsumerRun(t, "full_candidate_build", 90*time.Second, source, cgoEnv, goTool, "build", "-trimpath", "-tags", "gohp_pcre2", "-o", binary, generatedPath, observerPath)
	redact := strings.NewReplacer(workspace, "<WORKSPACE>", repository, "<GOHP_SOURCE>", prefix, "<PCRE2_PREFIX>", sdk, "<SDKROOT>", cc, "<CC>")
	evidence["candidate_build_stderr_sha256"], evidence["candidate_build_stderr_sanitized"] = digest(build.stderr), redact.Replace(string(build.stderr))
	if err != nil {
		t.Fatalf("enabled full candidate build failed: %v stdout=%s stderr=%s", err, build.stdout, build.stderr)
	}
	phases["candidate_build"] = "pass"
	evidence["binary_sha256"] = digest(read(t, binary))
	run, err := preg83ConsumerRun(t, "full_candidate_execute", 30*time.Second, "", []string{"PATH="}, binary, filepath.Join(input, "cases.json"))
	evidence["candidate_stdout_sha256"], evidence["candidate_stdout_bytes"], evidence["candidate_stderr_sha256"], evidence["candidate_stderr_bytes"] = digest(run.stdout), len(run.stdout), digest(run.stderr), len(run.stderr)
	if err != nil || len(run.stderr) != 0 {
		evidence["candidate_stderr_sanitized"] = redact.Replace(string(run.stderr))
		t.Fatalf("original generated parse execution failed: %v stderr=%s", err, run.stderr)
	}
	phases["candidate_execute"] = "pass"
	evidence["candidate_exit"] = 0
	var actual map[string]any
	numericBundleDecode(t, run.stdout, &actual)
	if actual["schema"] != "gossip.parser-full-generated-observation/v1" {
		t.Fatal("generated observer schema mismatch")
	}
	actualCases, ok := actual["cases"].([]any)
	if !ok || len(actualCases) != len(expected) {
		t.Fatal("generated original case count mismatch")
	}
	evidence["generated_observation"] = actual
	mismatches := []string{}
	for i, want := range expected {
		if !reflect.DeepEqual(actualCases[i], want) {
			mismatches = append(mismatches, cases[i].Name)
		}
	}
	evidence["case_count"], evidence["parse_call_count"], evidence["mismatches"] = len(cases), 2*len(cases), mismatches
	canonical, _ := json.Marshal(actualCases)
	evidence["typed_cases_sha256"] = digest(canonical)
	if len(mismatches) != 0 {
		t.Fatalf("original generated parser differs from retained PHP: %v", mismatches)
	}
	phases["retained_comparison"] = "pass"
	evidence["status"] = "retained-comparison-pass"
	if mode == "fresh" {
		write(t, filepath.Join(input, "observer.php"), phpObserver)
		stdout, stderr := runPHPOracle(t, "full_parser_generated_native", nativePin.PHPImage, input, "/input/observer.php")
		evidence["fresh_native_stdout_sha256"], evidence["fresh_native_stdout_bytes"], evidence["fresh_native_stderr_bytes"] = digest(stdout), len(stdout), len(stderr)
		if len(stderr) != 0 {
			t.Fatal("fresh native diagnostics are not quiet")
		}
		var fresh map[string]any
		numericBundleDecode(t, stdout, &fresh)
		if fresh["schema"] != native["schema"] || fresh["php_version"] != native["php_version"] || fresh["pcre_version"] != native["pcre_version"] || !reflect.DeepEqual(fresh["profile"], native["profile"]) || !reflect.DeepEqual(fresh["cases"], actualCases) {
			t.Fatal("fresh pinned PHP and original generated parser mismatch")
		}
		evidence["fresh_native_php"] = "pass"
		evidence["status"] = "fresh-comparison-pass"
		phases["fresh_native_comparison"] = "pass"
	}
	evidence["scope"] = "Complete unchanged original Parser/Block/Frame, twelve fixed cases/twenty-four actual parse calls, typed ordered raw-byte values/state/errors/caller/reuse. Retained and fresh comparisons are distinct; corpus boundary review and later application gates remain open. No PHP8.4 or production claim."
	evidence["commands"] = map[string]string{"build": "go build -trimpath -tags gohp_pcre2 -o <PROGRAM> <GENERATED> <OBSERVER>", "execute": "<PROGRAM> <FIXED_CASES>", "generation": fmt.Sprint(phases["generation"])}
}
