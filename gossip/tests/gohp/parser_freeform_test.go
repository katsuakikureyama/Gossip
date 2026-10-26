package gohp_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

// The Go candidate selects one unchanged parser method, while the PHP oracle
// loads the complete original parser and its two original require_once files.
func TestOriginalWordPressParserFreeform(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE; cached pinned Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run as a non-root user")
	}
	manifest := read(t, "parser-freeform.json")
	type sourcePin struct {
		Path      string     `json:"source_path"`
		Hash      string     `json:"source_sha256"`
		Role      string     `json:"role"`
		Fragments []fragment `json:"fragments"`
	}
	var pin struct {
		unitPin
		CasesHash string      `json:"cases_sha256"`
		Sources   []sourcePin `json:"sources"`
	}
	if err := json.Unmarshal(manifest, &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.parser-freeform/v1" || len(pin.Sources) != 3 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid parser-freeform pin")
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
	wanted := []sourcePin{
		{Path: "wp-includes/class-wp-block-parser-block.php", Role: "complete_class_and_native_file", Fragments: []fragment{{Name: "WP_Block_Parser_Block", First: 15, Last: 90}}},
		{Path: "wp-includes/class-wp-block-parser.php", Role: "selected_class_method_and_complete_native_file", Fragments: []fragment{
			{Name: "WP_Block_Parser.prefix", First: 16, Last: 50}, {Name: "WP_Block_Parser.freeform", First: 314, Last: 316}, {Name: "WP_Block_Parser.close", First: 396, Last: 396},
		}},
		{Path: "wp-includes/class-wp-block-parser-frame.php", Role: "native_require_once_dependency", Fragments: []fragment{}},
	}
	definitions := []byte("<?php\n")
	originals := map[string][]byte{}
	selections := []map[string]any{}
	sourceRecords := []map[string]any{}
	for i, source := range pin.Sources {
		want := wanted[i]
		if source.Path != want.Path || source.Role != want.Role || !reflect.DeepEqual(source.Fragments, want.Fragments) {
			t.Fatal("unexpected source selection/order")
		}
		original := read(t, filepath.Join("../../..", source.Path))
		if digest(original) != source.Hash {
			t.Fatal("WordPress source drift")
		}
		originals[filepath.Base(source.Path)] = original
		sourceRecords = append(sourceRecords, map[string]any{"source_path": source.Path, "source_sha256": source.Hash, "role": source.Role})
		lines := bytes.SplitAfter(original, []byte("\n"))
		for _, part := range source.Fragments {
			if part.First < 1 || part.Last < part.First || part.Last > len(lines) {
				t.Fatal("invalid original range")
			}
			declaration := bytes.Join(lines[part.First-1:part.Last], nil)
			valid := false
			switch part.Name {
			case "WP_Block_Parser_Block":
				valid = bytes.HasPrefix(declaration, []byte("class WP_Block_Parser_Block {\n")) && bytes.HasSuffix(declaration, []byte("}\n"))
			case "WP_Block_Parser.prefix":
				valid = bytes.HasPrefix(declaration, []byte("class WP_Block_Parser {\n")) && bytes.HasSuffix(declaration, []byte("\tpublic $stack;\n"))
			case "WP_Block_Parser.freeform":
				valid = bytes.HasPrefix(declaration, []byte("\tpublic function freeform( $inner_html ) {\n")) && bytes.HasSuffix(declaration, []byte("\t}\n"))
			case "WP_Block_Parser.close":
				valid = bytes.Equal(declaration, []byte("}\n"))
			}
			if !valid {
				t.Fatal("original declaration boundary drift")
			}
			definitions = append(definitions, declaration...)
			selections = append(selections, map[string]any{"name": part.Name, "source_path": source.Path, "first_line": part.First, "last_line": part.Last, "sha256": digest(declaration)})
		}
	}
	drivers := read(t, "parser-freeform-drivers.php.txt")
	definitions = append(definitions, drivers...)
	nativeDrivers := append([]byte("<?php\n"), drivers...)
	casesBytes := read(t, "parser-freeform-cases.json")
	if digest(casesBytes) != pin.CasesHash {
		t.Fatal("case corpus drift")
	}
	var cases []struct {
		Name   string `json:"name"`
		First  string `json:"first"`
		Second string `json:"second"`
	}
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 12 {
		t.Fatal("expected12 freeform cases")
	}
	block := func(html string) map[string]any {
		return map[string]any{"type": "object", "class": "WP_Block_Parser_Block", "properties": [][]any{
			{"blockName", nil}, {"attrs", []any{}}, {"innerBlocks", []any{}}, {"innerHTML", html}, {"innerContent", []any{html}},
		}}
	}
	expected := []map[string]any{}
	seen := map[string]bool{}
	for _, c := range cases {
		if c.Name == "" || seen[c.Name] || len(c.First) > 4096 || len(c.Second) > 4096 || !utf8.ValidString(c.First) || !utf8.ValidString(c.Second) {
			t.Fatal("invalid freeform case")
		}
		seen[c.Name] = true
		expected = append(expected, map[string]any{"name": c.Name, "parser": map[string]any{
			"type": "object", "class": "WP_Block_Parser", "properties": [][]any{{"document", nil}, {"offset", nil}, {"output", nil}, {"stack", nil}},
		}, "first": block(c.First), "second": block(c.Second), "distinct_blocks": true, "arguments_after": map[string]string{"first": c.First, "second": c.Second}})
	}
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
	phpObserver, goObserver := read(t, "parser-freeform-observer.php"), read(t, "parser-freeform-observer.go.txt")
	for name, data := range originals {
		write(t, filepath.Join(input, name), data)
	}
	for name, data := range map[string][]byte{"helpers.php": definitions, "drivers.php": nativeDrivers, "parser-freeform-cases.json": casesBytes, "observer.php": phpObserver, "wordpress-license.txt": read(t, "../../../license.txt")} {
		write(t, filepath.Join(input, name), data)
	}
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly", "--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate", "--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456", "--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local", "--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "parser_freeform_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c", "go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/helpers.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
	if len(buildErr) != 0 {
		t.Fatalf("compiler stderr: %s", buildErr)
	}
	generated, reportBytes := read(t, filepath.Join(candidate, "main.go")), read(t, filepath.Join(candidate, "report.json"))
	var report struct {
		Schema          string `json:"schema"`
		Status          string `json:"status"`
		Target          string `json:"php_target"`
		SourceHash      string `json:"source_sha256"`
		GeneratedHash   string `json:"generated_sha256"`
		CompilerVersion string `json:"compiler_version"`
	}
	if err := json.Unmarshal(reportBytes, &report); err != nil {
		t.Fatal(err)
	}
	if report.Schema != "gohp.build/v0.1" || report.Status != "generated" || report.Target != pin.PHPTarget || report.SourceHash != digest(definitions) || report.GeneratedHash != digest(generated) {
		t.Fatal("report/source/output mismatch")
	}
	_, compileErr := container(t, "parser_freeform_build", append(buildArgs, "--entrypoint=go", pin.GoImage, "build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("build stderr: %s", compileErr)
	}
	actual, actualErr := container(t, "parser_freeform_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly", "--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "parser_freeform_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	nativeHashes := map[string]string{}
	for name, data := range originals {
		nativeHashes[name] = digest(data)
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "manifest_sha256": digest(manifest), "wordpress_commit": pin.WordPressCommit, "sources": sourceRecords, "fragments": selections,
		"complete_original_class_count": 1, "selected_original_class_count": 1, "original_method_count": 1, "original_constructor_count": 1, "original_function_count": 0, "owned_driver_count": 1,
		"method_calls": 2 * len(cases), "block_constructions": 2 * len(cases), "native_source_files_sha256": nativeHashes, "native_drivers_sha256": digest(nativeDrivers), "compiled_bundle_sha256": digest(definitions), "drivers_sha256": digest(drivers),
		"cases_sha256": digest(casesBytes), "case_count": len(cases), "gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes), "go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)),
		"php_image": pin.PHPImage, "php_version": phpVersion, "php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "Go: unchanged original Block class plus selected original parser class prefix/freeform/closing brace; PHP: complete original parser/Block/Frame files; one owned driver;12 cases,24 real freeform calls/new Blocks; no parse/proceed/tokenizer/stack/render/bootstrap claim; historical configurations unchanged/not rerun",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
