package gohp_test

import (
	"archive/tar"
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

type fragment struct {
	Name  string `json:"name"`
	First int    `json:"first_line"`
	Last  int    `json:"last_line"`
}

type unitPin struct {
	Schema          string     `json:"schema"`
	WordPressCommit string     `json:"wordpress_commit"`
	SourcePath      string     `json:"source_path"`
	SourceHash      string     `json:"source_sha256"`
	GoHPCommit      string     `json:"gohp_commit"`
	GoImage         string     `json:"go_image"`
	PHPImage        string     `json:"php_image"`
	PHPTarget       string     `json:"php_target"`
	Fragments       []fragment `json:"fragments"`
	Expected        string     `json:"expected_stdout"`
}

func digest(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func read(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func write(t *testing.T, path string, data []byte) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}

// Only immutable committed GoHP files are exported. The caller's working tree,
// untracked files, .git configuration and credentials are never mounted.
func exportGoHP(t *testing.T, repository, revision, target string) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "-C", repository, "archive", "--format=tar", revision)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	data, err := cmd.Output()
	if err != nil {
		t.Fatalf("status=infrastructure_failed: export pinned GoHP: %v %s", err, stderr.String())
	}
	if len(data) > 16<<20 {
		t.Fatal("GoHP source archive exceeds fixture budget")
	}
	archive := tar.NewReader(bytes.NewReader(data))
	for {
		header, err := archive.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if header.Typeflag == tar.TypeXGlobalHeader {
			continue
		}
		if !filepath.IsLocal(header.Name) {
			t.Fatal("nonlocal archive path")
		}
		path := filepath.Join(target, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				t.Fatal(err)
			}
		case tar.TypeReg:
			contents, err := io.ReadAll(archive)
			if err != nil {
				t.Fatal(err)
			}
			write(t, path, contents)
		default:
			t.Fatalf("unsupported archive entry type: %d", header.Typeflag)
		}
	}
	return digest(data)
}

// Docker executes only fixed test commands. Each invocation owns its random
// container name, so a timeout can clean up that exact container, not others.
func container(t *testing.T, stage string, args ...string) ([]byte, []byte) {
	t.Helper()
	id := make([]byte, 8)
	if _, err := rand.Read(id); err != nil {
		t.Fatal(err)
	}
	name := "gossip-gohp-" + hex.EncodeToString(id)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	base := []string{"run", "--name", name, "--rm", "--pull=never", "--network=none", "--read-only",
		"--cap-drop=ALL", "--security-opt=no-new-privileges", "--user", fmt.Sprintf("%d:%d", os.Getuid(), os.Getgid()),
		"--memory=768m", "--pids-limit=256", "--cpus=2"}
	command := exec.CommandContext(ctx, "docker", append(base, args...)...)
	var stdout, stderr bytes.Buffer
	command.Stdout, command.Stderr = &stdout, &stderr
	err := command.Run()
	if ctx.Err() != nil {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		cleanup, cleanupErr := exec.CommandContext(cleanupCtx, "docker", "rm", "--force", name).CombinedOutput()
		t.Fatalf("status=infrastructure_failed: %s timeout; cleanup=%v %s", stage, cleanupErr, cleanup)
	}
	if err != nil {
		t.Fatalf("status=%s_failed: %v\nstdout=%s\nstderr=%s", stage, err, stdout.Bytes(), stderr.Bytes())
	}
	return stdout.Bytes(), stderr.Bytes()
}

func TestOriginalWordPressReturnHelpers(t *testing.T) {
	repository := os.Getenv("GOSSIP_GOHP_SOURCE")
	if repository == "" {
		t.Skip("set GOSSIP_GOHP_SOURCE to a local GoHP clone containing the pinned commit; cached pinned Docker images are also required")
	}
	if os.Getuid() == 0 {
		t.Fatal("run this local fixture as a non-root user")
	}
	var pin unitPin
	if err := json.Unmarshal(read(t, "return-helpers.json"), &pin); err != nil {
		t.Fatal(err)
	}
	if pin.Schema != "gossip.wordpress-unit/v1" || pin.SourcePath != "wp-includes/functions.php" || len(pin.Fragments) != 5 || pin.PHPTarget != "8.3" ||
		!regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.GoHPCommit) ||
		!regexp.MustCompile(`^golang@sha256:[0-9a-f]{64}$`).MatchString(pin.GoImage) ||
		!regexp.MustCompile(`^wordpress@sha256:[0-9a-f]{64}$`).MatchString(pin.PHPImage) {
		t.Fatal("invalid unit pin")
	}
	// Tie this unit to the fork baseline as well as to the whole original file.
	var baseline struct {
		Commit string `json:"commit"`
	}
	if err := json.Unmarshal(read(t, "../../wordpress-source.json"), &baseline); err != nil {
		t.Fatal(err)
	}
	if baseline.Commit != pin.WordPressCommit {
		t.Fatal("unit and fork baseline disagree")
	}
	original := read(t, filepath.Join("../../..", pin.SourcePath))
	if digest(original) != pin.SourceHash {
		t.Fatal("original WordPress source drift; review before changing extraction ranges")
	}
	lines := bytes.SplitAfter(original, []byte("\n"))
	var bundle bytes.Buffer
	bundle.WriteString("<?php\n")
	var selections []map[string]any
	for _, part := range pin.Fragments {
		if part.First < 1 || part.Last < part.First || part.Last > len(lines) {
			t.Fatal("invalid fragment range")
		}
		excerpt := bytes.Join(lines[part.First-1:part.Last], nil)
		if !bytes.HasPrefix(excerpt, []byte("function "+part.Name+"() {")) || !bytes.HasSuffix(excerpt, []byte("}\n")) {
			t.Fatal("fragment does not contain the expected complete original function")
		}
		selections = append(selections, map[string]any{"name": part.Name, "original_first_line": part.First,
			"original_last_line": part.Last, "bundle_first_line": bytes.Count(bundle.Bytes(), []byte("\n")) + 1, "sha256": digest(excerpt)})
		bundle.Write(excerpt) // No rewriting of whitespace, comments or function body.
		bundle.WriteByte('\n')
	}
	bundle.Write(read(t, "calls.php")) // Owned observation calls, not WordPress implementation.

	workspace := t.TempDir()
	source := filepath.Join(workspace, "gohp")
	input := filepath.Join(workspace, "input")
	candidate := filepath.Join(workspace, "candidate")
	for _, dir := range []string{source, input, candidate} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}
	archiveHash := exportGoHP(t, repository, pin.GoHPCommit, source)
	write(t, filepath.Join(input, "unit.php"), bundle.Bytes())
	write(t, filepath.Join(input, "wordpress-license.txt"), read(t, "../../../license.txt"))
	buildArgs := []string{"--mount", "type=bind,source=" + source + ",target=/gohp,readonly",
		"--mount", "type=bind,source=" + input + ",target=/input,readonly",
		"--mount", "type=bind,source=" + candidate + ",target=/candidate",
		"--workdir=/gohp", "--tmpfs", "/tmp:rw,nosuid,noexec,size=268435456",
		"--env=GOCACHE=/candidate/cache", "--env=GOPATH=/candidate/gopath", "--env=GOTOOLCHAIN=local",
		"--env=GOPROXY=off", "--env=GOSUMDB=off", "--env=GOWORK=off", "--env=CGO_ENABLED=0", "--env=GOFLAGS=-mod=readonly"}
	buildStdout, buildStderr := container(t, "generation", append(buildArgs, "--entrypoint=sh", pin.GoImage, "-eu", "-c",
		"go version; go build -trimpath -o /candidate/gohp ./cmd/gohp; /candidate/gohp -in /input/unit.php -out /candidate/main.go -php-version 8.3 -report /candidate/report.json")...)
	if len(buildStderr) != 0 {
		t.Fatalf("unexpected generator stderr: %s", buildStderr)
	}
	generated := read(t, filepath.Join(candidate, "main.go"))
	reportBytes := read(t, filepath.Join(candidate, "report.json"))
	var report struct {
		Schema          string `json:"schema"`
		Status          string `json:"status"`
		Target          string `json:"php_target"`
		SourceHash      string `json:"source_sha256"`
		GeneratedHash   string `json:"generated_sha256"`
		CompilerVersion string `json:"compiler_version"`
	}
	// Additive/unknown report fields are intentionally tolerated. Report paths
	// never select filesystem locations; only caller-owned fixed paths are read.
	if err := json.Unmarshal(reportBytes, &report); err != nil {
		t.Fatal(err)
	}
	if report.Schema != "gohp.build/v0.1" || report.Status != "generated" || report.Target != pin.PHPTarget ||
		report.SourceHash != digest(bundle.Bytes()) || report.GeneratedHash != digest(generated) {
		t.Fatal("status=rejected: compiler report does not match exact input/output bytes")
	}
	t.Log("status=generated; report hashes verified; not yet a compatibility result")
	_, compileStderr := container(t, "build", append(buildArgs, "--entrypoint=go", pin.GoImage,
		"build", "-trimpath", "-o", "/candidate/program", "/candidate/main.go")...)
	if len(compileStderr) != 0 {
		t.Fatalf("unexpected build stderr: %s", compileStderr)
	}
	t.Log("status=compiled; same pinned GoHP runtime; not yet a compatibility result")
	actual, actualErr := container(t, "candidate_runtime", "--mount", "type=bind,source="+filepath.Join(candidate, "program")+",target=/program,readonly",
		"--entrypoint=/program", pin.GoImage)
	expected, expectedErr := container(t, "php_oracle", "--mount", "type=bind,source="+input+",target=/input,readonly", "--entrypoint=php", pin.PHPImage,
		"-d", "display_errors=stderr", "-d", "log_errors=0", "-d", "error_reporting=-1", "/input/unit.php")
	if len(actualErr) != 0 || len(expectedErr) != 0 || !bytes.Equal(actual, expected) || string(expected) != pin.Expected {
		t.Fatalf("status=differential_failed: candidate=%q stderr=%q PHP=%q stderr=%q fixture=%q", actual, actualErr, expected, expectedErr, pin.Expected)
	}
	evidence := map[string]any{
		"schema": "gossip.translation-observation/v1", "status": "differential_pass",
		"wordpress_commit": pin.WordPressCommit, "source_path": pin.SourcePath, "source_sha256": pin.SourceHash,
		"fragments": selections, "bundle_sha256": digest(bundle.Bytes()), "gohp_commit": pin.GoHPCommit,
		"gohp_archive_sha256": archiveHash, "go_image": pin.GoImage, "php_image": pin.PHPImage,
		"go_version": strings.TrimSpace(string(buildStdout)), "compiler_version": report.CompilerVersion,
		"compiler_sha256":  digest(read(t, filepath.Join(candidate, "gohp"))),
		"generated_sha256": digest(generated), "report_sha256": digest(reportBytes),
		"binary_sha256": digest(read(t, filepath.Join(candidate, "program"))), "php_target": pin.PHPTarget,
		"stdout": string(actual), "stderr": "", "exit_code": 0,
		"scope": "five extracted, unchanged constant-return helpers; stdout/empty stderr/exit only; not whole-file translation or return-type equivalence",
	}
	encoded, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))
}
