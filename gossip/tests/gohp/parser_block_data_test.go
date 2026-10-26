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

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

// Complete original WordPress data class -> public property snapshot -> real
// element-class decision, without a substitute class or full parser claim.
func TestOriginalWordPressParserBlockData(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE; cached pinned Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run as a non-root user")
	}
	manifest := read(t, "parser-block-data.json")
	var pin struct {
		unitPin
		CasesHash string `json:"cases_sha256"`
		Sources   []struct {
			Path      string     `json:"source_path"`
			Hash      string     `json:"source_sha256"`
			Fragments []fragment `json:"fragments"`
		} `json:"sources"`
	}
	if err := json.Unmarshal(manifest, &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.parser-block-data/v1" || len(pin.Sources) != 3 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid parser-block pin")
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
	type selection struct {
		path        string
		first, last int
	}
	wanted := map[string]selection{
		"_wp_array_get":                     {"wp-includes/functions.php", 5233, 5272},
		"WP_Block_Parser_Block":             {"wp-includes/class-wp-block-parser-block.php", 15, 90},
		"wp_should_add_elements_class_name": {"wp-includes/block-supports/elements.php", 33, 97},
	}
	definitions := []byte("<?php\n")
	seen := map[string]bool{}
	var selections []map[string]any
	for _, source := range pin.Sources {
		if len(source.Fragments) != 1 {
			t.Fatal("expected one declaration per source")
		}
		part := source.Fragments[0]
		if want, ok := wanted[part.Name]; !ok || want != (selection{source.Path, part.First, part.Last}) || seen[part.Name] {
			t.Fatal("unexpected/duplicate original fragment")
		}
		seen[part.Name] = true
		original := read(t, filepath.Join("../../..", source.Path))
		if digest(original) != source.Hash {
			t.Fatal("WordPress source drift")
		}
		lines := bytes.SplitAfter(original, []byte("\n"))
		if part.Last > len(lines) {
			t.Fatal("invalid original range")
		}
		declaration := bytes.Join(lines[part.First-1:part.Last], nil)
		prefix := "function " + part.Name + "("
		if part.Name == "WP_Block_Parser_Block" {
			prefix = "class WP_Block_Parser_Block {"
		}
		if !bytes.HasPrefix(declaration, []byte(prefix)) || !bytes.HasSuffix(declaration, []byte("}\n")) {
			t.Fatal("incomplete original declaration")
		}
		definitions = append(definitions, declaration...)
		selections = append(selections, map[string]any{"name": part.Name, "source_path": source.Path, "source_sha256": source.Hash,
			"first_line": part.First, "last_line": part.Last, "sha256": digest(declaration)})
	}
	if len(seen) != len(wanted) {
		t.Fatal("missing declaration")
	}
	drivers := read(t, "parser-block-drivers.php.txt")
	definitions = append(definitions, drivers...)
	casesBytes := read(t, "parser-block-cases.json")
	if digest(casesBytes) != pin.CasesHash {
		t.Fatal("case corpus drift")
	}
	var cases []struct {
		Name         string `json:"name"`
		BlockName    any    `json:"block_name"`
		Attrs        any    `json:"attrs"`
		InnerBlocks  any    `json:"inner_blocks"`
		InnerHTML    any    `json:"inner_html"`
		InnerContent any    `json:"inner_content"`
		Options      any    `json:"options"`
		Expected     *bool  `json:"expected"`
	}
	numericBundleDecode(t, casesBytes, &cases)
	if len(cases) != 16 {
		t.Fatal("expected16 parser-block cases")
	}
	var expected []map[string]any
	seenCases := map[string]bool{}
	for _, c := range cases {
		if c.Name == "" || seenCases[c.Name] || c.Expected == nil {
			t.Fatal("invalid parser-block expectation")
		}
		seenCases[c.Name] = true
		properties := map[string]any{"array": []any{
			[]any{"blockName", c.BlockName}, []any{"attrs", c.Attrs}, []any{"innerBlocks", c.InnerBlocks},
			[]any{"innerHTML", c.InnerHTML}, []any{"innerContent", c.InnerContent},
		}}
		expected = append(expected, map[string]any{"name": c.Name, "type": "object", "class": "WP_Block_Parser_Block",
			"properties": properties, "snapshot": properties, "decision": *c.Expected,
			"arguments_after": map[string]any{"block_name": c.BlockName, "attrs": c.Attrs, "inner_blocks": c.InnerBlocks,
				"inner_html": c.InnerHTML, "inner_content": c.InnerContent, "options": c.Options}})
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
	// Reuse the bounded node transport, not the old test entrypoint. Each cut
	// delimiter and import replacement must occur exactly once; record both hashes.
	goTransportOriginal := read(t, "array-get-observer.go.txt")
	phpTransportOriginal := read(t, "array-get-observer.php")
	goCut, phpStart, phpEnd := "// init is compiled beside GoHP's generated main.go.", "const ARRAY_GET_MAX_DEPTH", "$cases = json_decode("
	if bytes.Count(goTransportOriginal, []byte(goCut)) != 1 || bytes.Count(phpTransportOriginal, []byte(phpStart)) != 1 ||
		bytes.Count(phpTransportOriginal, []byte(phpEnd)) != 1 {
		t.Fatal("transport adapter drift")
	}
	goTransport := append([]byte(nil), goTransportOriginal[:bytes.Index(goTransportOriginal, []byte(goCut))]...)
	if bytes.Count(goTransport, []byte("\n\t\"reflect\"")) != 1 {
		t.Fatal("Go transport import drift")
	}
	goTransport = bytes.Replace(goTransport, []byte("\n\t\"reflect\""), nil, 1)
	phpTransport := append([]byte("<?php\n"), phpTransportOriginal[bytes.Index(phpTransportOriginal, []byte(phpStart)):bytes.Index(phpTransportOriginal, []byte(phpEnd))]...)
	phpObserver, goObserver := read(t, "parser-block-observer.php"), read(t, "parser-block-observer.go.txt")
	for name, data := range map[string][]byte{"helpers.php": definitions, "parser-block-cases.json": casesBytes, "observer.php": phpObserver, "transport.php": phpTransport, "wordpress-license.txt": read(t, "../../../license.txt")} {
		write(t, filepath.Join(input, name), data)
	}
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	write(t, filepath.Join(candidate, "transport.go"), goTransport)
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456", "--env=GOCACHE=/candidate/cache",
		"--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local", "--env=GOPROXY=off", "--env=GOSUMDB=off",
		"--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "parser_block_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/helpers.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
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
	t.Log("status=generated; complete original class and two functions unchanged")
	_, compileErr := container(t, "parser_block_build", append(buildArgs, "--entrypoint=go", pin.GoImage, "build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/transport.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("build stderr: %s", compileErr)
	}
	t.Log("status=compiled; no compatibility claim before observations")
	actual, actualErr := container(t, "parser_block_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "parser_block_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "manifest_sha256": digest(manifest), "wordpress_commit": pin.WordPressCommit,
		"fragments": selections, "original_function_count": 2, "original_class_count": 1, "original_constructor_count": 1, "owned_driver_count": 1,
		"compiled_bundle_sha256": digest(definitions), "native_bundle_sha256": digest(definitions), "drivers_sha256": digest(drivers),
		"cases_sha256": digest(casesBytes), "case_count": len(cases), "public_property_count": 5, "constructor_argument_count": 5,
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "php_image": pin.PHPImage, "php_version": phpVersion, "php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver),
		"go_transport":  map[string]string{"original_sha256": digest(goTransportOriginal), "adapted_sha256": digest(goTransport)},
		"php_transport": map[string]string{"original_sha256": digest(phpTransportOriginal), "adapted_sha256": digest(phpTransport)},
		"binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "one complete original WP_Block_Parser_Block class and constructor, two unchanged original helpers and one owned source glue function;16 typed class/snapshot/decision/argument-state comparisons; data-only inner_blocks, no full parser/registry/render/CSS/HTML/bootstrap claim; historical configurations unchanged/not rerun",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
