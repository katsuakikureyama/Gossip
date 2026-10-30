package gohp_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"
)

// PASS means a native reference and a fail-closed admission diagnostic were
// captured. It explicitly does not mean generated hooks or bootstrap work.
func TestOriginalWordPressHooksBaseline(t *testing.T) {
	if os.Getenv("GOSSIP_GOHP_HOOKS_BASELINE") == "" {
		t.Skip("set GOSSIP_GOHP_HOOKS_BASELINE=1 with explicit source/image/fresh evidence paths")
	}
	if os.Getenv("GOSSIP_GOHP_HOOKS_BASELINE") != "1" {
		t.Fatal("hooks baseline opt-in must be 1")
	}
	repository, output := os.Getenv("GOSSIP_GOHP_HOOKS_SOURCE"), os.Getenv("GOSSIP_GOHP_HOOKS_EVIDENCE_OUT")
	if !filepath.IsAbs(repository) || !filepath.IsAbs(output) || os.Getuid() == 0 {
		t.Fatal("requires absolute paths and a non-root user")
	}
	if _, err := os.Lstat(output); !os.IsNotExist(err) {
		t.Fatal("evidence output must be a fresh unused path")
	}
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" || runtime.Version() != "go1.23.0" {
		t.Fatal("measured toolchain profile requires Darwin ARM64 Go1.23.0")
	}
	pinBytes := read(t, "hooks-original-baseline.json")
	var pin struct {
		Schema                  string `json:"schema"`
		GoHPCommit              string `json:"gohp_commit"`
		WordPressCommit         string `json:"wordpress_commit"`
		PHPImage                string `json:"php_image"`
		PHPTarget               string `json:"php_target"`
		Entry                   string `json:"entry"`
		ManifestSHA256          string `json:"closure_manifest_sha256"`
		CasesSHA256             string `json:"cases_sha256"`
		ObserverSHA256          string `json:"observer_sha256"`
		ExpectedSourceAdmission string `json:"expected_source_admission"`
	}
	numericBundleDecode(t, pinBytes, &pin)
	if pin.Schema != "gossip.hooks-original-baseline/v1" || pin.GoHPCommit != "b2d1482268a7710c515c07bf919af3d40b09d17c" || pin.WordPressCommit != "54106f6c545fc8b881f6b39b7963002bece6c28c" || pin.PHPTarget != "8.3" || pin.Entry != "wp-includes/plugin.php" || pin.ExpectedSourceAdmission != "rejected" || pin.PHPImage != "php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd" || os.Getenv("GOSSIP_PHP_RUNNER_IMAGE") != pin.PHPImage {
		t.Fatal("fixed hooks baseline identity drift")
	}
	manifest, casesBytes, observer := read(t, "hooks-source-closure.json"), read(t, "hooks-original-cases.json"), read(t, "hooks-original-observer.php")
	if digest(manifest) != pin.ManifestSHA256 || digest(casesBytes) != pin.CasesSHA256 || digest(observer) != pin.ObserverSHA256 {
		t.Fatal("hooks inputs/observer drift")
	}
	var closure struct {
		Schema   string `json:"schema"`
		ID       string `json:"id"`
		Upstream struct {
			Repository string `json:"repository"`
			Revision   string `json:"revision"`
			Version    string `json:"version"`
		} `json:"upstream"`
		Selection  map[string]string `json:"selection"`
		PHPTargets []string          `json:"php_targets"`
		Files      []struct {
			Path   string `json:"path"`
			Role   string `json:"role"`
			Bytes  int    `json:"bytes"`
			SHA256 string `json:"sha256"`
		} `json:"files"`
	}
	numericBundleDecode(t, manifest, &closure)
	if closure.Schema != "gohp.corpus.manifest/v0.1" || closure.Upstream.Revision != pin.WordPressCommit || len(closure.Files) != 5 || !reflect.DeepEqual(closure.PHPTargets, []string{"8.3"}) {
		t.Fatal("hooks whole-source closure identity drift")
	}
	verified, sources := map[string][]byte{}, map[string]string{}
	for _, file := range closure.Files {
		data := read(t, filepath.Join("../../..", file.Path))
		if len(data) != file.Bytes || digest(data) != file.SHA256 {
			t.Fatalf("original source drift: %s", file.Path)
		}
		verified[file.Path], sources[file.Path] = data, file.SHA256
	}
	var cases []string
	numericBundleDecode(t, casesBytes, &cases)
	expectedCases := []string{"no_callbacks", "priority_order", "accepted_arguments", "callback_identity_removal", "actions", "reentrant_filter", "mutation_during_iteration", "sentinel_identity"}
	if !reflect.DeepEqual(cases, expectedCases) {
		t.Fatal("hooks workflow identity/order drift")
	}
	goTool, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}
	evidence := map[string]any{
		"schema": "gossip.hooks-original-baseline-evidence/v1", "status": "failed", "application_acceptance": "not-established", "go_equivalence": "not-run",
		"gohp_commit": pin.GoHPCommit, "wordpress_commit": pin.WordPressCommit, "wordpress_sources": sources,
		"pin_sha256": digest(pinBytes), "closure_manifest_sha256": digest(manifest), "cases_sha256": digest(casesBytes), "observer_sha256": digest(observer), "harness_sha256": digest(read(t, "hooks_original_baseline_test.go")),
		"php_image": pin.PHPImage, "php_target": pin.PHPTarget, "entry": pin.Entry,
		"go_version": runtime.Version(), "goos": runtime.GOOS, "goarch": runtime.GOARCH, "go_tool_sha256": digest(read(t, goTool)),
		"phases": map[string]string{"native_reference": "not-run", "whole_source_admission": "not-run", "generated_build": "not-run", "generated_execute": "not-run", "differential_comparison": "not-run"},
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
		t.Logf("HOOKS_BASELINE_EVIDENCE status=%v workflows=%v admission=%v sha256=%s", evidence["status"], evidence["workflow_count"], evidence["source_admission"], digest(append(encoded, '\n')))
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
	write(t, filepath.Join(input, "closure.json"), manifest)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "observer.php"), observer)
	stdout, stderr := runPHPOracle(t, "hooks_original_native", pin.PHPImage, input, "/input/observer.php")
	evidence["native_stdout_sha256"], evidence["native_stdout_bytes"], evidence["native_stderr_sha256"], evidence["native_stderr_bytes"] = digest(stdout), len(stdout), digest(stderr), len(stderr)
	if len(stderr) != 0 {
		t.Fatal("native hook observation diagnostics are not quiet")
	}
	var native map[string]any
	numericBundleDecode(t, stdout, &native)
	evidence["native_observation"] = native
	environment, ok := native["native_environment"].(map[string]any)
	expectedProfile := map[string]any{"pcre.jit": "0", "pcre.backtrack_limit": "100000", "pcre.recursion_limit": "10000", "memory_limit": "128M"}
	if !ok || native["schema"] != "gossip.hooks-original-native-observation/v1" || environment["php_version"] != "8.3.36" || environment["php_version_id"] != json.Number("80336") || environment["pcre_version"] != "10.42 2022-12-12" || environment["int_size"] != json.Number("8") || environment["lc_ctype"] != "C" || environment["php_os"] != "Linux" || environment["php_arch"] != "aarch64" || !reflect.DeepEqual(environment["profile"], expectedProfile) {
		t.Fatal("native environment/profile drift")
	}
	rows, ok := native["workflows"].([]any)
	if !ok || len(rows) != len(cases) || native["workflow_count"] != json.Number("8") {
		t.Fatal("native workflow count mismatch")
	}
	for i, item := range rows {
		row, ok := item.(map[string]any)
		if !ok || row["name"] != cases[i] {
			t.Fatal("native workflow identity/order mismatch")
		}
		for _, field := range []string{"initial", "result", "events", "state_after"} {
			value, ok := row[field].(map[string]any)
			if !ok || value["type"] != "array" {
				t.Fatalf("missing typed workflow %s/%s", cases[i], field)
			}
		}
	}
	phases["native_reference"], evidence["native_exit"], evidence["workflow_count"] = "recorded", 0, len(rows)
	canonical, _ := json.Marshal(rows)
	evidence["native_typed_workflows_sha256"] = digest(canonical)
	generated, report := filepath.Join(candidate, "hooks.go"), filepath.Join(candidate, "report.json")
	env := preg83ConsumerEnv(map[string]string{"GOTOOLCHAIN": "local", "GOPROXY": "off", "GOSUMDB": "off", "GOWORK": "off", "GOFLAGS": "", "CGO_ENABLED": "0"})
	probe, probeErr := preg83ConsumerRun(t, "hooks_whole_source_admission", 60*time.Second, source, env, goTool, "run", "./cmd/gohp", "-manifest", filepath.Join(input, "closure.json"), "-source-root", input, "-entry", pin.Entry, "-php-version", pin.PHPTarget, "-out", generated, "-report", report)
	redact := strings.NewReplacer(workspace, "<WORKSPACE>", repository, "<GOHP_SOURCE>")
	evidence["admission_stdout_bytes"], evidence["admission_stdout_sha256"], evidence["admission_stderr_sha256"], evidence["admission_stderr_sanitized"] = len(probe.stdout), digest(probe.stdout), digest(probe.stderr), redact.Replace(string(probe.stderr))
	code := regexp.MustCompile(`PHP[0-9]{3}`).Find(probe.stderr)
	var diagnosticSource string
	for _, name := range []string{"wp-includes/plugin.php", "wp-includes/class-wp-hook.php", "wp-includes/class-wp-filter-sentinel.php"} {
		if strings.Contains(string(probe.stderr), name) {
			diagnosticSource = name
			break
		}
	}
	exit, ok := probeErr.(*exec.ExitError)
	if !ok || exit.ExitCode() != 1 || len(probe.stdout) != 0 || len(code) == 0 || diagnosticSource == "" {
		t.Fatalf("expected explicit original-source rejection, got %v stdout=%s stderr=%s", probeErr, probe.stdout, probe.stderr)
	}
	evidence["admission_exit"], evidence["diagnostic_code"], evidence["diagnostic_source"] = exit.ExitCode(), string(code), diagnosticSource
	for _, path := range []string{generated, report} {
		if _, err := os.Lstat(path); !os.IsNotExist(err) {
			t.Fatal("rejected admission must not publish generated artifacts")
		}
	}
	artifacts, err := os.ReadDir(candidate)
	if err != nil || len(artifacts) != 0 {
		t.Fatal("rejected admission left unexpected candidate artifacts")
	}
	phases["whole_source_admission"] = "rejected"
	evidence["source_admission"], evidence["generated_artifact_count"] = "rejected", 0
	evidence["status"] = "native-reference-recorded-source-rejected"
	evidence["scope"] = "Native-only eight-workflow original plugin API baseline and one complete unchanged-source admission rejection. No generated hook implementation/build/execution/equivalence, full bootstrap, G83-BOOT or application acceptance."
}
