package gohp_test

import (
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

// A distinct corpus and observer identity: the earlier twelve-case series
// deliberately established a JSON baseline before every parse and is unchanged.
func TestOriginalWordPressParserCarryState(t *testing.T) {
	if os.Getenv("GOSSIP_GOHP_PARSER_CARRY_STATE") == "" {
		t.Skip("set GOSSIP_GOHP_PARSER_CARRY_STATE=1 and explicit source/prefix/generated-dir/image/evidence paths")
	}
	if os.Getenv("GOSSIP_GOHP_PARSER_CARRY_STATE") != "1" {
		t.Fatal("carry-state opt-in must be 1")
	}
	repository, prefix := os.Getenv("GOSSIP_GOHP_FULL_PARSER_SOURCE"), os.Getenv("GOSSIP_PCRE2_PREFIX")
	generatedDir, output := os.Getenv("GOSSIP_GOHP_FULL_PARSER_GENERATED_DIR"), os.Getenv("GOSSIP_GOHP_CARRY_EVIDENCE_OUT")
	if !filepath.IsAbs(repository) || !filepath.IsAbs(prefix) || !filepath.IsAbs(generatedDir) || !filepath.IsAbs(output) || os.Getuid() == 0 {
		t.Fatal("requires explicit absolute paths and a non-root user")
	}
	if _, err := os.Lstat(output); !os.IsNotExist(err) {
		t.Fatal("evidence output must be a fresh unused path")
	}
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" || runtime.Version() != "go1.23.0" {
		t.Fatal("measured profile requires Darwin ARM64 Go1.23.0")
	}
	pinBytes := read(t, "parser-carry-state.json")
	var pin struct {
		Schema                  string `json:"schema"`
		GeneratedPinSHA256      string `json:"generated_pin_sha256"`
		ReferenceManifestSHA256 string `json:"reference_manifest_sha256"`
		DocumentsSHA256         string `json:"documents_sha256"`
		SequencesSHA256         string `json:"sequences_sha256"`
		GoObserverSHA256        string `json:"go_observer_sha256"`
		PHPObserverSHA256       string `json:"php_observer_sha256"`
	}
	numericBundleDecode(t, pinBytes, &pin)
	generatedPinBytes, manifest := read(t, "parser-full-generated.json"), read(t, "parser-full-reference.json")
	if pin.Schema != "gossip.parser-carry-state/v1" || digest(generatedPinBytes) != pin.GeneratedPinSHA256 || digest(manifest) != pin.ReferenceManifestSHA256 {
		t.Fatal("carry-state series or existing pins drifted")
	}
	var generation fullParserGeneratedPin
	numericBundleDecode(t, generatedPinBytes, &generation)
	if generation.GoHPCommit != "b2d1482268a7710c515c07bf919af3d40b09d17c" || generation.GenerationCommit != "35524a636073f4abbc8bb69b9e36263e71ac62d8" || generation.ReferenceManifestSHA256 != digest(manifest) {
		t.Fatal("fixed compiler/generation identity changed")
	}
	var nativePin struct {
		Schema          string            `json:"schema"`
		WordPressCommit string            `json:"wordpress_commit"`
		PHPImage        string            `json:"php_image"`
		PHPTarget       string            `json:"php_target"`
		CasesSHA256     string            `json:"cases_sha256"`
		ObserverSHA256  string            `json:"observer_sha256"`
		Sources         map[string]string `json:"sources"`
	}
	numericBundleDecode(t, manifest, &nativePin)
	if nativePin.Schema != "gossip.parser-full-reference/v1" || nativePin.WordPressCommit != "54106f6c545fc8b881f6b39b7963002bece6c28c" || nativePin.PHPTarget != "8.3" || os.Getenv("GOSSIP_PHP_RUNNER_IMAGE") != nativePin.PHPImage {
		t.Fatal("requires unchanged original and exact fixed native PHP image")
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
		t.Fatal("requires unchanged full original closure")
	}
	documents, sequences := read(t, "parser-full-reference-cases.json"), read(t, "parser-carry-state-cases.json")
	goObserver, phpObserver := read(t, "parser-carry-state-observer.go.txt"), read(t, "parser-carry-state-observer.php")
	if digest(documents) != pin.DocumentsSHA256 || digest(documents) != nativePin.CasesSHA256 || digest(sequences) != pin.SequencesSHA256 || digest(goObserver) != pin.GoObserverSHA256 || digest(phpObserver) != pin.PHPObserverSHA256 || digest(read(t, "parser-source-closure.json")) != generation.ClosureSHA256 {
		t.Fatal("carry-state inputs/observers/source closure drifted")
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
		"schema": "gossip.parser-carry-state-evidence/v1", "status": "failed", "application_acceptance": "not-established",
		"gohp_commit": generation.GoHPCommit, "generation_commit": generation.GenerationCommit,
		"pin_sha256": digest(pinBytes), "generated_pin_sha256": digest(generatedPinBytes), "reference_manifest_sha256": digest(manifest),
		"documents_sha256": digest(documents), "sequences_sha256": digest(sequences), "go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "harness_sha256": digest(read(t, "parser_carry_state_test.go")),
		"wordpress_commit": nativePin.WordPressCommit, "wordpress_sources": nativePin.Sources, "closure_manifest_sha256": generation.ClosureSHA256, "php_image": nativePin.PHPImage,
		"go_version": runtime.Version(), "goos": runtime.GOOS, "goarch": runtime.GOARCH, "go_tool_sha256": digest(read(t, goTool)),
		"header_sha256": digest(read(t, header)), "static_archive_sha256": digest(read(t, archive)), "clang_sha256": digest(read(t, cc)), "sdk_settings_sha256": digest(sdkSettings),
		"profile": map[string]any{"cgo": true, "tags": []string{"gohp_pcre2"}, "pcre2": "10.47_1", "sdk_version": "15.5", "limits": "match100000/depth10000/heap131072KiB/subject4096bytes", "jit": false, "utf_ucp": false},
		"phases":  map[string]string{"generation": "not-run", "candidate_build": "not-run", "candidate_execute": "not-run", "native_execute": "not-run", "comparison": "not-run"},
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
		t.Logf("PARSER_CARRY_EVIDENCE status=%v scenarios=%v calls=%v sha256=%s", evidence["status"], evidence["scenario_count"], evidence["parse_call_count"], digest(append(encoded, '\n')))
	}()
	phases := evidence["phases"].(map[string]string)
	workspace := t.TempDir()
	source, input, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "input"), filepath.Join(workspace, "candidate")
	for _, path := range []string{source, input, candidate} {
		if err := os.MkdirAll(path, 0755); err != nil {
			t.Fatal(err)
		}
	}
	evidence["gohp_archive_sha256"] = exportGoHP(t, repository, generation.GoHPCommit, source)
	for path, data := range verified {
		write(t, filepath.Join(input, path), data)
	}
	write(t, filepath.Join(input, "documents.json"), documents)
	write(t, filepath.Join(input, "sequences.json"), sequences)
	write(t, filepath.Join(input, "observer.php"), phpObserver)
	generated, reportBytes := read(t, filepath.Join(generatedDir, "parser.go")), read(t, filepath.Join(generatedDir, "report.json"))
	if digest(generated) != generation.GeneratedSHA256 || digest(reportBytes) != generation.ReusableReportSHA256 {
		t.Fatal("reused generation artifacts changed")
	}
	check := exec.Command("git", "-C", repository, "diff", "--exit-code", generation.GenerationCommit, generation.GoHPCommit, "--", "cmd", "internal", "runtime", "go.mod", "go.sum")
	if out, err := check.CombinedOutput(); err != nil {
		t.Fatalf("generation/runtime changed: %v %s", err, out)
	}
	var report map[string]any
	numericBundleDecode(t, reportBytes, &report)
	if report["schema"] != "gohp.build/v0.2" || report["status"] != "generated" || report["php_target"] != "8.3" || report["regex_profile"] != "pcre2-php83-v1" || report["manifest_sha256"] != generation.ClosureSHA256 || report["generated_sha256"] != digest(generated) || report["entry"] != "wp-includes/class-wp-block-parser.php" {
		t.Fatal("original generated report identity mismatch")
	}
	phases["generation"] = "reused-checked-artifact"
	evidence["generated_sha256"], evidence["report_sha256"] = digest(generated), digest(reportBytes)
	generatedPath, observerPath, binary := filepath.Join(candidate, "parser.go"), filepath.Join(candidate, "observer.go"), filepath.Join(candidate, "program")
	write(t, generatedPath, generated)
	write(t, observerPath, goObserver)
	baseEnv := preg83ConsumerEnv(map[string]string{"GOTOOLCHAIN": "local", "GOPROXY": "off", "GOSUMDB": "off", "GOWORK": "off", "GOFLAGS": ""})
	cgoEnv := preg83ConsumerWithEnv(baseEnv, map[string]string{"CGO_ENABLED": "1", "CC": cc, "SDKROOT": sdk, "CGO_CFLAGS": "-I" + filepath.Join(prefix, "include") + " -O2 -Wall -Wextra -Werror", "CGO_LDFLAGS": archive})
	build, err := preg83ConsumerRun(t, "carry_candidate_build", 90*time.Second, source, cgoEnv, goTool, "build", "-trimpath", "-tags", "gohp_pcre2", "-o", binary, generatedPath, observerPath)
	redact := strings.NewReplacer(workspace, "<WORKSPACE>", repository, "<GOHP_SOURCE>", prefix, "<PCRE2_PREFIX>", sdk, "<SDKROOT>", cc, "<CC>")
	evidence["candidate_build_stderr_sha256"], evidence["candidate_build_stderr_sanitized"] = digest(build.stderr), redact.Replace(string(build.stderr))
	if err != nil || len(build.stdout) != 0 {
		t.Fatalf("candidate build failed: %v stdout=%s stderr=%s", err, build.stdout, build.stderr)
	}
	phases["candidate_build"] = "pass"
	evidence["binary_sha256"] = digest(read(t, binary))
	run, err := preg83ConsumerRun(t, "carry_candidate_execute", 30*time.Second, "", []string{"PATH="}, binary, filepath.Join(input, "documents.json"), filepath.Join(input, "sequences.json"))
	evidence["candidate_stdout_sha256"], evidence["candidate_stdout_bytes"], evidence["candidate_stderr_sha256"], evidence["candidate_stderr_bytes"] = digest(run.stdout), len(run.stdout), digest(run.stderr), len(run.stderr)
	if err != nil || len(run.stderr) != 0 {
		evidence["candidate_stderr_sanitized"] = redact.Replace(string(run.stderr))
		t.Fatalf("candidate execute failed: %v stderr=%s", err, run.stderr)
	}
	phases["candidate_execute"], evidence["candidate_exit"] = "pass", 0
	var actual map[string]any
	numericBundleDecode(t, run.stdout, &actual)
	evidence["candidate_observation"] = actual
	stdout, stderr := runPHPOracle(t, "parser_carry_state_native", nativePin.PHPImage, input, "/input/observer.php")
	evidence["native_stdout_sha256"], evidence["native_stdout_bytes"], evidence["native_stderr_sha256"], evidence["native_stderr_bytes"] = digest(stdout), len(stdout), digest(stderr), len(stderr)
	if len(stderr) != 0 {
		t.Fatal("native runtime diagnostics are not quiet")
	}
	phases["native_execute"], evidence["native_exit"] = "pass", 0
	var native map[string]any
	numericBundleDecode(t, stdout, &native)
	evidence["native_observation"] = native
	environment, ok := native["native_environment"].(map[string]any)
	expectedProfile := map[string]any{"pcre.jit": "0", "pcre.backtrack_limit": "100000", "pcre.recursion_limit": "10000", "memory_limit": "128M"}
	if !ok || environment["php_version"] != "8.3.36" || environment["php_version_id"] != json.Number("80336") || environment["pcre_version"] != "10.42 2022-12-12" || environment["int_size"] != json.Number("8") || environment["lc_ctype"] != "C" || environment["php_os"] != "Linux" || environment["php_arch"] != "aarch64" || !reflect.DeepEqual(environment["profile"], expectedProfile) {
		t.Fatal("native PHP environment/profile changed")
	}
	for _, observation := range []map[string]any{actual, native} {
		rows, ok := observation["scenarios"].([]any)
		if observation["schema"] != "gossip.parser-carry-state-observation/v1" || observation["parse_call_count"] != json.Number("20") || !ok || len(rows) != 4 {
			t.Fatal("carry observation schema/count mismatch")
		}
		for _, item := range rows {
			row, ok := item.(map[string]any)
			if !ok || row["reuse_reset"] != true {
				t.Fatal("parser reuse invariant failed")
			}
		}
	}
	canonical, _ := json.Marshal(actual["scenarios"])
	nativeCanonical, _ := json.Marshal(native["scenarios"])
	evidence["typed_scenarios_sha256"], evidence["native_typed_scenarios_sha256"] = digest(canonical), digest(nativeCanonical)
	evidence["scenario_count"], evidence["parse_call_count"] = 4, 20
	if !reflect.DeepEqual(actual["scenarios"], native["scenarios"]) {
		t.Fatal("original PHP and generated carry-state scenarios differ")
	}
	phases["comparison"], evidence["status"], evidence["mismatch_count"] = "pass", "fresh-comparison-pass", 0
	evidence["scope"] = "Four fixed carry-in sequences, twenty original parse calls per backend; JSON syntax/null state and independently seeded PCRE backtrack error. Does not establish original WordPress PCRE resource-error paths, arbitrary globals, JSON-domain coverage or whole application acceptance."
}
