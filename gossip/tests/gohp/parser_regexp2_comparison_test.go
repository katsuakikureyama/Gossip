package gohp_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

const regexp2Revision = "a6ef7b6f88fff76c0c7d4f5fe832535de8cfda11"

// A successful test means complete observations, NOT backend compatibility.
func TestParserRegexp2Comparison(t *testing.T) {
	repository, goTool := os.Getenv("GOSSIP_REGEXP2_SOURCE"), os.Getenv("GOSSIP_REGEXP2_GO")
	if repository == "" || goTool == "" {
		t.Skip("set GOSSIP_REGEXP2_SOURCE and GOSSIP_REGEXP2_GO for isolated regexp2 evaluation")
	}
	if !filepath.IsAbs(repository) || !filepath.IsAbs(goTool) {
		t.Fatal("requires absolute source/toolchain paths")
	}
	referenceBytes := read(t, "../../../docs/evidence/parser-token-reference-2026-10-27.json")
	if digest(referenceBytes) != "5cbf560247afcd2c5f1462cab2f95f16bb822ac2d4d1756be16221cffdc6dc60" {
		t.Fatal("reference drift")
	}
	var ref struct {
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
	if err := json.Unmarshal(referenceBytes, &ref); err != nil {
		t.Fatal(err)
	}
	for p, h := range ref.Sources {
		if digest(read(t, filepath.Join("../../..", p))) != h {
			t.Fatal("source drift: " + p)
		}
	}
	pattern, err := base64.StdEncoding.Strict().DecodeString(ref.Native.Pattern)
	if err != nil || digest(pattern) != ref.PatternSHA || !bytes.HasPrefix(pattern, []byte("/")) || !bytes.HasSuffix(pattern, []byte("/s")) {
		t.Fatal("pattern drift")
	}
	caseBytes := read(t, "parser-token-reference-cases.json")
	if digest(caseBytes) != ref.CasesSHA {
		t.Fatal("corpus drift")
	}
	var cases []struct {
		Name     string `json:"name"`
		Document string `json:"document_base64"`
		Offset   int    `json:"offset"`
		Limit    int    `json:"backtrack_limit"`
	}
	if err := json.Unmarshal(caseBytes, &cases); err != nil {
		t.Fatal(err)
	}
	if len(cases) != 24 || len(ref.Native.Cases) != 24 {
		t.Fatal("requires 24 retained cases")
	}
	root := t.TempDir()
	backend := filepath.Join(root, "backend")
	// Reuse the shared committed-git-archive exporter: no dirty checkout bytes.
	archiveSHA := exportGoHP(t, repository, regexp2Revision, backend)
	module := read(t, filepath.Join(backend, "go.mod"))
	if string(module) != "module github.com/dlclark/regexp2/v2\n\ngo 1.25\n" {
		t.Fatal("upstream module contract changed")
	}
	probe := read(t, "testdata/parser-regexp2-probe.go")
	probeDir := filepath.Join(root, "probe")
	write(t, filepath.Join(probeDir, "main.go"), probe)
	mod := []byte("module gossip.local/regexp2-probe\n\ngo 1.25\n\nrequire github.com/dlclark/regexp2/v2 v2.0.0\n\nreplace github.com/dlclark/regexp2/v2 => ../backend\n")
	write(t, filepath.Join(probeDir, "go.mod"), mod)
	version := regexp2Command(t, 10*time.Second, probeDir, goTool, "version")
	if string(version) != "go version go1.25.14 darwin/arm64\n" {
		t.Fatalf("unadmitted probe toolchain: %s", version)
	}
	binary := filepath.Join(root, "probe-bin")
	regexp2Command(t, 60*time.Second, probeDir, goTool, "build", "-trimpath", "-buildvcs=false", "-o", binary, ".")
	patternPath := filepath.Join(root, "pattern.bin")
	write(t, patternPath, pattern[1:len(pattern)-2])
	var observations []any
	counts := map[string]int{"equal": 0, "different": 0, "not-comparable": 0}
	for i, c := range cases {
		if c.Name != ref.Native.Cases[i].Name {
			t.Fatal("case order drift")
		}
		subject, err := base64.StdEncoding.Strict().DecodeString(c.Document)
		if err != nil || len(subject) > 4096 || c.Offset < -len(subject) || c.Offset > len(subject) || (c.Limit != 1 && c.Limit != 100000) {
			t.Fatal("invalid case domain")
		}
		offset := c.Offset
		if offset < 0 {
			offset += len(subject)
		}
		subjectPath := filepath.Join(root, "subject.bin")
		write(t, subjectPath, subject)
		out := regexp2Command(t, 5*time.Second, root, binary, patternPath, subjectPath, strconv.Itoa(offset))
		var raw regexp2Observation
		if err := json.Unmarshal(out, &raw); err != nil {
			t.Fatal(err)
		}
		if raw.Options != 528 || raw.Stack != 100000 || raw.Timeout != 100 || raw.GoVersion != "go1.25.14" {
			t.Fatalf("profile drift: %+v", raw)
		}
		shape, err := regexp2PHPShape(raw, subject)
		if err != nil {
			t.Fatal(err)
		}
		expected := ref.Native.Cases[i].Preg
		delete(expected, "error_message")
		result, reason := "different", "candidate output or API behavior differs"
		if c.Limit == 1 {
			result, reason = "not-comparable", "PCRE match-count limit has no equivalent in the selected regexp2 stack/time profile"
		} else if shape != nil {
			encoded, err := json.Marshal(shape)
			if err != nil {
				t.Fatal(err)
			}
			if reflect.DeepEqual(jsonObservation(t, encoded), expected) {
				result, reason = "equal", "observed output equality only; resource-limit contracts are not equivalent"
			}
		}
		counts[result]++
		observations = append(observations, map[string]any{"name": c.Name, "input_base64": c.Document, "php_offset": c.Offset, "normalized_offset": offset, "php_match_limit": c.Limit, "raw": raw, "php_shaped": shape, "comparison": result, "reason": reason, "stdout_sha256": digest(out), "stderr_bytes": 0, "exit_code": 0})
	}
	report := map[string]any{
		"schema": "gossip.parser-regexp2-comparison/v1", "status": "comparison-observed", "compatibility_status": "not-established", "counts": counts, "cases": observations,
		"regexp2_commit": regexp2Revision, "source_archive_sha256": archiveSHA, "module_sha256": digest(module), "license_sha256": digest(read(t, filepath.Join(backend, "LICENSE"))), "temporary_module_sha256": digest(mod),
		"reference_sha256": digest(referenceBytes), "cases_sha256": digest(caseBytes), "pattern_sha256": digest(pattern), "probe_sha256": digest(probe), "harness_sha256": digest(read(t, "parser_regexp2_comparison_test.go")), "archive_helper_sha256": digest(read(t, "helpers_test.go")), "executable_sha256": digest(read(t, binary)), "go_tool_sha256": digest(read(t, goTool)), "go_version": strings.TrimSpace(string(version)),
		"profile": "RE2|Singleline; string API; stack slots100000; timeout100ms; no PCRE match-count equivalent", "source_rewrite": false, "php_reference_reused": true, "php84_status": "not-run", "gohp_status": "not-run", "next_token_status": "not-run",
	}
	encoded, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("REGEXP2_COMPARISON_EVIDENCE " + string(encoded))
}

type regexp2Observation struct {
	Status     string `json:"status"`
	Error      string `json:"error"`
	ErrorClass string `json:"error_class"`
	Options    int    `json:"options"`
	Stack      int    `json:"stack_slots"`
	Timeout    int    `json:"timeout_ms"`
	GoVersion  string `json:"go_version"`
	Groups     []struct {
		Number       int    `json:"number"`
		Name         string `json:"name"`
		Participated bool   `json:"participated"`
		RuneIndex    int    `json:"rune_index"`
		RuneLength   int    `json:"rune_length"`
		Start        int    `json:"byte_start"`
		Length       int    `json:"byte_length"`
		Base64       string `json:"base64"`
	} `json:"groups"`
}

func regexp2PHPShape(raw regexp2Observation, subject []byte) (map[string]any, error) {
	if raw.Status == "error" {
		if raw.Error == "" || raw.ErrorClass == "" {
			return nil, fmt.Errorf("missing error evidence")
		}
		return nil, nil
	}
	if raw.Error != "" || raw.ErrorClass != "" {
		return nil, fmt.Errorf("unexpected diagnostics")
	}
	entries := []any{}
	presence := map[string]bool{"closer": false, "namespace": false, "name": false, "attrs": false, "void": false}
	result := pcre2Scalar("int", 0)
	if raw.Status == "match" {
		if len(raw.Groups) != 6 {
			return nil, fmt.Errorf("group count drift")
		}
		names := []string{"0", "closer", "namespace", "name", "attrs", "void"}
		last := 0
		for i, g := range raw.Groups {
			if g.Number != i || g.Name != names[i] {
				return nil, fmt.Errorf("group identity drift")
			}
			data, err := base64.StdEncoding.Strict().DecodeString(g.Base64)
			if err != nil {
				return nil, err
			}
			if g.Participated {
				if g.Start < 0 || g.Length < 0 || g.Start > len(subject)-g.Length || !bytes.Equal(data, subject[g.Start:g.Start+g.Length]) {
					return nil, fmt.Errorf("capture byte range drift")
				}
				last = i
			} else if g.Start != -1 || g.Length != 0 || len(data) != 0 {
				return nil, fmt.Errorf("unmatched group drift")
			}
		}
		if !raw.Groups[0].Participated {
			return nil, fmt.Errorf("missing whole match")
		}
		for i, g := range raw.Groups[:last+1] {
			value := pcre2Array([]any{pcre2Entry(pcre2Scalar("int", 0), map[string]any{"type": "string", "base64": g.Base64}), pcre2Entry(pcre2Scalar("int", 1), pcre2Scalar("int", g.Start))})
			if i > 0 {
				presence[g.Name] = true
				entries = append(entries, pcre2Entry(map[string]any{"type": "string", "base64": base64.StdEncoding.EncodeToString([]byte(g.Name))}, value))
			}
			entries = append(entries, pcre2Entry(pcre2Scalar("int", i), value))
		}
		result = pcre2Scalar("int", 1)
	} else if raw.Status != "no-match" || len(raw.Groups) != 0 {
		return nil, fmt.Errorf("unknown outcome: %s", raw.Status)
	}
	return map[string]any{"return": result, "error_code": 0, "matches_before": pcre2Scalar("null", nil), "matches_after": pcre2Array(entries), "named_capture_presence": presence}, nil
}

func regexp2Command(t *testing.T, limit time.Duration, dir, exe string, args ...string) []byte {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()
	cmd := exec.CommandContext(ctx, exe, args...)
	cmd.Dir = dir
	cmd.Env = []string{"PATH=/usr/bin:/bin", "LC_ALL=C", "GOTOOLCHAIN=local", "GOPROXY=off", "GOSUMDB=off", "GOWORK=off", "CGO_ENABLED=0", "GOCACHE=" + filepath.Join(os.TempDir(), "gossip-regexp2-go-cache"), "TMPDIR=" + t.TempDir()}
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil || stderr.Len() != 0 {
		t.Fatalf("probe infrastructure/diagnostic failure, not regex mismatch: %v %s", err, stderr.String())
	}
	return stdout.Bytes()
}
