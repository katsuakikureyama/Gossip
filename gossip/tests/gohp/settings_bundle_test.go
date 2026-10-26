package gohp_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

// Separate integrated configuration; historical results and pins are unchanged.
func TestOriginalWordPressSettingsHelperBundle(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a clone containing the pinned commit; cached Docker images required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	manifest := read(t, "settings-helper-bundle.json")
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
	if pin.Schema != "gossip.settings-helper-bundle/v1" || len(pin.Sources) != 3 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) || phporacle.ValidateImage(pin.PHPImage) != nil {
		t.Fatal("invalid edge-bundle pin")
	}
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil {
		t.Fatal(err)
	}
	if baseline.Commit != pin.WordPressCommit {
		t.Fatal("boundary and WordPress baseline disagree")
	}
	type selection struct {
		path        string
		first, last int
	}
	wanted := map[string]selection{
		"__return_true":                      {"wp-includes/functions.php", 7181, 7183},
		"__return_false":                     {"wp-includes/functions.php", 7196, 7198},
		"__return_zero":                      {"wp-includes/functions.php", 7209, 7211},
		"__return_empty_array":               {"wp-includes/functions.php", 7222, 7224},
		"__return_null":                      {"wp-includes/functions.php", 7235, 7237},
		"__return_empty_string":              {"wp-includes/functions.php", 7250, 7252},
		"wp_validate_boolean":                {"wp-includes/functions.php", 7932, 7942},
		"wp_array_slice_assoc":               {"wp-includes/functions.php", 5173, 5183},
		"wp_is_heic_image_mime_type":         {"wp-includes/functions.php", 9465, 9474},
		"bool_from_yn":                       {"wp-includes/functions.php", 1621, 1623},
		"wp_is_numeric_array":                {"wp-includes/functions.php", 5448, 5454},
		"_autop_newline_preservation_helper": {"wp-includes/formatting.php", 811, 813},
		"zeroise":                            {"wp-includes/formatting.php", 2848, 2850},
		"trailingslashit":                    {"wp-includes/formatting.php", 2883, 2885},
		"untrailingslashit":                  {"wp-includes/formatting.php", 2898, 2900},
		"stripslashes_from_strings_only":     {"wp-includes/formatting.php", 2934, 2936},
		"wp_slash":                           {"wp-includes/formatting.php", 5933, 5943},
		"_is_utf8_charset":                   {"wp-includes/compat.php", 102, 111},
		"array_all":                          {"wp-includes/compat.php", 635, 643},
		"array_all_guard":                    {"wp-includes/compat.php", 623, 644},
		"array_find":                         {"wp-includes/compat.php", 566, 574},
		"array_find_guard":                   {"wp-includes/compat.php", 554, 575},
		"array_find_key":                     {"wp-includes/compat.php", 589, 597},
		"array_find_key_guard":               {"wp-includes/compat.php", 577, 598},
		"array_any":                          {"wp-includes/compat.php", 612, 620},
		"array_any_guard":                    {"wp-includes/compat.php", 600, 621},
		"array_first":                        {"wp-includes/compat.php", 657, 665},
		"array_first_guard":                  {"wp-includes/compat.php", 646, 666},
		"array_last":                         {"wp-includes/compat.php", 679, 685},
		"array_last_guard":                   {"wp-includes/compat.php", 668, 686},
		"_wp_array_get":                      {"wp-includes/functions.php", 5233, 5272},
	}
	fragments := map[string][]byte{}
	var selections []map[string]any
	for _, source := range pin.Sources {
		if source.Path != "wp-includes/functions.php" && source.Path != "wp-includes/compat.php" && source.Path != "wp-includes/formatting.php" {
			t.Fatal("unexpected original source")
		}
		original := read(t, filepath.Join("../../..", source.Path))
		if digest(original) != source.Hash {
			t.Fatal("WordPress source drift; review original extraction")
		}
		lines := bytes.SplitAfter(original, []byte("\n"))
		for _, part := range source.Fragments {
			if expected, ok := wanted[part.Name]; !ok || expected != (selection{source.Path, part.First, part.Last}) || fragments[part.Name] != nil || part.Last > len(lines) {
				t.Fatal("unexpected/duplicate original fragment")
			}
			fragments[part.Name] = bytes.Join(lines[part.First-1:part.Last], nil)
			selections = append(selections, map[string]any{"name": part.Name, "source_path": source.Path, "source_sha256": source.Hash,
				"first_line": part.First, "last_line": part.Last, "sha256": digest(fragments[part.Name])})
		}
	}
	if len(fragments) != len(wanted) || !bytes.Contains(fragments["array_all_guard"], fragments["array_all"]) {
		t.Fatal("missing original helper or full guarded polyfill")
	}
	// Native retains the complete loading guard; only unchanged declarations compile.
	definitions, nativeDefinitions := []byte("<?php\n"), []byte("<?php\n")
	for _, name := range []string{"__return_true", "__return_false", "__return_zero", "__return_empty_array", "__return_null", "__return_empty_string", "wp_validate_boolean", "wp_array_slice_assoc", "wp_is_heic_image_mime_type", "bool_from_yn", "wp_is_numeric_array", "_autop_newline_preservation_helper", "zeroise", "trailingslashit", "untrailingslashit", "stripslashes_from_strings_only", "wp_slash", "_is_utf8_charset", "array_all", "array_any", "array_find", "array_find_key", "array_first", "array_last", "_wp_array_get"} {
		definitions = append(definitions, fragments[name]...)
		if strings.HasPrefix(name, "array_") {
			if !bytes.Contains(fragments[name+"_guard"], fragments[name]) {
				t.Fatal("incomplete original guarded polyfill")
			}
			name += "_guard"
		}
		nativeDefinitions = append(nativeDefinitions, fragments[name]...)
	}
	var drivers []byte
	driverHashes := map[string]string{}
	for _, name := range []string{"yn-bundle-drivers.php.txt", "numeric-array-drivers.php.txt", "numeric-bundle-drivers.php.txt", "array-search-drivers.php.txt", "search-bundle-drivers.php.txt", "edge-bundle-drivers.php.txt", "settings-bundle-drivers.php.txt"} {
		data := read(t, name)
		driverHashes[name] = digest(data)
		drivers = append(drivers, data...)
		drivers = append(drivers, '\n')
	}
	definitions = append(definitions, drivers...)
	nativeDefinitions = append(nativeDefinitions, drivers...)
	casesBytes := read(t, "numeric-array-cases.json")
	var cases []struct {
		Name           string `json:"name"`
		InputType      string `json:"input_type"`
		Input          any    `json:"input"`
		Expected       *bool  `json:"expected"`
		ExpectedIsList *bool  `json:"expected_is_list"`
		InputAfter     any    `json:"input_after"`
	}
	decoder := json.NewDecoder(bytes.NewReader(casesBytes))
	decoder.UseNumber()
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cases); err != nil {
		t.Fatal(err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		t.Fatal("trailing corpus data")
	}
	if len(cases) != 24 {
		t.Fatal("expected twenty-four native reference cases")
	}
	expected := make([]map[string]any, 0, len(cases))
	seen := map[string]bool{}
	for _, c := range cases {
		if c.Name == "" || seen[c.Name] || c.Expected == nil || (c.InputType == "array") != (c.ExpectedIsList != nil) {
			t.Fatal("invalid native reference expectation")
		}
		seen[c.Name] = true
		validateNumericArrayTransport(t, c.InputType, c.Input)
		validateNumericArrayTransport(t, c.InputType, c.InputAfter)
		expected = append(expected, map[string]any{"name": c.Name, "input_type": c.InputType, "type": "bool", "value": *c.Expected,
			"is_list": c.ExpectedIsList, "input_after": map[string]any{"type": c.InputType, "value": c.InputAfter}})
	}
	polyfillCases := read(t, "numeric-array-polyfill-cases.json")
	polyfillExpected := numericPolyfillExpected(t, polyfillCases)
	legacyCases := read(t, "yn-bundle-cases.json")
	compositionCases := read(t, "numeric-bundle-cases.json")
	searchCases := read(t, "array-search-cases.json")
	searchCompositionCases := read(t, "search-bundle-cases.json")
	edgeCases := read(t, "array-edge-cases.json")
	edgeCompositionCases := read(t, "edge-bundle-cases.json")
	lookupCases := read(t, "array-get-cases.json")
	settingsCases := read(t, "settings-bundle-cases.json")
	for _, retained := range []struct {
		data []byte
		hash string
	}{
		{legacyCases, "fbb7752edf4b84aa6205f6c50a07a37bee397d8da6ecedc9d2c6be5c4a4d6a32"},
		{casesBytes, "9806417cf73c2d67e7ddfbef8309ad9dbfbafefc83298d026c0310798c3e9ffb"},
		{polyfillCases, "a57aaa77efd06c73fac6e0c99728efd28711d7b686ae38e3d179e3a2347e2f30"},
		{compositionCases, "79cfd8fcd559943c339a98e4ecd0feaf00f891717f7800d7bea3b80618f7abc8"},
		{searchCases, "3ebab8acdcac68a054e8c6b6c54df9f21e5805b162fafd3067444a6cfe6f6216"},
		{searchCompositionCases, "ce5017329da976cf389467a81e8c53bd549639b3c79d37fe9913b68fc614d278"},
		{edgeCases, "f6108f0abfc56c533df2992568a6d3e011ef1e1db97c89af8e10164e0da0bd2e"},
		{edgeCompositionCases, "d053915decd72078c311e8d2a9fcc59e119bd943043155c2a3edea6fffba61bc"},
		{lookupCases, "b0624f70429ae85a400a536d77eb779626298277062f8f33b7642c9021b23463"},
	} {
		if digest(retained.data) != retained.hash {
			t.Fatal("retained corpus changed")
		}
	}
	legacyExpected := numericBundleLegacyExpected(t, legacyCases)
	compositionExpected := numericBundleCompositionExpected(t, compositionCases)
	searchExpected := arraySearchExpected(t, searchCases)
	searchCompositionExpected := searchBundleExpected(t, searchCompositionCases)
	retainedExpected := map[string]any{"legacy": legacyExpected, "numeric": map[string]any{"predicate": expected, "polyfill": polyfillExpected}, "composition": compositionExpected}
	previousExpected := map[string]any{"retained": retainedExpected, "search": searchExpected, "composition": searchCompositionExpected}
	edgeExpected := arrayEdgeExpected(t, edgeCases)
	edgeCompositionExpected := edgeBundleExpected(t, edgeCompositionCases)
	edgeBundleExpected := map[string]any{"retained": previousExpected, "edge": edgeExpected, "composition": edgeCompositionExpected}
	// Retained measured unit expectations include PHP-normalized ordered keys.
	// They are not substituted for new calls: both runtimes execute every case.
	lookupReferenceBytes := read(t, "../../../docs/evidence/array-get-unit-2026-10-26.json")
	var lookupReference map[string]any
	numericBundleDecode(t, lookupReferenceBytes, &lookupReference)
	lookupExpected, ok := lookupReference["typed_observations"].([]any)
	fragment, fragmentOK := lookupReference["fragment"].(map[string]any)
	if !ok || !fragmentOK || len(lookupExpected) != 24 || lookupReference["status"] != "differential_pass" ||
		lookupReference["schema"] != "gossip.array-get-unit/v1" || lookupReference["cases_sha256"] != digest(lookupCases) ||
		lookupReference["source_sha256"] != digest(read(t, "../../../wp-includes/functions.php")) || fragment["sha256"] != digest(fragments["_wp_array_get"]) {
		t.Fatal("lookup unit expectations drifted from unchanged source/corpus")
	}
	settingsExpected := settingsBundleExpected(t, settingsCases)
	expectedBytes, err := json.Marshal(map[string]any{"retained": edgeBundleExpected, "lookup": lookupExpected, "composition": settingsExpected})
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
	observerHashes := map[string]map[string]string{}
	adapt := func(name string, replacements ...string) []byte {
		original := read(t, name)
		data := append([]byte(nil), original...)
		for i := 0; i < len(replacements); i += 2 {
			if bytes.Count(data, []byte(replacements[i])) != 1 {
				t.Fatalf("adapter drift in %s: %q", name, replacements[i])
			}
			data = bytes.Replace(data, []byte(replacements[i]), []byte(replacements[i+1]), 1)
		}
		observerHashes[name] = map[string]string{"original_sha256": digest(original), "adapted_sha256": digest(data)}
		return data
	}
	legacyGo := adapt("yn-bundle-observer.go.txt",
		"func init() {", "func observeYNBundle() []ynBundleObservation {",
		"/input/cases.json", "/input/yn-cases.json",
		"\tif err := json.NewEncoder(os.Stdout).Encode(observations); err != nil {\n\t\tpanic(err)\n\t}\n\tos.Exit(0)",
		"\treturn observations")
	numericGo := adapt("numeric-array-observer.go.txt",
		"func init() {", "func observeNumericArray() numericArrayOutput {",
		"\tif err := json.NewEncoder(os.Stdout).Encode(numericArrayOutput{Predicate: predicate, Polyfill: polyfill}); err != nil {\n\t\tpanic(err)\n\t}\n\tos.Exit(0)",
		"\treturn numericArrayOutput{Predicate: predicate, Polyfill: polyfill}")
	legacyPHP := adapt("yn-bundle-observer.php", "require '/input/helpers.php';", "require_once '/input/helpers.php';", "/input/cases.json", "/input/yn-cases.json")
	reference := adapt("numeric-array-observer.php", "require '/input/helpers.php';", "require_once '/input/helpers.php';")
	numericPHP := adapt("numeric-array-differential-observer.php")
	retainedPHP := adapt("numeric-bundle-observer.php")
	retainedGo := adapt("numeric-bundle-observer.go.txt",
		"func init() {", "func observeNumericBundle() map[string]any {",
		"\n\t\"os\"", "",
		"\tif err := json.NewEncoder(os.Stdout).Encode(map[string]any{\"legacy\": legacy, \"numeric\": numeric, \"composition\": observations}); err != nil {\n\t\tpanic(err)\n\t}\n\tos.Exit(0)",
		"\treturn map[string]any{\"legacy\":legacy,\"numeric\":numeric,\"composition\":observations}")
	searchGo := adapt("array-search-observer.go.txt",
		"func init() {", "func observeSearchUnit() []map[string]any {",
		"\n\t\"os\"", "",
		"/input/cases.json", "/input/search-cases.json",
		"\tif err := json.NewEncoder(os.Stdout).Encode(observations); err != nil {\n\t\tpanic(err)\n\t}\n\tos.Exit(0)",
		"\treturn observations")
	searchPHP := adapt("array-search-observer.php",
		"require '/input/helpers.php';", "require_once '/input/helpers.php';",
		"require '/input/transport.php';", "// Shared transport already loaded by retained observer.",
		"/input/cases.json", "/input/search-cases.json")
	previousGo := adapt("search-bundle-observer.go.txt",
		"func init() {", "func observeSearchBundle() map[string]any {",
		"\n\t\"encoding/json\"", "",
		"\n\t\"os\"", "",
		"\tif err := json.NewEncoder(os.Stdout).Encode(map[string]any{\"retained\": retained, \"search\": search, \"composition\": observations}); err != nil {\n\t\tpanic(err)\n\t}\n\tos.Exit(0)",
		"\treturn map[string]any{\"retained\":retained,\"search\":search,\"composition\":observations}")
	edgeGo := adapt("array-edge-unit-observer.go.txt",
		"func init() {", "func observeEdgeUnit() []map[string]any {",
		"\n\t\"os\"", "",
		"/input/cases.json", "/input/edge-cases.json",
		"\tif err := json.NewEncoder(os.Stdout).Encode(observations); err != nil {\n\t\tpanic(err)\n\t}\n\tos.Exit(0)",
		"\treturn observations")
	previousPHP := adapt("search-bundle-observer.php")
	edgePHP := adapt("array-edge-observer.php",
		"require '/input/helpers.php';", "require_once '/input/helpers.php';",
		"require '/input/transport.php';", "// Shared transport already loaded.",
		"/input/cases.json", "/input/edge-cases.json")
	retainedEdgeGo := adapt("edge-bundle-observer.go.txt",
		"func init() {", "func observeEdgeBundle() map[string]any {",
		"\n\t\"encoding/json\"", "",
		"\n\t\"os\"", "",
		"\tif err := json.NewEncoder(os.Stdout).Encode(map[string]any{\"retained\": retained, \"edge\": edges, \"composition\": observations}); err != nil {\n\t\tpanic(err)\n\t}\n\tos.Exit(0)",
		"\treturn map[string]any{\"retained\": retained, \"edge\": edges, \"composition\": observations}")
	retainedEdgePHP := adapt("edge-bundle-observer.php")
	lookupGo := adapt("array-get-observer.go.txt",
		"func init() {", "func observeArrayGetUnit() []arrayGetObservation {",
		"/input/cases.json", "/input/lookup-cases.json",
		"\tif err := json.NewEncoder(os.Stdout).Encode(observations); err != nil {\n\t\tpanic(err)\n\t}\n\tos.Exit(0)",
		"\treturn observations")
	lookupPHP := adapt("array-get-observer.php",
		"require '/input/helpers.php';", "require_once '/input/helpers.php';",
		"/input/cases.json", "/input/lookup-cases.json")
	phpObserver := read(t, "settings-bundle-observer.php")
	goObserver := read(t, "settings-bundle-observer.go.txt")
	for name, data := range map[string][]byte{
		"helpers.php": nativeDefinitions, "unit.php": definitions, "cases.json": casesBytes,
		"polyfill-cases.json": polyfillCases, "yn-cases.json": legacyCases, "composition-cases.json": compositionCases,
		"reference.php": reference, "yn-observer.php": legacyPHP, "numeric-observer.php": numericPHP,
		"observer.php": phpObserver, "wordpress-license.txt": read(t, "../../../license.txt"),
		"retained-observer.php": retainedPHP, "search-observer.php": searchPHP,
		"search-cases.json": searchCases, "search-composition-cases.json": searchCompositionCases,
		"previous-observer.php": previousPHP, "edge-observer.php": edgePHP,
		"edge-cases.json": edgeCases, "edge-composition-cases.json": edgeCompositionCases,
		"retained-edge-observer.php": retainedEdgePHP, "lookup-observer.php": lookupPHP,
		"lookup-cases.json": lookupCases, "settings-composition-cases.json": settingsCases,
	} {
		write(t, filepath.Join(input, name), data)
	}
	write(t, filepath.Join(candidate, "legacy.go"), legacyGo)
	write(t, filepath.Join(candidate, "numeric.go"), numericGo)
	write(t, filepath.Join(candidate, "observer.go"), goObserver)
	write(t, filepath.Join(candidate, "retained.go"), retainedGo)
	write(t, filepath.Join(candidate, "search.go"), searchGo)
	write(t, filepath.Join(candidate, "previous.go"), previousGo)
	write(t, filepath.Join(candidate, "edge.go"), edgeGo)
	write(t, filepath.Join(candidate, "retained_edge.go"), retainedEdgeGo)
	write(t, filepath.Join(candidate, "lookup.go"), lookupGo)
	phpVersion := requirePHPOracle(t, pin.PHPImage)
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456",
		"--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local",
		"--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildOut, buildErr := container(t, "settings_bundle_generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/unit.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
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
	if report.Schema != "gohp.build/v0.1" || report.Status != "generated" || report.Target != pin.PHPTarget ||
		report.SourceHash != digest(definitions) || report.GeneratedHash != digest(generated) {
		t.Fatal("report does not match exact source/output bytes")
	}
	t.Log("status=generated; original declarations unchanged; top-level loading guard is not compiled")
	_, compileErr := container(t, "settings_bundle_build", append(buildArgs, "--entrypoint=go", pin.GoImage,
		"build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go", "/candidate/legacy.go", "/candidate/numeric.go", "/candidate/retained.go", "/candidate/search.go", "/candidate/previous.go", "/candidate/edge.go", "/candidate/retained_edge.go", "/candidate/lookup.go", "/candidate/observer.go")...)
	if len(compileErr) != 0 {
		t.Fatalf("build stderr: %s", compileErr)
	}
	t.Log("status=compiled; no compatibility claim before observations")
	actual, actualErr := container(t, "settings_bundle_candidate",
		"--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=/program", pin.GoImage)
	oracle, oracleErr := runPHPOracle(t, "settings_bundle_php", pin.PHPImage, input, "/input/observer.php")
	observed := jsonObservation(t, actual)
	if len(actualErr) != 0 || len(oracleErr) != 0 || !reflect.DeepEqual(observed, jsonObservation(t, oracle)) || !reflect.DeepEqual(observed, jsonObservation(t, expectedBytes)) {
		t.Fatalf("differential_failed: Go=%s stderr=%s PHP=%s stderr=%s expected=%s", actual, actualErr, oracle, oracleErr, expectedBytes)
	}
	evidence := map[string]any{
		"schema": pin.Schema, "status": "differential_pass", "manifest_sha256": digest(manifest),
		"wordpress_commit": pin.WordPressCommit, "fragments": selections, "original_declaration_count": 25, "owned_driver_count": 48,
		"compiled_bundle_sha256": digest(definitions), "native_bundle_sha256": digest(nativeDefinitions), "drivers_sha256": digest(drivers),
		"loading_guard": "native executes unchanged guard and verifies user-defined polyfill; compiler extracts unchanged declarations only",
		"cases_sha256":  digest(casesBytes), "polyfill_cases_sha256": digest(polyfillCases), "predicate_case_count": len(cases), "polyfill_case_count": len(polyfillExpected), "case_count": len(cases) + len(polyfillExpected) + len(legacyExpected) + len(compositionExpected) + len(searchExpected) + len(searchCompositionExpected) + len(edgeExpected) + len(edgeCompositionExpected) + len(lookupExpected) + len(settingsExpected),
		"lookup_case_count": len(lookupExpected), "settings_composition_case_count": len(settingsExpected),
		"lookup_cases_sha256": digest(lookupCases), "settings_cases_sha256": digest(settingsCases), "lookup_unit_evidence_sha256": digest(lookupReferenceBytes),
		"edge_case_count": len(edgeExpected), "edge_composition_case_count": len(edgeCompositionExpected), "edge_cases_sha256": digest(edgeCases), "edge_composition_cases_sha256": digest(edgeCompositionCases),
		"search_case_count": len(searchExpected), "search_composition_case_count": len(searchCompositionExpected),
		"search_cases_sha256": digest(searchCases), "search_composition_cases_sha256": digest(searchCompositionCases),
		"legacy_case_count": len(legacyExpected), "composition_case_count": len(compositionExpected),
		"legacy_cases_sha256": digest(legacyCases), "composition_cases_sha256": digest(compositionCases), "driver_files": driverHashes, "observer_adapters": observerHashes,
		"gohp_commit": pin.GoHPCommit, "gohp_archive_sha256": archiveHash, "compiler_version": report.CompilerVersion,
		"compiler_sha256": digest(read(t, filepath.Join(candidate, "gohp"))), "generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"go_image": pin.GoImage, "go_version": strings.TrimSpace(string(buildOut)), "php_image": pin.PHPImage, "php_version": phpVersion, "php_target": pin.PHPTarget, "php_minimum": phporacle.Minimum, "php_lc_ctype": "C",
		"go_observer_sha256": digest(goObserver), "php_observer_sha256": digest(phpObserver), "reference_observer_sha256": digest(reference),
		"binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "typed_observations": observed, "stderr": "", "exit_code": 0,
		"scope": "twenty-five unchanged original declarations and48 owned drivers in one binary;262 retained edge-bundle cases,24 retained nested lookup cases,8 new settings compositions; actual typed results and caller inputs; bounded fixture domains only; no full WordPress runtime compatibility",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}

func settingsBundleExpected(t *testing.T, data []byte) []map[string]any {
	t.Helper()
	var cases []struct {
		Name      string `json:"name"`
		Operation string `json:"operation"`
		Input     any    `json:"input"`
		Path      any    `json:"path"`
		Default   any    `json:"default"`
		Selector  any    `json:"selector"`
		Expected  any    `json:"expected"`
	}
	numericBundleDecode(t, data, &cases)
	if len(cases) != 8 {
		t.Fatal("expected eight settings compositions")
	}
	seen, counts := map[string]bool{}, map[string]int{}
	var observations []map[string]any
	for _, c := range cases {
		if c.Name == "" || seen[c.Name] {
			t.Fatal("invalid settings composition name")
		}
		seen[c.Name] = true
		kind := ""
		switch c.Operation {
		case "gossip_settings_bool":
			kind = "bool"
			if _, ok := c.Expected.(bool); !ok {
				t.Fatal("expected bool composition result")
			}
		case "gossip_settings_edges":
			kind = "array"
			if _, ok := c.Expected.(map[string]any); !ok {
				t.Fatal("expected ordered array composition result")
			}
		default:
			t.Fatal("unknown settings composition")
		}
		counts[c.Operation]++
		observations = append(observations, map[string]any{"name": c.Name, "type": kind, "value": c.Expected,
			"args_after": []any{c.Input, c.Path, c.Default, c.Selector}})
	}
	if counts["gossip_settings_bool"] != 4 || counts["gossip_settings_edges"] != 4 {
		t.Fatal("expected four cases per settings composition")
	}
	return observations
}
