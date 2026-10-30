package gohp_test

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"testing"
	"time"
)

type parserHostBoundariesPin struct {
	Schema                  string `json:"schema"`
	GoHPCommit              string `json:"gohp_commit"`
	RegexProfile            string `json:"regex_profile"`
	GeneratedSHA256         string `json:"generated_sha256"`
	ReportSHA256            string `json:"report_sha256"`
	RetainedEvidenceSHA256  string `json:"retained_evidence_sha256"`
	PreviousPinSHA256       string `json:"previous_pin_sha256"`
	ReferenceManifestSHA256 string `json:"reference_manifest_sha256"`
	ClosureManifestSHA256   string `json:"closure_manifest_sha256"`
	CasesSHA256             string `json:"cases_sha256"`
	ObserverSHA256          string `json:"observer_sha256"`
	DeadlineMilliseconds    int    `json:"deadline_milliseconds"`
}

type parserHostBoundaryInput struct {
	Name     string `json:"name"`
	Document string `json:"document_base64"`
	Expected string `json:"expected"`
}

// This is a host-boundary observation, not another PHP equivalence run.
// Generation and all earlier cases are reused as retained identities only.
func TestOriginalWordPressParserHostBoundaries(t *testing.T) {
	if os.Getenv("GOSSIP_GOHP_PARSER_HOST_BOUNDARIES") == "" {
		t.Skip("requires explicit host-boundary opt-in, source, retained artifacts, PCRE prefix and fresh evidence path")
	}
	if os.Getenv("GOSSIP_GOHP_PARSER_HOST_BOUNDARIES") != "1" {
		t.Fatal("host-boundary opt-in must be 1")
	}
	repository, prefix := os.Getenv("GOSSIP_GOHP_FULL_PARSER_SOURCE"), os.Getenv("GOSSIP_PCRE2_PREFIX")
	artifacts, output := os.Getenv("GOSSIP_GOHP_HOST_RETAINED_ARTIFACT_DIR"), os.Getenv("GOSSIP_GOHP_HOST_EVIDENCE_OUT")
	for _, path := range []string{repository, prefix, artifacts, output} {
		if !filepath.IsAbs(path) {
			t.Fatal("all source, retained artifact, prefix and output paths must be explicit and absolute")
		}
	}
	if os.Getuid() == 0 || runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" || runtime.Version() != "go1.23.0" {
		t.Fatal("requires the measured nonroot Darwin ARM64 Go1.23.0 profile")
	}
	if _, err := os.Lstat(output); !os.IsNotExist(err) {
		t.Fatal("evidence output must be a fresh unused path")
	}
	pinBytes := read(t, "parser-host-boundaries.json")
	var pin parserHostBoundariesPin
	numericBundleDecode(t, pinBytes, &pin)
	if pin.Schema != "gossip.parser-host-boundaries/v1" || pin.GoHPCommit != "e6a71a5f267658e8c575ffab5db6712246e543b2" || pin.RegexProfile != "pcre2-php83-match1-v1" || pin.DeadlineMilliseconds != 100 {
		t.Fatal("host-boundary pin/profile changed")
	}
	manifestBytes, closure := read(t, "parser-full-reference.json"), read(t, "parser-source-closure.json")
	var manifest parserMatch1ReferenceManifest
	numericBundleDecode(t, manifestBytes, &manifest)
	if digest(manifestBytes) != pin.ReferenceManifestSHA256 || digest(closure) != pin.ClosureManifestSHA256 || manifest.WordPressCommit != "54106f6c545fc8b881f6b39b7963002bece6c28c" || len(manifest.Sources) != 5 {
		t.Fatal("original source manifest or closure changed")
	}
	for path, expected := range manifest.Sources {
		if digest(read(t, filepath.Join("../../..", path))) != expected {
			t.Fatalf("original source drift: %s", path)
		}
	}
	oldPin := read(t, "parser-match1-generated.json")
	retained := read(t, "../../../docs/evidence/parser-match1-generated-2026-10-30.json")
	generated, reportBytes := read(t, filepath.Join(artifacts, "parser.go")), read(t, filepath.Join(artifacts, "report.json"))
	if digest(oldPin) != pin.PreviousPinSHA256 || digest(retained) != pin.RetainedEvidenceSHA256 || digest(generated) != pin.GeneratedSHA256 || digest(reportBytes) != pin.ReportSHA256 {
		t.Fatal("retained generation, report, pin or successful record identity changed")
	}
	var report, retainedRecord map[string]any
	numericBundleDecode(t, reportBytes, &report)
	numericBundleDecode(t, retained, &retainedRecord)
	if report["schema"] != "gohp.build/v0.2" || report["status"] != "generated" || report["php_target"] != "8.3" || report["regex_profile"] != pin.RegexProfile || report["manifest_sha256"] != digest(closure) || report["generated_sha256"] != digest(generated) || report["entry"] != "wp-includes/class-wp-block-parser.php" || retainedRecord["gohp_commit"] != pin.GoHPCommit || retainedRecord["regex_profile"] != pin.RegexProfile || retainedRecord["status"] != "retained-comparison-pass" || retainedRecord["generated_sha256"] != digest(generated) || retainedRecord["report_sha256"] != digest(reportBytes) {
		t.Fatal("retained source/report/evidence binding mismatch")
	}
	encodedSources, err := json.Marshal(retainedRecord["wordpress_sources"])
	if err != nil {
		t.Fatal(err)
	}
	var retainedSources map[string]string
	if err := json.Unmarshal(encodedSources, &retainedSources); err != nil || !reflect.DeepEqual(retainedSources, manifest.Sources) {
		t.Fatal("retained generated source map mismatch")
	}
	if !strings.Contains(string(generated), ".PregMatch83LocalMatch1(") || strings.Contains(string(generated), ".PregMatch83Local(") {
		t.Fatal("retained parser does not dispatch the named profile")
	}
	casesBytes, observer := read(t, "parser-host-boundaries-cases.json"), read(t, "parser-host-boundaries-observer.go.txt")
	if digest(casesBytes) != pin.CasesSHA256 || digest(observer) != pin.ObserverSHA256 {
		t.Fatal("fixed host-boundary cases or observer changed")
	}
	var cases []parserHostBoundaryInput
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 4 {
		t.Fatal("requires exactly four separate byte-bound cases")
	}
	documents := make([][]byte, len(cases))
	for i, item := range cases {
		documents[i], err = base64.StdEncoding.Strict().DecodeString(item.Document)
		want := 4097
		if i == 0 {
			want = 4096
		}
		if err != nil || len(documents[i]) != want || base64.StdEncoding.EncodeToString(documents[i]) != item.Document {
			t.Fatal("fixed boundary input identity changed")
		}
	}
	header, archive := filepath.Join(prefix, "include", "pcre2.h"), filepath.Join(prefix, "lib", "libpcre2-8.a")
	if digest(read(t, header)) != preg83ConsumerHeaderSHA || digest(read(t, archive)) != preg83ConsumerArchiveSHA {
		t.Fatal("PCRE2 differs from the measured static profile")
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
		"schema": "gossip.parser-host-boundaries-evidence/v1", "status": "failed", "application_acceptance": "not-established",
		"generation": "not-run-retained-artifact", "fresh_native_php": "not-run", "native_comparison": "not-applicable-host-boundary",
		"gohp_commit": pin.GoHPCommit, "regex_profile": pin.RegexProfile, "wordpress_commit": manifest.WordPressCommit, "wordpress_sources": manifest.Sources,
		"pin_sha256": digest(pinBytes), "cases_sha256": digest(casesBytes), "observer_sha256": digest(observer), "harness_sha256": digest(read(t, "parser_host_boundaries_test.go")),
		"retained_evidence_sha256": digest(retained), "previous_pin_sha256": digest(oldPin), "reference_manifest_sha256": digest(manifestBytes), "closure_manifest_sha256": digest(closure),
		"generated_sha256": digest(generated), "report_sha256": digest(reportBytes), "go_version": runtime.Version(), "goos": runtime.GOOS, "goarch": runtime.GOARCH,
		"go_tool_sha256": digest(read(t, goTool)), "header_sha256": digest(read(t, header)), "static_archive_sha256": digest(read(t, archive)), "clang_sha256": digest(read(t, cc)), "sdk_settings_sha256": digest(sdkSettings),
		"profile": map[string]any{"cgo": true, "tags": []string{"gohp_pcre2"}, "pcre2": "10.47_1", "sdk_version": "15.5", "limits": "match1/depth10000/heap131072KiB/subject4096bytes", "jit": false, "utf_ucp": false},
		"phases":  map[string]string{"retained_artifact_validation": "pass", "candidate_build": "not-run", "subject_guard": "not-run", "host_deadline": "not-run"},
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
		t.Logf("PARSER_HOST_BOUNDARIES_EVIDENCE status=%v sha256=%s", evidence["status"], digest(append(encoded, '\n')))
	}()
	workspace := t.TempDir()
	source, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "candidate")
	for _, path := range []string{source, candidate} {
		if err := os.MkdirAll(path, 0755); err != nil {
			t.Fatal(err)
		}
	}
	evidence["gohp_archive_sha256"] = exportGoHP(t, repository, pin.GoHPCommit, source)
	candidateSource, observerPath := filepath.Join(candidate, "parser.go"), filepath.Join(candidate, "observer.go")
	casesPath, binary := filepath.Join(candidate, "cases.json"), filepath.Join(candidate, "program")
	write(t, candidateSource, generated)
	write(t, observerPath, observer)
	write(t, casesPath, casesBytes)
	baseEnv := preg83ConsumerEnv(map[string]string{"GOTOOLCHAIN": "local", "GOPROXY": "off", "GOSUMDB": "off", "GOWORK": "off", "GOFLAGS": ""})
	cgoEnv := preg83ConsumerWithEnv(baseEnv, map[string]string{"CGO_ENABLED": "1", "CC": cc, "SDKROOT": sdk, "CGO_CFLAGS": "-I" + filepath.Join(prefix, "include") + " -O2 -Wall -Wextra -Werror", "CGO_LDFLAGS": archive})
	build, err := preg83ConsumerRun(t, "parser_host_candidate_build", 60*time.Second, source, cgoEnv, goTool, "build", "-trimpath", "-tags", "gohp_pcre2", "-o", binary, candidateSource, observerPath)
	redact := strings.NewReplacer(workspace, "<WORKSPACE>", repository, "<GOHP_SOURCE>", prefix, "<PCRE2_PREFIX>", sdk, "<SDKROOT>", cc, "<CC>")
	evidence["candidate_build_stderr_sha256"], evidence["candidate_build_stderr_sanitized"] = digest(build.stderr), redact.Replace(string(build.stderr))
	if err != nil || len(build.stdout) != 0 {
		t.Fatalf("host candidate build: %v stdout=%s stderr=%s", err, build.stdout, build.stderr)
	}
	phases := evidence["phases"].(map[string]string)
	phases["candidate_build"] = "pass"
	evidence["binary_sha256"] = digest(read(t, binary))
	run, err := preg83ConsumerRun(t, "parser_subject_guard", 5*time.Second, "", []string{"PATH="}, binary, "subject-guard", casesPath)
	evidence["subject_stdout_sha256"], evidence["subject_stderr_sha256"], evidence["subject_stderr_bytes"] = digest(run.stdout), digest(run.stderr), len(run.stderr)
	if err != nil || len(run.stderr) != 0 {
		evidence["subject_stderr_sanitized"] = redact.Replace(string(run.stderr))
		t.Fatalf("subject guard execution: %v stderr=%s", err, run.stderr)
	}
	var actual map[string]any
	numericBundleDecode(t, run.stdout, &actual)
	parserHostValidateObservation(t, actual, cases, documents)
	evidence["subject_observation"], evidence["subject_exit"] = actual, 0
	phases["subject_guard"] = "pass"
	deadline, err := parserHostDeadline(binary, casesPath, pin.DeadlineMilliseconds, digest(documents[0]))
	evidence["deadline_observation"] = deadline
	if err != nil {
		t.Fatal(err)
	}
	phases["host_deadline"] = "pass"
	evidence["status"] = "host-boundaries-observed"
	evidence["commands"] = map[string]string{"build": "go build -trimpath -tags gohp_pcre2 -o <PROGRAM> <RETAINED_PARSER_COPY> <NEW_OBSERVER>", "subject_guard": "<PROGRAM> subject-guard <FIXED_HOST_CASES>", "host_deadline": "<PROGRAM> deadline-workload <FIXED_HOST_CASES>; arm100ms after completed-original-call readiness; cancel and Wait"}
}

func parserHostValidateObservation(t *testing.T, actual map[string]any, cases []parserHostBoundaryInput, documents [][]byte) {
	t.Helper()
	if actual["schema"] != "gossip.parser-host-boundaries-observation/v1" || actual["original_parse_call_count"] != json.Number("8") || actual["adapter_seed_count"] != json.Number("4") {
		t.Fatal("host observation schema/count mismatch")
	}
	rows, ok := actual["cases"].([]any)
	if !ok || len(rows) != 4 {
		t.Fatal("host observation case count mismatch")
	}
	for i, item := range rows {
		row, ok := item.(map[string]any)
		if !ok || row["name"] != cases[i].Name || row["expected"] != cases[i].Expected || row["input_sha256"] != digest(documents[i]) || row["input_bytes"] != json.Number(fmt.Sprint(len(documents[i]))) || row["caller_unchanged"] != true || row["seed_pcre_error_code"] != json.Number("2") || row["seed_json_error_code"] != json.Number("4") || row["after_json_error_code"] != json.Number("4") || row["reset_json_error_code"] != json.Number("4") || row["reset_pcre_error_code"] != json.Number("0") {
			t.Fatalf("host observation row %d input/error contract changed", i)
		}
		if !reflect.DeepEqual(row["before"], map[string]any{"document": nil, "offset": nil, "output": nil, "stack": nil}) {
			t.Fatal("constructor null property state changed")
		}
		count := json.Number("0")
		if i == 0 {
			count = "1"
			if row["returned"] != true || row["panic_type"] != "" || row["panic_message"] != "" || row["freeform_inner_html_base64"] != cases[i].Document || (row["after_pcre_error_code"] != json.Number("0") && row["after_pcre_error_code"] != json.Number("2")) {
				t.Fatal("bounded original call did not complete with unchanged raw bytes")
			}
		} else {
			message, ok := row["panic_message"].(string)
			if !ok || row["returned"] != false || row["panic_type"] == "" || !strings.Contains(message, "subject size or byte offset outside bounds") || row["freeform_inner_html_base64"] != "" || row["after_pcre_error_code"] != json.Number("2") {
				t.Fatal("oversized call was not an explicit unsupported boundary with preserved PCRE state")
			}
		}
		if !reflect.DeepEqual(row["after"], map[string]any{"document_base64": cases[i].Document, "offset": json.Number("0"), "output_count": count, "stack_count": json.Number("0")}) || !reflect.DeepEqual(row["empty_reset"], map[string]any{"document_base64": "", "offset": json.Number("0"), "output_count": json.Number("0"), "stack_count": json.Number("0")}) {
			t.Fatal("partial parser state or following empty reset changed")
		}
	}
}

// The deadline is armed only after an original parse completes. The new
// workload repeatedly executes original parse calls; it does not stall stdin,
// sleep or pretend that a single PCRE2 call has been interrupted.
func parserHostDeadline(binary, casesPath string, milliseconds int, inputHash string) (map[string]any, error) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(context.Canceled)
	startup := time.AfterFunc(5*time.Second, func() { cancel(fmt.Errorf("readiness timeout")) })
	defer startup.Stop()
	cmd := exec.CommandContext(ctx, binary, "deadline-workload", casesPath)
	cmd.Env = []string{"PATH="}
	cmd.WaitDelay = time.Second
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	defer func() {
		cancel(context.Canceled)
		if cmd.ProcessState == nil {
			_ = cmd.Wait()
		}
	}()
	reader := bufio.NewReader(io.LimitReader(pipe, 8193))
	line, err := reader.ReadBytes('\n')
	if err != nil || len(line) > 1024 {
		return nil, fmt.Errorf("deadline workload readiness: %v", err)
	}
	var ready map[string]any
	decoder := json.NewDecoder(bytes.NewReader(line))
	decoder.UseNumber()
	if err := decoder.Decode(&ready); err != nil || ready["schema"] != "gossip.parser-host-deadline-ready/v1" || ready["first_parse_completed"] != true || ready["input_bytes"] != json.Number("4096") || ready["input_sha256"] != inputHash || (ready["pcre_error_code"] != json.Number("0") && ready["pcre_error_code"] != json.Number("2")) {
		return nil, fmt.Errorf("deadline workload invalid readiness")
	}
	if !startup.Stop() && ctx.Err() != nil {
		return nil, fmt.Errorf("deadline workload readiness exceeded startup guard")
	}
	started := time.Now()
	deadline := time.AfterFunc(time.Duration(milliseconds)*time.Millisecond, func() { cancel(context.DeadlineExceeded) })
	defer deadline.Stop()
	trailing, readErr := io.ReadAll(reader)
	waitErr := cmd.Wait()
	observation := map[string]any{"status": "host-timeout", "readiness": ready, "deadline_milliseconds": milliseconds, "elapsed_nanoseconds": time.Since(started).Nanoseconds(), "wait_completed": cmd.ProcessState != nil, "complete_parse_observation": false, "php_error_result": "not-produced", "single_stalled_pcre_interruption": "not-established", "stdout_sha256": digest(append(append([]byte{}, line...), trailing...)), "stderr_sha256": digest(stderr.Bytes()), "stderr_bytes": stderr.Len()}
	var exit *exec.ExitError
	if !errors.Is(context.Cause(ctx), context.DeadlineExceeded) || !errors.As(waitErr, &exit) || readErr != nil || len(trailing) != 0 || stderr.Len() != 0 || cmd.ProcessState == nil {
		return observation, fmt.Errorf("host deadline did not cancel/reap cleanly: cause=%v wait=%v read=%v trailing=%d stderr=%d", context.Cause(ctx), waitErr, readErr, len(trailing), stderr.Len())
	}
	status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus)
	if !ok || !status.Signaled() || status.Signal() != syscall.SIGKILL {
		return observation, fmt.Errorf("host deadline exit is not SIGKILL")
	}
	observation["termination"], observation["exit_code"] = "SIGKILL", cmd.ProcessState.ExitCode()
	return observation, nil
}
