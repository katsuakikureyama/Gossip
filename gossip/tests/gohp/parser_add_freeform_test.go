package gohp_test

import (
	"bytes"
	"encoding/base64"
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

// The Go candidate selects two unchanged parser methods, while the PHP oracle
// loads the complete original parser and its two original require_once files.
func TestOriginalWordPressParserAddFreeform(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE; cached pinned Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run as a non-root user")
	}
	manifest := read(t, "parser-add-freeform.json")
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
	if pin.Schema != "gossip.parser-add-freeform/v1" || len(pin.Sources) != 3 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid parser-add-freeform pin")
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
		{Path: "wp-includes/class-wp-block-parser.php", Role: "selected_class_methods_and_complete_native_file", Fragments: []fragment{
			{Name: "WP_Block_Parser.prefix", First: 16, Last: 50}, {Name: "WP_Block_Parser.freeform", First: 314, Last: 316}, {Name: "WP_Block_Parser.add_freeform", First: 327, Last: 335}, {Name: "WP_Block_Parser.close", First: 396, Last: 396},
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
			case "WP_Block_Parser.add_freeform":
				valid = bytes.HasPrefix(declaration, []byte("\tpublic function add_freeform( $length = null ) {\n")) && bytes.HasSuffix(declaration, []byte("\t}\n"))
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
	drivers := read(t, "parser-add-freeform-drivers.php.txt")
	definitions = append(definitions, drivers...)
	nativeDrivers := append([]byte("<?php\n"), drivers...)
	casesBytes := read(t, "parser-add-freeform-cases.json")
	if digest(casesBytes) != pin.CasesHash {
		t.Fatal("case corpus drift")
	}
	var cases []struct {
		Name         string   `json:"name"`
		Document     string   `json:"document"`
		Offset       int64    `json:"offset"`
		Length       any      `json:"length"`
		Mode         string   `json:"mode"`
		Seed         []string `json:"seed"`
		ExpectedHTML *string  `json:"expected_html"`
	}
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 16 {
		t.Fatal("expected16 add-freeform cases")
	}
	b64 := func(s string) string { return base64.StdEncoding.EncodeToString([]byte(s)) }
	stringNode := func(encoded string) map[string]any { return map[string]any{"type": "string", "base64": encoded} }
	expected := []map[string]any{}
	seen := map[string]bool{}
	appended := 0
	for _, c := range cases {
		if c.Name == "" || seen[c.Name] || len(c.Document) > 4096 || !utf8.ValidString(c.Document) || c.Offset < -1000000 || c.Offset > 1000000 || (c.Mode != "omitted" && c.Mode != "explicit") || len(c.Seed) > 4 {
			t.Fatal("invalid add-freeform case")
		}
		seen[c.Name] = true
		if c.Length != nil {
			number, ok := c.Length.(json.Number)
			if !ok {
				t.Fatal("length requires int/null")
			}
			n, err := number.Int64()
			if err != nil || n < -1000000 || n > 1000000 {
				t.Fatal("length out of bounds")
			}
		}
		output := []any{}
		seedBase64 := []string{}
		for _, seed := range c.Seed {
			if !utf8.ValidString(seed) || len(seed) > 4096 {
				t.Fatal("invalid seed")
			}
			encoded := b64(seed)
			seedBase64 = append(seedBase64, encoded)
			output = append(output, stringNode(encoded))
		}
		if c.ExpectedHTML != nil {
			raw, err := base64.StdEncoding.DecodeString(*c.ExpectedHTML)
			if err != nil || len(raw) > 4096 || base64.StdEncoding.EncodeToString(raw) != *c.ExpectedHTML {
				t.Fatal("expected HTML must be bounded canonical base64")
			}
			html := stringNode(*c.ExpectedHTML)
			output = append(output, map[string]any{"type": "array", "properties": [][]any{{"blockName", nil}, {"attrs", []any{}}, {"innerBlocks", []any{}}, {"innerHTML", html}, {"innerContent", []any{html}}}})
			appended++
		}
		expected = append(expected, map[string]any{"name": c.Name, "return_type": "null", "return_value": nil,
			"parser":          map[string]any{"type": "object", "class": "WP_Block_Parser", "document_base64": b64(c.Document), "offset": c.Offset, "stack": []any{}, "output": output},
			"arguments_after": map[string]any{"document_base64": b64(c.Document), "offset": c.Offset, "length": c.Length, "mode": c.Mode, "seed_base64": seedBase64}})
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
	phpObserver, goObserver := read(t, "parser-add-freeform-observer.php"), read(t, "parser-add-freeform-observer.go.txt")
	for name, data := range originals {
		write(t, filepath.Join(input, name), data)
	}
	for name, data := range map[string][]byte{"helpers.php": definitions, "drivers.php": nativeDrivers, "parser-add-freeform-cases.json": casesBytes, "observer.php": phpObserver, "wordpress-license.txt": read(t, "../../../license.txt")} {
		write(t, filepath.Join(input, name), data)
	}
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly", "--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate", "--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456", "--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local", "--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "parser_add_freeform_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c", "go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/helpers.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
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
	_, compileErr := container(t, "parser_add_freeform_build", append(buildArgs, "--entrypoint=go", pin.GoImage, "build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("build stderr: %s", compileErr)
	}
	actual, actualErr := container(t, "parser_add_freeform_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly", "--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "parser_add_freeform_php", pin.PHPImage, input, "/input/observer.php")
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
		"complete_original_class_count": 1, "selected_original_class_count": 1, "original_method_count": 2, "original_constructor_count": 1, "original_function_count": 0, "owned_driver_count": 1,
		"method_calls": len(cases) + appended, "add_freeform_calls": len(cases), "freeform_calls": appended, "block_constructions": appended, "skip_cases": len(cases) - appended, "native_source_files_sha256": nativeHashes, "native_drivers_sha256": digest(nativeDrivers), "compiled_bundle_sha256": digest(definitions), "drivers_sha256": digest(drivers),
		"cases_sha256": digest(casesBytes), "case_count": len(cases), "gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes), "go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)),
		"php_image": pin.PHPImage, "php_version": phpVersion, "php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "Go: complete original Block plus selected original Parser prefix/freeform/add_freeform/closing; PHP: complete original parser/Block/Frame files; one owned state-initializing source driver;16 cases, actual ordered array output/return/state/input-after and raw byte slices; no parse/proceed/tokenizer/stack/render/bootstrap claim; historical configurations unchanged/not rerun",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
