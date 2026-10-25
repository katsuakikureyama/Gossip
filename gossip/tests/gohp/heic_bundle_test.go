package gohp_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

// This is a new, independent HEIC helper bundle. It preserves the earlier
// eighty-nine observations while checking the expanded configuration.
func TestOriginalWordPressHEICHelperBundle(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	manifest := read(t, "heic-helper-bundle.json")
	var pin struct {
		unitPin
		Sources []struct {
			Path      string     `json:"source_path"`
			Hash      string     `json:"source_sha256"`
			Fragments []fragment `json:"fragments"`
		} `json:"sources"`
	}
	if err := json.Unmarshal(manifest, &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.heic-helper-bundle/v1" || len(pin.Sources) != 3 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid HEIC helper bundle pin")
	}
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil {
		t.Fatal(err)
	}
	if baseline.Commit != pin.WordPressCommit {
		t.Fatal("HEIC bundle and fork baseline disagree")
	}

	wanted := map[string]string{
		"__return_true": "wp-includes/functions.php", "__return_false": "wp-includes/functions.php",
		"__return_zero": "wp-includes/functions.php", "__return_empty_array": "wp-includes/functions.php",
		"__return_null": "wp-includes/functions.php", "__return_empty_string": "wp-includes/functions.php",
		"wp_validate_boolean": "wp-includes/functions.php", "_autop_newline_preservation_helper": "wp-includes/formatting.php",
		"zeroise": "wp-includes/formatting.php", "trailingslashit": "wp-includes/formatting.php",
		"untrailingslashit": "wp-includes/formatting.php", "stripslashes_from_strings_only": "wp-includes/formatting.php",
		"wp_slash":                   "wp-includes/formatting.php",
		"wp_array_slice_assoc":       "wp-includes/functions.php",
		"wp_is_heic_image_mime_type": "wp-includes/functions.php",
		"_is_utf8_charset":           "wp-includes/compat.php",
	}
	var bundle bytes.Buffer
	bundle.WriteString("<?php\n")
	var selections []map[string]any
	seenSources, seenFunctions := map[string]bool{}, map[string]bool{}
	for _, sourcePin := range pin.Sources {
		if (sourcePin.Path != "wp-includes/functions.php" && sourcePin.Path != "wp-includes/formatting.php" && sourcePin.Path != "wp-includes/compat.php") || seenSources[sourcePin.Path] {
			t.Fatal("unexpected or duplicate original source")
		}
		seenSources[sourcePin.Path] = true
		original := read(t, filepath.Join("../../..", sourcePin.Path))
		if digest(original) != sourcePin.Hash {
			t.Fatal("original source drift")
		}
		lines := bytes.SplitAfter(original, []byte("\n"))
		for _, part := range sourcePin.Fragments {
			if wanted[part.Name] != sourcePin.Path || seenFunctions[part.Name] || part.First < 1 || part.Last < part.First || part.Last > len(lines) {
				t.Fatal("unexpected, duplicate or invalid fragment")
			}
			seenFunctions[part.Name] = true
			excerpt := bytes.Join(lines[part.First-1:part.Last], nil)
			if part.Name == "wp_array_slice_assoc" && (part.First != 5173 || part.Last != 5183 || digest(excerpt) != "887f2bc5b08b05075261d1b5ff38dfbe90c9de9a08efd86f3a03f5fd6367a213") {
				t.Fatal("slice fragment must retain the complete original including : array")
			}
			if part.Name == "_is_utf8_charset" && (part.First != 102 || part.Last != 111 || digest(excerpt) != "241a7691e4461f13f8656fb12b8275fb4eaa71e34b983ef63301965286df5739") {
				t.Fatal("UTF-8 charset fragment must retain the complete unchanged original")
			}
			if part.Name == "wp_is_heic_image_mime_type" && (part.First != 9465 || part.Last != 9474 || digest(excerpt) != "de4452e658e48a6cac0d02b501b425bb8ac60872fdbbe9ab200a107f98223585") {
				t.Fatal("HEIC MIME-label fragment must retain the complete unchanged original")
			}
			if !bytes.HasPrefix(excerpt, []byte("function "+part.Name+"(")) || !bytes.HasSuffix(excerpt, []byte("}\n")) {
				t.Fatal("incomplete original function")
			}
			selections = append(selections, map[string]any{
				"name": part.Name, "source_path": sourcePin.Path, "source_sha256": sourcePin.Hash,
				"original_first_line": part.First, "original_last_line": part.Last,
				"bundle_first_line": bytes.Count(bundle.Bytes(), []byte("\n")) + 1, "sha256": digest(excerpt),
			})
			bundle.Write(excerpt)
			bundle.WriteByte('\n')
		}
	}
	if len(seenSources) != 3 || len(seenFunctions) != 16 {
		t.Fatal("bundle must contain exactly sixteen unchanged original functions")
	}
	drivers := read(t, "heic-bundle-drivers.php.txt")
	driverFirstLine := bytes.Count(bundle.Bytes(), []byte("\n")) + 1
	bundle.Write(drivers) // Owned composition calls; never substitutes original bodies.
	definitions := bundle.Bytes()

	casesBytes := read(t, "heic-bundle-cases.json")
	decoder := json.NewDecoder(bytes.NewReader(casesBytes))
	decoder.UseNumber()
	var cases []struct {
		Name      string `json:"name"`
		Operation string `json:"operation"`
		Args      []any  `json:"args"`
		ArgsAfter []any  `json:"args_after"`
		Expected  struct {
			Type  string `json:"type"`
			Value any    `json:"value"`
		} `json:"expected"`
	}
	if err := decoder.Decode(&cases); err != nil {
		t.Fatal(err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		t.Fatal("HEIC helper bundle fixture has trailing JSON values")
	}
	if len(cases) != 119 {
		t.Fatal("expected one hundred nineteen owned HEIC helper bundle cases")
	}
	expected := make([]map[string]any, 0, len(cases))
	seenCases, operations := map[string]bool{}, map[string]int{}
	for _, c := range cases {
		if c.Name == "" || seenCases[c.Name] {
			t.Fatal("missing or duplicate case name")
		}
		seenCases[c.Name] = true
		if c.Args == nil || c.ArgsAfter == nil || len(c.Args) != len(c.ArgsAfter) {
			t.Fatalf("case %s requires explicit argument and post-call lists of matching length", c.Name)
		}
		// All transports are bounded before either isolated observer runs.
		normalizeArraySliceAssocValue(t, c.Args, 0, c.Name+" args")
		normalizeArraySliceAssocValue(t, c.ArgsAfter, 0, c.Name+" args_after")
		normalizeArraySliceAssocValue(t, c.Expected.Value, 0, c.Name+" expected")
		heicPairOperation := c.Operation == "gossip_slice_heic" || c.Operation == "gossip_slice_unescape_heic"
		if c.Operation == "wp_is_heic_image_mime_type" || c.Operation == "gossip_heic_boolean" {
			if len(c.Args) != 1 {
				t.Fatalf("%s requires exactly one scalar MIME-label argument", c.Operation)
			}
			normalizeHEICBundleScalar(t, c.Args[0], c.Name+" HEIC args")
			normalizeHEICBundleScalar(t, c.ArgsAfter[0], c.Name+" HEIC args_after")
		}
		pairResult := false
		switch c.Operation {
		case "wp_array_slice_assoc", "gossip_slice_slash", "gossip_slash_slice", "gossip_slice_shallow_roundtrip", "gossip_slice_charset", "gossip_slice_unescape_charset", "gossip_slice_heic", "gossip_slice_unescape_heic":
			pairResult = true
		}
		switch c.Operation {
		case "wp_array_slice_assoc", "gossip_slice_slash", "gossip_slash_slice", "gossip_slice_shallow_roundtrip", "gossip_slice_shape", "gossip_slice_boolean", "gossip_slice_charset", "gossip_slice_unescape_charset", "gossip_slice_heic", "gossip_slice_unescape_heic":
			if len(c.Args) != 2 {
				t.Fatalf("%s requires pair input and selector list", c.Operation)
			}
			for _, args := range [][]any{c.Args, c.ArgsAfter} {
				pairs, pairOK := args[0].([]any)
				keys, keysOK := args[1].([]any)
				if !pairOK || !keysOK {
					t.Fatalf("%s arguments must be pair and selector lists", c.Name)
				}
				normalizeArraySliceAssocPairs(t, pairs, c.Name+" pairs")
				normalizeArraySliceAssocSelectors(t, keys, c.Name+" keys")
				if heicPairOperation {
					for pairIndex, rawPair := range pairs {
						pair := rawPair.([]any)
						normalizeHEICBundleScalar(t, pair[1], c.Name+" root pair value "+strconv.Itoa(pairIndex))
					}
				}
			}
			if !reflect.DeepEqual(c.Args[1], c.ArgsAfter[1]) {
				t.Fatalf("%s must expect unchanged selector values", c.Name)
			}
			if heicPairOperation && !reflect.DeepEqual(
				normalizeArraySliceAssocPairs(t, c.Args[0].([]any), c.Name+" canonical input pairs"),
				normalizeArraySliceAssocPairs(t, c.ArgsAfter[0].([]any), c.Name+" canonical input-after pairs"),
			) {
				t.Fatalf("%s must retain canonical root pairs after the call", c.Name)
			}
		default:
			if !reflect.DeepEqual(c.Args, c.ArgsAfter) {
				t.Fatalf("%s must expect unchanged ordinary arguments", c.Name)
			}
		}
		if pairResult {
			pairs, ok := c.Expected.Value.([]any)
			if !ok || c.Expected.Type != "array" {
				t.Fatalf("%s requires an array result represented as ordered pairs", c.Name)
			}
			normalizeArraySliceAssocPairs(t, pairs, c.Name+" expected pairs")
			if heicPairOperation {
				for pairIndex, rawPair := range pairs {
					pair := rawPair.([]any)
					normalizeHEICBundleScalar(t, pair[1], c.Name+" expected root pair value "+strconv.Itoa(pairIndex))
				}
			}
		}
		if err := validateIntegratedBundleJSON(c.Args); err != nil {
			t.Fatalf("invalid case arguments %s: %v", c.Name, err)
		}
		if err := validateIntegratedBundleJSON(c.Expected.Value); err != nil {
			t.Fatalf("invalid expected value %s: %v", c.Name, err)
		}
		if c.Expected.Type != integratedBundleJSONType(c.Expected.Value) {
			t.Fatalf("expected type mismatch for %s", c.Name)
		}
		operations[c.Operation]++
		expected = append(expected, map[string]any{
			"name": c.Name, "type": c.Expected.Type, "value": c.Expected.Value, "args_after": c.ArgsAfter,
		})
	}
	for name := range wanted {
		want := 1
		if name == "wp_array_slice_assoc" {
			want = 14
		}
		if name == "_is_utf8_charset" {
			want = 28
		}
		if name == "wp_is_heic_image_mime_type" {
			want = 24
		}
		if operations[name] != want {
			t.Fatalf("missing direct original probe: %s", name)
		}
	}
	for _, name := range []string{"gossip_zero_pad", "gossip_false_boolean", "gossip_empty_string_path", "gossip_path_pad"} {
		if operations[name] != 1 {
			t.Fatalf("missing scalar composition probe: %s", name)
		}
	}
	if operations["gossip_capture_newline_unslash"] != 4 || operations["gossip_unslash_array_keys"] != 2 || operations["gossip_unslash_array_shape"] != 2 ||
		operations["gossip_slash_scalar_roundtrip"] != 3 || operations["gossip_slash_list_shallow_roundtrip"] != 3 ||
		operations["gossip_slash_capture_newline_unslash"] != 1 || operations["gossip_slash_array_shape"] != 1 || operations["wp_array_slice_assoc"] != 14 || operations["gossip_slice_slash"] != 2 || operations["gossip_slash_slice"] != 1 ||
		operations["gossip_slice_shallow_roundtrip"] != 2 || operations["gossip_slice_shape"] != 1 || operations["gossip_slice_boolean"] != 2 || operations["gossip_slice_charset"] != 2 || operations["gossip_slice_unescape_charset"] != 2 || operations["gossip_slice_heic"] != 2 || operations["gossip_slice_unescape_heic"] != 2 || operations["gossip_heic_boolean"] != 2 || operations["gossip_charset_boolean"] != 2 || len(operations) != 38 {
		t.Fatal("unexpected HEIC composition coverage")
	}
	expectedBytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	phpVersion := requirePHPOracle(t, pin.PHPImage)
	workspace := t.TempDir()
	source, input, candidate := filepath.Join(workspace, "gohp"), filepath.Join(workspace, "input"), filepath.Join(workspace, "candidate")
	for _, dir := range []string{source, input, candidate} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}
	archiveHash := exportGoHP(t, repository, pin.GoHPCommit, source)
	phpObserver, goObserver := read(t, "heic-bundle-observer.php"), read(t, "heic-bundle-observer.go.txt")
	write(t, filepath.Join(input, "helpers.php"), definitions)
	write(t, filepath.Join(input, "cases.json"), casesBytes)
	write(t, filepath.Join(input, "observer.php"), phpObserver)
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456",
		"--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local",
		"--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "heic_bundle_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/helpers.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
	if len(buildErr) != 0 {
		t.Fatalf("unexpected compiler stderr: %s", buildErr)
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
		t.Fatal("generation report does not match bundle bytes")
	}
	t.Log("status=generated; HEIC bundle source/report hashes verified")
	_, compileErr := container(t, "heic_bundle_build", append(buildArgs, "--entrypoint=go", pin.GoImage, "build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("unexpected build stderr: %s", compileErr)
	}
	actual, actualErr := container(t, "heic_bundle_candidate", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly", "--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "heic_bundle_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("status=differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "wordpress_commit": pin.WordPressCommit,
		"manifest_sha256": digest(manifest), "selections": selections, "original_function_count": 16, "owned_driver_count": 22,
		"driver_first_line": driverFirstLine, "driver_sha256": digest(drivers), "bundle_sha256": digest(definitions),
		"cases_sha256": digest(casesBytes), "case_count": len(cases), "direct_call_count": 79, "composed_call_count": 40,
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "php_image": pin.PHPImage, "php_target": pin.PHPTarget, "go_version": strings.TrimSpace(string(buildOut)),
		"php_version": phpVersion, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "binary_sha256": digest(read(t, filepath.Join(candidate, "program"))),
		"typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "sixteen unchanged WordPress helpers and twenty-two owned composition drivers translated together in one binary; 119 bounded observations preserve the earlier root associative key/value, UTF-8 charset and nested-list coverage while adding twenty-four scalar MIME-label calls plus scalar-root HEIC slice compositions; observes actual returns and caller arguments after each call; no image-byte inspection, decoding, upload validation, nested associative values, recursive unescaping, full-file/core/startup, HTTP, DB, theme or plugin compatibility",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}

func normalizeHEICBundleScalar(t *testing.T, value any, field string) any {
	t.Helper()
	switch typed := value.(type) {
	case nil, bool, string:
		return typed
	case json.Number:
		integer, err := typed.Int64()
		if err != nil || integer < -1_000_000 || integer > 1_000_000 {
			t.Fatalf("%s must be a modest signed integer: %v", field, value)
		}
		return integer
	case []any, map[string]any:
		t.Fatalf("%s must be a scalar; HEIC operations do not transport nested array/object values", field)
		return nil
	default:
		t.Fatalf("%s has unsupported JSON value type %T", field, value)
		return nil
	}
}
