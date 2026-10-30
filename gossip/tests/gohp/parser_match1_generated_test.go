package gohp_test

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

type parserMatch1GeneratedPin struct {
	Schema                  string `json:"schema"`
	GoHPCommit              string `json:"gohp_commit"`
	RegexProfile            string `json:"regex_profile"`
	NativeReferenceSHA256   string `json:"native_reference_sha256"`
	ReferenceManifestSHA256 string `json:"reference_manifest_sha256"`
	CasesSHA256             string `json:"cases_sha256"`
	GoObserverSHA256        string `json:"go_observer_sha256"`
	PHPObserverSHA256       string `json:"php_observer_sha256"`
	ClosureManifestSHA256   string `json:"closure_manifest_sha256"`
}

type parserMatch1ReferenceManifest struct {
	Schema          string            `json:"schema"`
	WordPressCommit string            `json:"wordpress_commit"`
	PHPImage        string            `json:"php_image"`
	PHPTarget       string            `json:"php_target"`
	CasesSHA256     string            `json:"cases_sha256"`
	ObserverSHA256  string            `json:"observer_sha256"`
	Sources         map[string]string `json:"sources"`
}

// TestOriginalWordPressParserMatch1Generated reuses the separately retained
// PHP match-limit-one record. It never invokes PHP or Docker: only a fresh
// archive-exported GoHP compilation and generated candidate are executed.
func TestOriginalWordPressParserMatch1Generated(t *testing.T) {
	if os.Getenv("GOSSIP_GOHP_PARSER_MATCH1_GENERATED") == "" {
		t.Skip("set GOSSIP_GOHP_PARSER_MATCH1_GENERATED=1 plus explicit source, PCRE2 prefix, evidence, and artifact paths")
	}
	if os.Getenv("GOSSIP_GOHP_PARSER_MATCH1_GENERATED") != "1" {
		t.Fatal("parser match1 generated opt-in must be 1")
	}
	repository, prefix := os.Getenv("GOSSIP_GOHP_FULL_PARSER_SOURCE"), os.Getenv("GOSSIP_PCRE2_PREFIX")
	output, artifactDir := os.Getenv("GOSSIP_GOHP_MATCH1_EVIDENCE_OUT"), os.Getenv("GOSSIP_GOHP_MATCH1_ARTIFACT_DIR")
	if !filepath.IsAbs(repository) || !filepath.IsAbs(prefix) || !filepath.IsAbs(output) || !filepath.IsAbs(artifactDir) || os.Getuid() == 0 {
		t.Fatal("requires explicit absolute source, prefix, evidence, and artifact paths and a non-root user")
	}
	if _, err := os.Lstat(output); !os.IsNotExist(err) {
		t.Fatal("evidence output must be a fresh unused path")
	}
	if _, err := os.Lstat(artifactDir); !os.IsNotExist(err) {
		t.Fatal("artifact directory must be a fresh unused path")
	}
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" || runtime.Version() != "go1.23.0" {
		t.Fatal("measured profile requires Darwin ARM64 Go1.23.0")
	}

	pinBytes := read(t, "parser-match1-generated.json")
	var pin parserMatch1GeneratedPin
	numericBundleDecode(t, pinBytes, &pin)
	if pin.Schema != "gossip.parser-match1-generated/v1" || len(pin.GoHPCommit) != 40 || strings.Trim(pin.GoHPCommit, "0123456789abcdef") != "" || pin.RegexProfile != "pcre2-php83-match1-v1" {
		t.Fatal("match1 generated pin identity changed")
	}

	manifestBytes := read(t, "parser-full-reference.json")
	var manifest parserMatch1ReferenceManifest
	numericBundleDecode(t, manifestBytes, &manifest)
	if digest(manifestBytes) != pin.ReferenceManifestSHA256 || manifest.Schema != "gossip.parser-full-reference/v1" || manifest.WordPressCommit != "54106f6c545fc8b881f6b39b7963002bece6c28c" || manifest.PHPTarget != "8.3" || len(manifest.Sources) != 5 {
		t.Fatal("original reference manifest changed")
	}
	verified := map[string][]byte{}
	for path, hash := range manifest.Sources {
		data := read(t, filepath.Join("../../..", path))
		if digest(data) != hash {
			t.Fatalf("original source drift: %s", path)
		}
		verified[path] = data
	}

	casesBytes := read(t, "parser-resource-boundaries-cases.json")
	goObserver := read(t, "parser-full-generated-observer.go.txt")
	phpObserver := read(t, "parser-resource-match1-observer.php")
	closureBytes := read(t, "parser-source-closure.json")
	if digest(casesBytes) != pin.CasesSHA256 || digest(goObserver) != pin.GoObserverSHA256 || digest(phpObserver) != pin.PHPObserverSHA256 || digest(closureBytes) != pin.ClosureManifestSHA256 {
		t.Fatal("match1 corpus, observers, or closure identity drifted")
	}
	var cases []struct {
		Name     string `json:"name"`
		Document string `json:"document_base64"`
	}
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 12 {
		t.Fatal("requires all twelve unchanged resource-boundary cases")
	}
	seen := map[string]bool{}
	for _, item := range cases {
		document, err := base64.StdEncoding.Strict().DecodeString(item.Document)
		if err != nil || base64.StdEncoding.EncodeToString(document) != item.Document || len(document) > 4096 || item.Name == "" || seen[item.Name] {
			t.Fatal("invalid fixed match1 case")
		}
		seen[item.Name] = true
	}

	nativeReferenceBytes := read(t, "../../../docs/evidence/parser-resource-match1-2026-10-30.json")
	const nativeReferenceSHA256 = "59d3e4cfe3553677f743b93d571966ffd538be9cbe1a0610c0788cc8488df60c"
	if digest(nativeReferenceBytes) != nativeReferenceSHA256 || pin.NativeReferenceSHA256 != nativeReferenceSHA256 {
		t.Fatal("retained native match1 reference identity changed")
	}
	var nativeRecord map[string]any
	numericBundleDecode(t, nativeReferenceBytes, &nativeRecord)
	if nativeRecord["schema"] != "gossip.parser-resource-match1-evidence/v1" || nativeRecord["status"] != "native-resource-reference-recorded" || nativeRecord["profile"] != "php83-original-parser-match1-v1" || nativeRecord["reference_manifest_sha256"] != digest(manifestBytes) || nativeRecord["cases_sha256"] != digest(casesBytes) || nativeRecord["php_observer_sha256"] != digest(phpObserver) || nativeRecord["case_count"] != json.Number("12") || nativeRecord["parse_call_count"] != json.Number("24") {
		t.Fatal("retained native match1 record no longer binds the fixed corpus")
	}
	encodedSources, err := json.Marshal(nativeRecord["wordpress_sources"])
	if err != nil {
		t.Fatal(err)
	}
	var recordedSources map[string]string
	if err := json.Unmarshal(encodedSources, &recordedSources); err != nil || !reflect.DeepEqual(recordedSources, manifest.Sources) || nativeRecord["wordpress_commit"] != manifest.WordPressCommit || nativeRecord["php_image"] != manifest.PHPImage {
		t.Fatal("retained native match1 source identity mismatch")
	}
	native, ok := nativeRecord["native_observation"].(map[string]any)
	if !ok {
		t.Fatal("retained native match1 observation is missing")
	}
	expectedNativeProfile := map[string]any{"pcre.jit": "0", "pcre.backtrack_limit": "1", "pcre.recursion_limit": "10000", "memory_limit": "128M"}
	if native["schema"] != "gossip.parser-full-native-observation/v1" || native["php_version"] != "8.3.36" || native["php_version_id"] != json.Number("80336") || native["pcre_version"] != "10.42 2022-12-12" || native["int_size"] != json.Number("8") || native["lc_ctype"] != "C" || native["php_os"] != "Linux" || native["php_arch"] != "aarch64" || !reflect.DeepEqual(native["profile"], expectedNativeProfile) {
		t.Fatal("retained native PHP environment/profile changed")
	}
	nativeCases, ok := native["cases"].([]any)
	if !ok || len(nativeCases) != len(cases) {
		t.Fatal("retained native match1 case count mismatch")
	}
	checkRows := func(label string, rows []any) map[string]int {
		t.Helper()
		if len(rows) != len(cases) {
			t.Fatalf("%s case count = %d, want %d", label, len(rows), len(cases))
		}
		counts := map[string]int{}
		for i, item := range rows {
			row, ok := item.(map[string]any)
			if !ok || row["name"] != cases[i].Name || row["caller_unchanged"] != true || row["reuse_reset"] != true {
				t.Fatalf("%s row %d identity/caller/reuse mismatch", label, i)
			}
			input, ok := row["input"].(map[string]any)
			if !ok || input["document_base64"] != cases[i].Document {
				t.Fatalf("%s row %d input binding mismatch", label, i)
			}
			for _, pass := range []string{"first_parse", "reuse_parse"} {
				call, ok := row[pass].(map[string]any)
				if !ok {
					t.Fatalf("%s row %d %s missing", label, i, pass)
				}
				code, ok := call["pcre_error_code"].(json.Number)
				if !ok || (code != "0" && code != "2") {
					t.Fatalf("%s row %d %s invalid PCRE error", label, i, pass)
				}
				counts[string(code)]++
			}
		}
		if counts["0"] != 8 || counts["2"] != 16 {
			t.Fatalf("%s PCRE error counts = %#v, want 0:8 and 2:16", label, counts)
		}
		return counts
	}
	nativeCounts := checkRows("retained native", nativeCases)
	nativeCanonical, err := json.Marshal(nativeCases)
	if err != nil {
		t.Fatal(err)
	}
	const nativeTypedCasesSHA256 = "4f4d9fcd79d53704981a9ae8e1c51f1387971818d21fe0223e1092a89884a728"
	if digest(nativeCanonical) != nativeTypedCasesSHA256 || nativeRecord["native_typed_cases_sha256"] != nativeTypedCasesSHA256 {
		t.Fatal("retained native typed cases changed")
	}

	header, archive := filepath.Join(prefix, "include", "pcre2.h"), filepath.Join(prefix, "lib", "libpcre2-8.a")
	if digest(read(t, header)) != preg83ConsumerHeaderSHA || digest(read(t, archive)) != preg83ConsumerArchiveSHA {
		t.Fatal("PCRE2 differs from measured static profile")
	}
	cc, sdk := preg83ConsumerToolchain(t)
	sdkSettings := read(t, filepath.Join(sdk, "SDKSettings.json"))
	if digest(sdkSettings) != "58a133735f0a55a624a1703067059f6e78925e51725ec6e1f966072e142c9c42" {
		t.Fatal("SDK15.5 settings changed")
	}
	goTool, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}

	evidence := map[string]any{
		"schema": "gossip.parser-match1-generated-evidence/v1", "status": "failed", "application_acceptance": "not-established", "fresh_native_php": "not-run-retained-reference-only",
		"gohp_commit": pin.GoHPCommit, "regex_profile": pin.RegexProfile, "pin_sha256": digest(pinBytes),
		"native_reference_sha256": digest(nativeReferenceBytes), "native_reference_reused": true, "reference_manifest_sha256": digest(manifestBytes),
		"cases_sha256": digest(casesBytes), "go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "closure_manifest_sha256": digest(closureBytes), "harness_sha256": digest(read(t, "parser_match1_generated_test.go")),
		"wordpress_commit": manifest.WordPressCommit, "wordpress_sources": manifest.Sources, "php_image": manifest.PHPImage,
		"expected_native_cases": nativeCases, "native_typed_cases_sha256": digest(nativeCanonical), "native_pcre_error_counts": nativeCounts,
		"go_version": runtime.Version(), "goos": runtime.GOOS, "goarch": runtime.GOARCH, "go_tool_sha256": digest(read(t, goTool)),
		"header_sha256": digest(read(t, header)), "static_archive_sha256": digest(read(t, archive)), "clang_sha256": digest(read(t, cc)), "sdk_settings_sha256": digest(sdkSettings),
		"profile": map[string]any{"cgo": true, "tags": []string{"gohp_pcre2"}, "pcre2": "10.47_1", "sdk_version": "15.5", "limits": "match1/depth10000/heap131072KiB/subject4096bytes", "jit": false, "utf_ucp": false},
		"phases":  map[string]string{"generation": "not-run", "candidate_build": "not-run", "candidate_execute": "not-run", "retained_comparison": "not-run"},
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
			return
		}
		t.Logf("PARSER_MATCH1_GENERATED_EVIDENCE status=%v cases=%v calls=%v sha256=%s", evidence["status"], evidence["case_count"], evidence["parse_call_count"], digest(append(encoded, '\n')))
	}()

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
	write(t, filepath.Join(input, "closure.json"), closureBytes)

	if err := os.Mkdir(artifactDir, 0700); err != nil {
		t.Fatal(err)
	}
	generatedPath, reportPath := filepath.Join(artifactDir, "parser.go"), filepath.Join(artifactDir, "report.json")
	baseEnv := preg83ConsumerEnv(map[string]string{"GOTOOLCHAIN": "local", "GOPROXY": "off", "GOSUMDB": "off", "GOWORK": "off", "GOFLAGS": ""})
	generation, err := preg83ConsumerRun(t, "parser_match1_generation", 60*time.Second, source, preg83ConsumerWithEnv(baseEnv, map[string]string{"CGO_ENABLED": "0"}), goTool,
		"run", "./cmd/gohp", "-manifest", filepath.Join(input, "closure.json"), "-source-root", input, "-entry", "wp-includes/class-wp-block-parser.php", "-php-version", "8.3", "-experimental-pcre2", "-experimental-pcre2-profile", pin.RegexProfile, "-out", generatedPath, "-report", reportPath)
	if err != nil || len(generation.stdout) != 0 || len(generation.stderr) != 0 {
		t.Fatalf("match1 generation failed: %v stdout=%s stderr=%s", err, generation.stdout, generation.stderr)
	}
	generated, reportBytes := read(t, generatedPath), read(t, reportPath)
	var report map[string]any
	numericBundleDecode(t, reportBytes, &report)
	if report["schema"] != "gohp.build/v0.2" || report["status"] != "generated" || report["php_target"] != "8.3" || report["regex_profile"] != pin.RegexProfile || report["manifest_sha256"] != digest(closureBytes) || report["generated_sha256"] != digest(generated) || report["entry"] != "wp-includes/class-wp-block-parser.php" {
		t.Fatal("match1 generated report identity mismatch")
	}
	if !strings.Contains(string(generated), ".PregMatch83LocalMatch1(") || strings.Contains(string(generated), ".PregMatch83Local(") {
		t.Fatal("generated parser did not select the distinct match1 runtime entry")
	}
	entries, err := os.ReadDir(artifactDir)
	if err != nil || len(entries) != 2 || entries[0].Name() != "parser.go" || entries[1].Name() != "report.json" || entries[0].IsDir() || entries[1].IsDir() {
		t.Fatal("artifact directory must retain only generated parser.go and report.json")
	}
	phases := evidence["phases"].(map[string]string)
	phases["generation"] = "pass"
	evidence["generated_sha256"], evidence["report_sha256"] = digest(generated), digest(reportBytes)

	observerPath, binary := filepath.Join(candidate, "observer.go"), filepath.Join(candidate, "program")
	candidateSource := filepath.Join(candidate, "parser.go")
	write(t, candidateSource, generated)
	write(t, observerPath, goObserver)
	cgoEnv := preg83ConsumerWithEnv(baseEnv, map[string]string{"CGO_ENABLED": "1", "CC": cc, "SDKROOT": sdk, "CGO_CFLAGS": "-I" + filepath.Join(prefix, "include") + " -O2 -Wall -Wextra -Werror", "CGO_LDFLAGS": archive})
	build, err := preg83ConsumerRun(t, "parser_match1_candidate_build", 90*time.Second, source, cgoEnv, goTool, "build", "-trimpath", "-tags", "gohp_pcre2", "-o", binary, candidateSource, observerPath)
	redact := strings.NewReplacer(workspace, "<WORKSPACE>", repository, "<GOHP_SOURCE>", prefix, "<PCRE2_PREFIX>", sdk, "<SDKROOT>", cc, "<CC>")
	evidence["candidate_build_stderr_sha256"], evidence["candidate_build_stderr_sanitized"] = digest(build.stderr), redact.Replace(string(build.stderr))
	if err != nil || len(build.stdout) != 0 {
		t.Fatalf("match1 candidate build failed: %v stdout=%s stderr=%s", err, build.stdout, build.stderr)
	}
	phases["candidate_build"] = "pass"
	evidence["binary_sha256"] = digest(read(t, binary))

	run, err := preg83ConsumerRun(t, "parser_match1_candidate_execute", 30*time.Second, "", []string{"PATH="}, binary, filepath.Join(input, "cases.json"))
	evidence["candidate_stdout_sha256"], evidence["candidate_stdout_bytes"], evidence["candidate_stderr_sha256"], evidence["candidate_stderr_bytes"] = digest(run.stdout), len(run.stdout), digest(run.stderr), len(run.stderr)
	if err != nil || len(run.stderr) != 0 {
		evidence["candidate_stderr_sanitized"] = redact.Replace(string(run.stderr))
		t.Fatalf("match1 candidate execution failed: %v stderr=%s", err, run.stderr)
	}
	evidence["candidate_exit"] = 0
	var actual map[string]any
	numericBundleDecode(t, run.stdout, &actual)
	if actual["schema"] != "gossip.parser-full-generated-observation/v1" {
		t.Fatal("generated match1 observation schema changed")
	}
	actualCases, ok := actual["cases"].([]any)
	if !ok {
		t.Fatal("generated match1 cases are missing")
	}
	actualCounts := checkRows("generated", actualCases)
	actualCanonical, err := json.Marshal(actualCases)
	if err != nil {
		t.Fatal(err)
	}
	phases["candidate_execute"] = "pass"
	evidence["candidate_observation"], evidence["typed_cases_sha256"], evidence["candidate_pcre_error_counts"] = actual, digest(actualCanonical), actualCounts
	evidence["case_count"], evidence["parse_call_count"] = len(actualCases), 2*len(actualCases)
	if !reflect.DeepEqual(actualCases, nativeCases) {
		t.Fatal("original PHP retained match1 cases and generated match1 cases differ")
	}
	phases["retained_comparison"], evidence["status"], evidence["mismatch_count"] = "pass", "retained-comparison-pass", 0
	evidence["scope"] = "Complete unchanged original Parser/Block/Frame, twelve fixed bounded documents and twenty-four generated parse calls at Match=1/Depth=10000/Heap=131072KiB/subject=4096. Native PHP is the separately retained match1 reference, not freshly invoked. Typed ordered values/state/errors/caller/reuse are compared; no PHP8.4, production, or application-acceptance claim."
	evidence["commands"] = map[string]string{"generation": "go run ./cmd/gohp -manifest <CLOSURE> -source-root <INPUT> -entry wp-includes/class-wp-block-parser.php -php-version 8.3 -experimental-pcre2 -experimental-pcre2-profile pcre2-php83-match1-v1 -out <ARTIFACT>/parser.go -report <ARTIFACT>/report.json", "build": "go build -trimpath -tags gohp_pcre2 -o <PROGRAM> <ARTIFACT>/parser.go <OBSERVER>", "execute": "<PROGRAM> <FIXED_CASES>"}
}
