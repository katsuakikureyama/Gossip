package gohp_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"
)

const (
	preg83ConsumerGoHPCommit = "d7650944802a33cfc2dab0006bb410374d828245"
	preg83ConsumerHeaderSHA  = "058462dc030e845ee627e19ac36347f561f5cc44d89eaf1c0010f2b363ab68d5"
	preg83ConsumerArchiveSHA = "ce40246eba1dfdd99e43142cfc33264610d7c965fab3062f386c68a6636a4a54"
	preg83ConsumerClangSHA   = "4c458256bcdf913774de1bb4a37768244d3b41f32ad55afb2dfec3432f46f4fe"
	preg83ConsumerRefSHA     = "5cbf560247afcd2c5f1462cab2f95f16bb822ac2d4d1756be16221cffdc6dc60"
	preg83ConsumerCasesSHA   = "ac92d6eb7cbf2fef591d73749e13c8110fe00934cb96a7e825dc7e5727ff83ef"
	preg83ConsumerPatternSHA = "3394500101eae41893b0978b3003f891dfafab54e7d7ebcbf820422753d8446e"
)

type preg83ConsumerCase struct {
	Name     string `json:"name"`
	Document string `json:"document_base64"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"backtrack_limit"`
}

type preg83ConsumerInputs struct {
	Manifest       []byte
	Reference      []byte
	CasesBytes     []byte
	DriverTemplate []byte
	Pattern        []byte
	Cases          []preg83ConsumerCase
	Selected       []preg83ConsumerCase
	Expected       []map[string]any
	SourceHashes   map[string]string
	PHPImage       string
}

// TestParserPreg83GeneratedContract is deliberately dependency-free. It pins
// the original source, retained PHP reference and the exact owned-driver seam;
// the native compiler/run proof below remains an explicit local opt-in.
func TestParserPreg83GeneratedContract(t *testing.T) {
	inputs := preg83ConsumerLoadInputs(t)
	t.Logf("verified_original_files=%d retained_cases=%d selected=%d excluded=backtrack_limit_one", len(inputs.SourceHashes), len(inputs.Cases), len(inputs.Selected))
}

// TestParserPreg83GeneratedConsumer builds the pinned GoHP compiler from a
// committed archive, then builds the generated program from that same module
// root. It compares only against the already-retained PHP 8.3 reference;
// it does not rerun PHP or claim a WordPress parser conversion.
func TestParserPreg83GeneratedConsumer(t *testing.T) {
	repository, prefix := os.Getenv("GOSSIP_PREG83_GOHP_SOURCE"), os.Getenv("GOSSIP_PCRE2_PREFIX")
	if repository == "" && prefix == "" {
		t.Skip("set GOSSIP_PREG83_GOHP_SOURCE and GOSSIP_PCRE2_PREFIX for the pinned local preg83 generated-consumer experiment")
	}
	if repository == "" || prefix == "" {
		t.Fatal("GOSSIP_PREG83_GOHP_SOURCE and GOSSIP_PCRE2_PREFIX must be supplied together")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" || runtime.Version() != "go1.23.0" {
		t.Fatalf("requires darwin/arm64 Go 1.23.0, got %s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	}
	if !filepath.IsAbs(repository) || !filepath.IsAbs(prefix) {
		t.Fatal("GoHP source and PCRE2 prefix must be absolute paths")
	}

	inputs := preg83ConsumerLoadInputs(t)
	header, archive := filepath.Join(prefix, "include", "pcre2.h"), filepath.Join(prefix, "lib", "libpcre2-8.a")
	if digest(read(t, header)) != preg83ConsumerHeaderSHA || digest(read(t, archive)) != preg83ConsumerArchiveSHA {
		t.Fatal("PCRE2 header/static archive differs from the measured profile")
	}
	resolvedCC, sdk := preg83ConsumerToolchain(t)
	sdkSettings := read(t, filepath.Join(sdk, "SDKSettings.json"))
	var sdkIdentity struct {
		Version string `json:"Version"`
	}
	if err := json.Unmarshal(sdkSettings, &sdkIdentity); err != nil || sdkIdentity.Version != "15.5" {
		t.Fatal("requires the measured macOS SDK 15.5")
	}
	goTool, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}
	goVersion, goVersionErr := preg83ConsumerRun(t, "go_version", 15*time.Second, "", nil, goTool, "version")
	if goVersionErr != nil || !strings.Contains(string(goVersion.stdout), "go1.23.0 ") {
		t.Fatalf("go tool version mismatch: stdout=%q stderr=%q err=%v", goVersion.stdout, goVersion.stderr, goVersionErr)
	}

	workspace := t.TempDir()
	source, input, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "input"), filepath.Join(workspace, "candidate")
	for _, path := range []string{source, input, candidate} {
		if err := os.MkdirAll(path, 0755); err != nil {
			t.Fatal(err)
		}
	}
	archiveHash := exportGoHP(t, repository, preg83ConsumerGoHPCommit, source)
	unit := strings.Replace(string(inputs.DriverTemplate), "@PATTERN_LITERAL@", preg83ConsumerPHPSingleQuoted(inputs.Pattern), 1)
	if strings.Count(unit, "@PATTERN_LITERAL@") != 0 {
		t.Fatal("driver placeholder remained after substitution")
	}
	unitPath, casesPath := filepath.Join(input, "unit.php"), filepath.Join(input, "cases.json")
	write(t, unitPath, []byte(unit))
	selectedJSON, err := json.Marshal(inputs.Selected)
	if err != nil {
		t.Fatal(err)
	}
	write(t, casesPath, append(selectedJSON, '\n'))
	observer := read(t, "parser-preg83-generated-observer.go.txt")
	observerPath := filepath.Join(candidate, "observer.go")
	write(t, observerPath, observer)

	baseEnv := preg83ConsumerEnv(map[string]string{
		"GOTOOLCHAIN": "local", "GOPROXY": "off", "GOSUMDB": "off", "GOWORK": "off", "GOFLAGS": "",
	})
	compilerPath := filepath.Join(candidate, "gohp")
	compilerBuild := preg83ConsumerWithEnv(baseEnv, map[string]string{"CGO_ENABLED": "0"})
	compiler, compilerErr := preg83ConsumerRun(t, "compiler_build", 90*time.Second, source, compilerBuild, goTool, "build", "-trimpath", "-o", compilerPath, "./cmd/gohp")
	if compilerErr != nil || len(compiler.stderr) != 0 {
		t.Fatalf("compiler build failed: stdout=%s stderr=%s err=%v", compiler.stdout, compiler.stderr, compilerErr)
	}
	generatedPath, reportPath := filepath.Join(candidate, "main.go"), filepath.Join(candidate, "report.json")
	generation, generationErr := preg83ConsumerRun(t, "generation", 30*time.Second, source, compilerBuild, compilerPath,
		"-experimental-pcre2", "-php-version", "8.3", "-in", unitPath, "-out", generatedPath, "-report", reportPath)
	if generationErr != nil || len(generation.stderr) != 0 || len(generation.stdout) != 0 {
		t.Fatalf("generation failed: stdout=%s stderr=%s err=%v", generation.stdout, generation.stderr, generationErr)
	}
	generated, reportBytes := read(t, generatedPath), read(t, reportPath)
	var report struct {
		Schema        string `json:"schema"`
		Status        string `json:"status"`
		Target        string `json:"php_target"`
		RegexProfile  string `json:"regex_profile"`
		SourceHash    string `json:"source_sha256"`
		GeneratedHash string `json:"generated_sha256"`
	}
	if err := json.Unmarshal(reportBytes, &report); err != nil {
		t.Fatal(err)
	}
	if report.Schema != "gohp.build/v0.1" || report.Status != "generated" || report.Target != "8.3" || report.RegexProfile != "pcre2-php83-v1" || report.SourceHash != digest([]byte(unit)) || report.GeneratedHash != digest(generated) {
		t.Fatal("compiler report/source/generated hash mismatch")
	}

	cgoEnv := preg83ConsumerWithEnv(baseEnv, map[string]string{
		"CGO_ENABLED": "1", "CC": resolvedCC, "SDKROOT": sdk,
		"CGO_CFLAGS":  "-I" + filepath.Join(prefix, "include") + " -O2 -Wall -Wextra -Werror",
		"CGO_LDFLAGS": archive,
	})
	binaryPath := filepath.Join(candidate, "program")
	candidateBuild, candidateBuildErr := preg83ConsumerRun(t, "candidate_build", 90*time.Second, source, cgoEnv, goTool,
		"build", "-trimpath", "-tags", "gohp_pcre2", "-o", binaryPath, generatedPath, observerPath)
	if candidateBuildErr != nil {
		t.Fatalf("candidate build failed: stdout=%s stderr=%s err=%v", candidateBuild.stdout, candidateBuild.stderr, candidateBuildErr)
	}
	// A Darwin linker may print a duplicate-static-archive warning. It is
	// retained in evidence, while runtime diagnostics remain an error.
	runEnv := []string{"PATH="}
	candidateRun, candidateRunErr := preg83ConsumerRun(t, "candidate_runtime", 30*time.Second, "", runEnv, binaryPath, casesPath)
	if candidateRunErr != nil || len(candidateRun.stderr) != 0 {
		t.Fatalf("candidate runtime failed: stdout=%s stderr=%s err=%v", candidateRun.stdout, candidateRun.stderr, candidateRunErr)
	}
	actual := preg83ConsumerObserved(t, candidateRun.stdout)
	mismatches := preg83ConsumerMismatchCount(inputs.Expected, actual)
	redact := strings.NewReplacer(workspace, "<WORKSPACE>", repository, "<GOHP_SOURCE>", prefix, "<PCRE2_PREFIX>", sdk, "<SDKROOT>", resolvedCC, "<CC>")
	evidence := map[string]any{
		"schema": "gossip.parser-preg83-generated/v1", "status": "differential_pass", "php_target": "8.3",
		"gohp_commit": preg83ConsumerGoHPCommit, "gohp_archive_sha256": archiveHash,
		"wordpress_sources": inputs.SourceHashes, "manifest_sha256": digest(inputs.Manifest), "reference_sha256": digest(inputs.Reference), "cases_sha256": digest(inputs.CasesBytes),
		"php_image":      inputs.PHPImage,
		"pattern_sha256": digest(inputs.Pattern), "selected_case_count": len(inputs.Selected), "excluded_case": "backtrack_limit_one",
		"excluded_reason": "the generated profile fixes match_limit=100000; the retained limit=1 error case is not rerun under a changed limit",
		"compiler_sha256": digest(read(t, compilerPath)), "compiler_input_sha256": digest([]byte(unit)), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"binary_sha256": digest(read(t, binaryPath)), "driver_template_sha256": digest(inputs.DriverTemplate), "driver_sha256": digest([]byte(unit)), "observer_sha256": digest(observer),
		"harness_sha256": digest(read(t, "parser_preg83_generated_test.go")), "header_sha256": digest(read(t, header)), "static_archive_sha256": digest(read(t, archive)),
		"clang_sha256": digest(read(t, resolvedCC)), "go_version": strings.TrimSpace(string(goVersion.stdout)), "go_runtime_version": runtime.Version(),
		"sdk_version": sdkIdentity.Version, "sdk_settings_sha256": digest(sdkSettings), "go_tool_sha256": digest(read(t, goTool)),
		"candidate_stdout_sha256": digest(candidateRun.stdout), "candidate_exit": 0, "candidate_stderr_bytes": len(candidateRun.stderr),
		"selected_inputs_sha256": digest(append(selectedJSON, '\n')), "reference_reused": true,
		"commands": map[string]string{
			"compiler_build":    "go build -trimpath -o <CANDIDATE>/gohp ./cmd/gohp",
			"generation":        "<CANDIDATE>/gohp -experimental-pcre2 -php-version 8.3 -in <INPUT>/unit.php -out <CANDIDATE>/main.go -report <CANDIDATE>/report.json",
			"candidate_build":   "go build -trimpath -tags gohp_pcre2 -o <CANDIDATE>/program <CANDIDATE>/main.go <CANDIDATE>/observer.go",
			"candidate_runtime": "<CANDIDATE>/program <INPUT>/cases.json",
		},
		"profile":                          map[string]any{"os": runtime.GOOS, "arch": runtime.GOARCH, "cgo": true, "tags": []string{"gohp_pcre2"}, "sdkroot_suffix": "MacOSX.sdk", "cgo_cflags": "-I<PCRE2_PREFIX>/include -O2 -Wall -Wextra -Werror", "cgo_ldflags": "<PCRE2_PREFIX>/lib/libpcre2-8.a"},
		"phases":                           map[string]any{"compiler_build": "pass", "generation": "pass", "candidate_build": "pass", "native_php": "retained_reference_reused", "candidate_runtime": "pass"},
		"candidate_build_stderr_sanitized": redact.Replace(string(candidateBuild.stderr)), "candidate_build_stderr_sha256": digest(candidateBuild.stderr), "candidate_observations": actual, "expected_observations": inputs.Expected,
		"matches": len(inputs.Expected) - mismatches, "mismatches": mismatches,
		"scope": "23 retained PHP 8.3 WordPress block-token preg_match cases at fixed match limit 100000; generated owned probe only; no full next_token/parser conversion, no PHP rerun, no PHP 8.4 claim",
	}
	if mismatches != 0 {
		evidence["status"] = "differential_mismatch"
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("PREG83_GENERATED_EVIDENCE " + string(encoded))
	if mismatches != 0 {
		t.Fatalf("%d of %d generated preg83 observations differ from retained PHP reference", mismatches, len(inputs.Expected))
	}
}

func preg83ConsumerLoadInputs(t *testing.T) preg83ConsumerInputs {
	t.Helper()
	manifest := read(t, "parser-token-reference.json")
	var pin struct {
		Schema          string            `json:"schema"`
		WordPressCommit string            `json:"wordpress_commit"`
		PHPTarget       string            `json:"php_target"`
		PHPImage        string            `json:"php_image"`
		CasesSHA        string            `json:"cases_sha256"`
		PatternSHA      string            `json:"pattern_sha256"`
		Sources         map[string]string `json:"sources"`
	}
	if err := json.Unmarshal(manifest, &pin); err != nil {
		t.Fatal(err)
	}
	wantSources := map[string]string{
		"wp-includes/class-wp-block-parser.php":       "8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b",
		"wp-includes/class-wp-block-parser-block.php": "7e21f483cf0c9deddd2144e026842bf906c72983908f7bd163c0716e8f6f325e",
		"wp-includes/class-wp-block-parser-frame.php": "e5883d82dff099a4b1a23786e9a53b7e29aa73654db83e7ddc444ba17df5e105",
		"license.txt":             "49e97f629ce763abb2f8f1f13da74cb5a8a9a72a0e72a7477d1540c98cef5141",
		"wp-includes/version.php": "6fbc8cf7e4667ba118b92186eea0c9ec35154476790c218d55ead45eb2a1f1d1",
	}
	if pin.Schema != "gossip.parser-token-reference/v1" || pin.WordPressCommit != "54106f6c545fc8b881f6b39b7963002bece6c28c" || pin.PHPTarget != "8.3" || pin.PHPImage != "php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd" || pin.CasesSHA != preg83ConsumerCasesSHA || pin.PatternSHA != preg83ConsumerPatternSHA || !reflect.DeepEqual(pin.Sources, wantSources) {
		t.Fatal("parser token reference pin drift")
	}
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil || baseline.Commit != pin.WordPressCommit {
		t.Fatal("WordPress baseline drift")
	}
	verified := make(map[string][]byte, len(pin.Sources))
	for path, hash := range pin.Sources {
		data := read(t, filepath.Join("../../..", path))
		if digest(data) != hash {
			t.Fatalf("original WordPress source drift: %s", path)
		}
		verified[path] = data
	}
	lines := bytes.Split(verified["wp-includes/class-wp-block-parser.php"], []byte("\n"))
	if len(lines) < 251 {
		t.Fatal("original parser lacks pinned pattern line")
	}
	literalLine := bytes.TrimSpace(lines[250])
	if !bytes.HasPrefix(literalLine, []byte("'")) || !bytes.HasSuffix(literalLine, []byte("',")) {
		t.Fatal("original parser pattern literal boundary drift")
	}
	pattern := append([]byte(nil), literalLine[1:len(literalLine)-2]...)
	if bytes.Contains(pattern, []byte("'")) || bytes.Contains(pattern, []byte(`\\`)) || digest(pattern) != preg83ConsumerPatternSHA {
		t.Fatal("original parser pattern requires reviewed extraction")
	}
	casesBytes := read(t, "parser-token-reference-cases.json")
	if digest(casesBytes) != preg83ConsumerCasesSHA {
		t.Fatal("retained token corpus drift")
	}
	var rawCases []preg83ConsumerCase
	if err := json.Unmarshal(casesBytes, &rawCases); err != nil || len(rawCases) != 24 {
		t.Fatal("invalid retained token corpus")
	}
	selected := make([]preg83ConsumerCase, 0, 23)
	seen := make(map[string]bool, len(rawCases))
	for _, c := range rawCases {
		document, err := base64.StdEncoding.Strict().DecodeString(c.Document)
		if err != nil || base64.StdEncoding.EncodeToString(document) != c.Document || c.Name == "" || seen[c.Name] || len(document) > 4096 || c.Offset < -len(document) || c.Offset > len(document) || (c.Limit != 100000 && c.Limit != 1) {
			t.Fatalf("retained case outside fixed domain: %s", c.Name)
		}
		seen[c.Name] = true
		if c.Limit == 1 {
			if c.Name != "backtrack_limit_one" {
				t.Fatal("unexpected non-default match limit case")
			}
			continue
		}
		selected = append(selected, c)
	}
	if len(selected) != 23 || len(seen) != 24 || !seen["backtrack_limit_one"] {
		t.Fatal("fixed 23-case selection drift")
	}
	reference := read(t, "../../../docs/evidence/parser-token-reference-2026-10-27.json")
	if digest(reference) != preg83ConsumerRefSHA {
		t.Fatal("retained PHP reference drift")
	}
	var retained struct {
		Status     string            `json:"status"`
		Sources    map[string]string `json:"sources"`
		CasesSHA   string            `json:"cases_sha256"`
		PatternSHA string            `json:"pattern_sha256"`
		Native     struct {
			Pattern string `json:"pattern_base64"`
			Cases   []struct {
				Name string         `json:"name"`
				Preg map[string]any `json:"preg_match"`
			} `json:"cases"`
		} `json:"native_observation"`
	}
	if err := json.Unmarshal(reference, &retained); err != nil {
		t.Fatal(err)
	}
	refPattern, err := base64.StdEncoding.Strict().DecodeString(retained.Native.Pattern)
	if err != nil || retained.Status != "native-reference-pass" || !reflect.DeepEqual(retained.Sources, wantSources) || retained.CasesSHA != preg83ConsumerCasesSHA || retained.PatternSHA != preg83ConsumerPatternSHA || !bytes.Equal(refPattern, pattern) || len(retained.Native.Cases) != 24 {
		t.Fatal("retained native reference no longer matches pinned source/corpus/profile")
	}
	expected := make([]map[string]any, 0, 23)
	for i, c := range rawCases {
		ref := retained.Native.Cases[i]
		if ref.Name != c.Name {
			t.Fatal("retained reference case identity drift")
		}
		if c.Limit == 1 {
			continue
		}
		preg := make(map[string]any, len(ref.Preg)-1)
		for key, value := range ref.Preg {
			if key != "error_message" {
				preg[key] = value
			}
		}
		expected = append(expected, map[string]any{"name": c.Name, "preg_match": preg})
	}
	driver := read(t, "parser-preg83-generated-driver.php.txt")
	if strings.Count(string(driver), "@PATTERN_LITERAL@") != 1 {
		t.Fatal("owned preg83 driver must contain exactly one pattern placeholder")
	}
	return preg83ConsumerInputs{Manifest: manifest, Reference: reference, CasesBytes: casesBytes, DriverTemplate: driver, Pattern: pattern, Cases: rawCases, Selected: selected, Expected: expected, SourceHashes: wantSources, PHPImage: pin.PHPImage}
}

// PHP single-quoted strings interpret only escaped quote and escaped backslash.
// Doubling every backslash and escaping every quote therefore restores exactly
// the original runtime pattern bytes; this never edits regex syntax itself.
func preg83ConsumerPHPSingleQuoted(pattern []byte) string {
	escaped := strings.ReplaceAll(string(pattern), `\`, `\\`)
	escaped = strings.ReplaceAll(escaped, "'", `\'`)
	return "'" + escaped + "'"
}

func preg83ConsumerToolchain(t *testing.T) (string, string) {
	t.Helper()
	sdk := os.Getenv("SDKROOT")
	if !filepath.IsAbs(sdk) || !strings.HasSuffix(filepath.Clean(sdk), "MacOSX.sdk") {
		t.Fatal("SDKROOT must be an absolute macOSX.sdk path")
	}
	cc := os.Getenv("CC")
	if cc == "" {
		cc = "clang"
	}
	if strings.ContainsAny(cc, " \t\n") {
		t.Fatal("CC must name one compiler executable, without embedded arguments")
	}
	resolved, err := exec.LookPath(cc)
	if err != nil || !filepath.IsAbs(resolved) {
		t.Fatalf("resolve CC: %v", err)
	}
	if digest(read(t, resolved)) != preg83ConsumerClangSHA {
		t.Fatal("resolved clang differs from the measured profile")
	}
	return resolved, sdk
}

type preg83ConsumerCommandResult struct{ stdout, stderr []byte }

func preg83ConsumerRun(t *testing.T, stage string, timeout time.Duration, dir string, env []string, command string, args ...string) (preg83ConsumerCommandResult, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir, cmd.Env = dir, env
	var result preg83ConsumerCommandResult
	cmd.Stdout, cmd.Stderr = &bytes.Buffer{}, &bytes.Buffer{}
	err := cmd.Run()
	result.stdout = cmd.Stdout.(*bytes.Buffer).Bytes()
	result.stderr = cmd.Stderr.(*bytes.Buffer).Bytes()
	if ctx.Err() != nil {
		t.Fatalf("status=infrastructure_failed: %s timeout", stage)
	}
	return result, err
}

func preg83ConsumerEnv(overrides map[string]string) []string {
	values := make(map[string]string)
	for _, value := range os.Environ() {
		if key, entry, ok := strings.Cut(value, "="); ok {
			values[key] = entry
		}
	}
	for key, value := range overrides {
		values[key] = value
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	// Stable ordering makes accidental environment-dependent evidence easier to
	// review without exposing the inherited values themselves.
	sort.Strings(keys)
	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, key+"="+values[key])
	}
	return result
}

func preg83ConsumerWithEnv(base []string, overrides map[string]string) []string {
	values := make(map[string]string, len(base)+len(overrides))
	for _, value := range base {
		key, entry, _ := strings.Cut(value, "=")
		values[key] = entry
	}
	for key, value := range overrides {
		values[key] = value
	}
	return preg83ConsumerEnv(values)
}

func preg83ConsumerObserved(t *testing.T, raw []byte) []map[string]any {
	t.Helper()
	var observed []map[string]any
	decoder := json.NewDecoder(bytes.NewReader(raw))
	if err := decoder.Decode(&observed); err != nil {
		t.Fatalf("candidate did not emit JSON observations: %v: %s", err, raw)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		t.Fatalf("candidate emitted trailing or invalid JSON: %v", err)
	}
	return observed
}

func preg83ConsumerMismatchCount(expected, actual []map[string]any) int {
	mismatches := 0
	if len(actual) != len(expected) {
		return len(expected)
	}
	for i := range expected {
		if !reflect.DeepEqual(expected[i], actual[i]) {
			mismatches++
		}
	}
	return mismatches
}
