package gohp_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// This opt-in experiment is not a GoHP runtime or a general preg_match adapter.
// It statically links an already installed, byte-pinned PCRE2 library; no fetch.
func TestParserPCRE2Comparison(t *testing.T) {
	prefix := os.Getenv("GOSSIP_PCRE2_PREFIX")
	if prefix == "" {
		t.Skip("set GOSSIP_PCRE2_PREFIX for the pinned darwin/arm64 PCRE2 10.47 experiment")
	}
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" || !filepath.IsAbs(prefix) {
		t.Fatal("this recorded build profile requires darwin/arm64 and an absolute prefix")
	}
	header := filepath.Join(prefix, "include/pcre2.h")
	library := filepath.Join(prefix, "lib/libpcre2-8.a")
	const headerHash = "058462dc030e845ee627e19ac36347f561f5cc44d89eaf1c0010f2b363ab68d5"
	const libraryHash = "ce40246eba1dfdd99e43142cfc33264610d7c965fab3062f386c68a6636a4a54"
	if digest(read(t, header)) != headerHash || digest(read(t, library)) != libraryHash {
		t.Fatal("PCRE2 header/static archive differs from the measured profile")
	}
	referenceBytes := read(t, "../../../docs/evidence/parser-token-reference-2026-10-27.json")
	if digest(referenceBytes) != "5cbf560247afcd2c5f1462cab2f95f16bb822ac2d4d1756be16221cffdc6dc60" {
		t.Fatal("retained PHP reference changed")
	}
	var reference struct {
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
	if err := json.Unmarshal(referenceBytes, &reference); err != nil {
		t.Fatal(err)
	}
	if reference.Status != "native-reference-pass" || len(reference.Native.Cases) != 24 {
		t.Fatal("incomplete reference")
	}
	for path, hash := range reference.Sources {
		if digest(read(t, filepath.Join("../../..", path))) != hash {
			t.Fatalf("original source drift: %s", path)
		}
	}
	pattern, err := base64.StdEncoding.Strict().DecodeString(reference.Native.Pattern)
	if err != nil || digest(pattern) != reference.PatternSHA || !bytes.HasPrefix(pattern, []byte("/")) || !bytes.HasSuffix(pattern, []byte("/s")) {
		t.Fatal("fixed pattern/delimiter drift")
	}
	casesBytes := read(t, "parser-token-reference-cases.json")
	if digest(casesBytes) != reference.CasesSHA {
		t.Fatal("corpus drift")
	}
	var cases []struct {
		Name     string `json:"name"`
		Document string `json:"document_base64"`
		Offset   int    `json:"offset"`
		Limit    int    `json:"backtrack_limit"`
	}
	if err := json.Unmarshal(casesBytes, &cases); err != nil {
		t.Fatal(err)
	}
	if len(cases) != 24 {
		t.Fatal("requires fixed 24 cases")
	}
	input := t.TempDir()
	patternPath := filepath.Join(input, "pattern.bin")
	// Only PHP delimiters and /s are removed; DOTALL represents that modifier.
	write(t, patternPath, pattern[1:len(pattern)-2])
	source := read(t, "testdata/parser-pcre2-probe.c")
	sourcePath := filepath.Join(input, "probe.c")
	write(t, sourcePath, source)
	binary := filepath.Join(input, "probe")
	compiler, err := exec.LookPath("clang")
	if err != nil {
		t.Fatal(err)
	}
	version := pcre2Command(t, 10*time.Second, compiler, "--version")
	versionLine := strings.SplitN(string(version), "\n", 2)[0] // No local installation paths.
	pcre2Command(t, 30*time.Second, compiler, "-std=c11", "-O2", "-Wall", "-Wextra", "-Werror", "-I"+filepath.Join(prefix, "include"), sourcePath, library, "-o", binary)
	observations := make([]any, 0, len(cases))
	mismatches := 0
	for i, c := range cases {
		if c.Name != reference.Native.Cases[i].Name {
			t.Fatal("case identity drift")
		}
		subject, err := base64.StdEncoding.Strict().DecodeString(c.Document)
		if err != nil || len(subject) > 4096 || c.Offset < -len(subject) || c.Offset > len(subject) || (c.Limit != 1 && c.Limit != 100000) {
			t.Fatal("case outside fixed domain")
		}
		offset := c.Offset
		if offset < 0 {
			offset += len(subject)
		} // PHP's in-range negative byte offset.
		subjectPath := filepath.Join(input, "subject.bin")
		write(t, subjectPath, subject)
		stdout := pcre2Command(t, 5*time.Second, binary, patternPath, subjectPath, strconv.Itoa(offset), strconv.Itoa(c.Limit))
		var raw pcre2Probe
		if err := json.Unmarshal(stdout, &raw); err != nil {
			t.Fatal(err)
		}
		if !strings.HasPrefix(raw.Version, "10.47 ") || raw.Header != "10.47" || raw.JIT || raw.MatchLimit != c.Limit || raw.DepthLimit != 10000 || raw.HeapLimit != 131072 || raw.CompileOptions != 32 || raw.Newline != 2 || raw.BSR != 1 || raw.CaptureCount != 5 {
			t.Fatalf("candidate profile drift: %+v", raw)
		}
		actual, err := pcre2PHPShape(raw, subject)
		if err != nil {
			t.Fatal(err)
		}
		// PHP's wrapper error text is not an engine result. Compare codes separately.
		expected := reference.Native.Cases[i].Preg
		delete(expected, "error_message")
		actualJSON, err := json.Marshal(actual)
		if err != nil {
			t.Fatal(err)
		}
		equal := reflect.DeepEqual(jsonObservation(t, actualJSON), expected)
		if !equal {
			mismatches++
		}
		observations = append(observations, map[string]any{"name": c.Name, "input_base64": c.Document, "php_offset": c.Offset, "normalized_offset": offset, "raw": raw, "php_shaped": actual, "equal_to_retained_preg_match": equal, "stdout_sha256": digest(stdout), "stderr_bytes": 0, "exit_code": 0})
	}
	status := "bounded-engine-comparison-pass"
	if mismatches != 0 {
		status = "bounded-engine-comparison-mismatch"
	}
	evidence := map[string]any{
		"schema": "gossip.parser-pcre2-comparison/v1", "status": status, "cases": observations, "matches": 24 - mismatches, "mismatches": mismatches,
		"reference_sha256": digest(referenceBytes), "cases_sha256": digest(casesBytes), "pattern_sha256": digest(pattern), "pattern_body_sha256": digest(pattern[1 : len(pattern)-2]),
		"probe_source_sha256": digest(source), "harness_sha256": digest(read(t, "parser_pcre2_comparison_test.go")), "header_sha256": headerHash, "static_archive_sha256": libraryHash, "executable_sha256": digest(read(t, binary)),
		"compiler": versionLine, "compiler_sha256": digest(read(t, compiler)), "compile_arguments": []string{"-std=c11", "-O2", "-Wall", "-Wextra", "-Werror", "-I<PCRE2_PREFIX>/include", "<INPUT>/probe.c", "<PCRE2_PREFIX>/lib/libpcre2-8.a", "-o", "<OUTPUT>/probe"},
		"go_version": runtime.Version(), "os": runtime.GOOS, "arch": runtime.GOARCH, "locale": "C", "php_reference_version": "8.3.36", "php_reference_pcre": "10.42 2022-12-12",
		"reference_reused": true, "next_token_status": "not-run", "gohp_status": "not-run", "cgo_status": "not-run", "regexp2_status": "not-run", "php84_status": "not-run",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("PCRE2_COMPARISON_EVIDENCE " + string(encoded))
	if mismatches != 0 {
		t.Fatalf("%d candidate cases differ; retained report includes both raw and PHP-shaped results", mismatches)
	}
}

type pcre2Probe struct {
	Version        string `json:"version"`
	Header         string `json:"header_version"`
	JIT            bool   `json:"jit"`
	MatchLimit     int    `json:"match_limit"`
	DepthLimit     int    `json:"depth_limit"`
	HeapLimit      int    `json:"heap_limit_kib"`
	CompileOptions int    `json:"compile_options"`
	Newline        int    `json:"newline"`
	BSR            int    `json:"bsr"`
	RC             int    `json:"return_code"`
	CaptureCount   int    `json:"capture_count"`
	Names          []struct {
		Number int    `json:"number"`
		Name   string `json:"name"`
	} `json:"names"`
	Captures []struct {
		Number int    `json:"number"`
		Start  int    `json:"start"`
		End    int    `json:"end"`
		Hex    string `json:"hex"`
	} `json:"captures"`
}

// A fixture-only mapping, not a general PHP wrapper. No expected values are used.
func pcre2PHPShape(raw pcre2Probe, subject []byte) (map[string]any, error) {
	names := make(map[int]string)
	wanted := map[int]string{1: "closer", 2: "namespace", 3: "name", 4: "attrs", 5: "void"}
	for _, n := range raw.Names {
		if _, exists := names[n.Number]; exists {
			return nil, fmt.Errorf("duplicate name entry")
		}
		names[n.Number] = n.Name
	}
	if !reflect.DeepEqual(names, wanted) {
		return nil, fmt.Errorf("unexpected original named captures")
	}
	entries := []any{}
	presence := map[string]bool{"closer": false, "namespace": false, "name": false, "attrs": false, "void": false}
	result := pcre2Scalar("int", 0)
	errorCode := 0
	switch {
	case raw.RC == -1: // PCRE2_ERROR_NOMATCH
	case raw.RC == -47: // PCRE2_ERROR_MATCHLIMIT, mapped to PREG_BACKTRACK_LIMIT_ERROR
		result, errorCode = pcre2Scalar("bool", false), 2
	case raw.RC > 0 && raw.RC <= 6:
		result = pcre2Scalar("int", 1)
		if len(raw.Captures) != raw.RC {
			return nil, fmt.Errorf("incomplete capture vector")
		}
		for i, c := range raw.Captures {
			data, err := hex.DecodeString(c.Hex)
			if err != nil || c.Number != i {
				return nil, fmt.Errorf("invalid capture")
			}
			if c.Start == -1 {
				if c.End != -1 || len(data) != 0 {
					return nil, fmt.Errorf("invalid unmatched capture")
				}
			} else if c.Start < 0 || c.End < c.Start || c.End > len(subject) || !bytes.Equal(data, subject[c.Start:c.End]) {
				return nil, fmt.Errorf("invalid byte range")
			}
			value := pcre2Array([]any{pcre2Entry(pcre2Scalar("int", 0), map[string]any{"type": "string", "base64": base64.StdEncoding.EncodeToString(data)}), pcre2Entry(pcre2Scalar("int", 1), pcre2Scalar("int", c.Start))})
			if name, ok := names[i]; ok {
				presence[name] = true
				entries = append(entries, pcre2Entry(map[string]any{"type": "string", "base64": base64.StdEncoding.EncodeToString([]byte(name))}, value))
			}
			entries = append(entries, pcre2Entry(pcre2Scalar("int", i), value))
		}
	default:
		return nil, fmt.Errorf("unmapped PCRE2 result: %d", raw.RC)
	}
	if raw.RC < 0 && len(raw.Captures) != 0 {
		return nil, fmt.Errorf("captures on failure")
	}
	return map[string]any{"return": result, "error_code": errorCode, "matches_before": pcre2Scalar("null", nil), "matches_after": pcre2Array(entries), "named_capture_presence": presence}, nil
}

func pcre2Scalar(kind string, value any) map[string]any {
	return map[string]any{"type": kind, "value": value}
}
func pcre2Array(entries []any) map[string]any {
	return map[string]any{"type": "array", "entries": entries}
}
func pcre2Entry(key, value any) map[string]any { return map[string]any{"key": key, "value": value} }

func pcre2Command(t *testing.T, limit time.Duration, executable string, args ...string) []byte {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()
	cmd := exec.CommandContext(ctx, executable, args...)
	cmd.Env = []string{"PATH=/usr/bin:/bin", "LC_ALL=C", "LANG=C", "TMPDIR=" + t.TempDir()}
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil || stderr.Len() != 0 {
		t.Fatalf("candidate infrastructure/diagnostic failure (not no-match): %v; %s", err, stderr.String())
	}
	return stdout.Bytes()
}
