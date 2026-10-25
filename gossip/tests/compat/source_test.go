package compat_test

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

type sourcePin struct {
	Schema     string            `json:"schema"`
	Repository string            `json:"repository"`
	Commit     string            `json:"commit"`
	WordPress  string            `json:"wordpress"`
	Files      map[string]string `json:"files"`
}

// readPinnedSource verifies the files actually checked out in this fork, not
// WordPress files bundled in a PHP runner image. Return the verified bytes so
// the integration test stages exactly what was hashed, even if the checkout
// changes while Docker is starting.
func readPinnedSource(t *testing.T) (sourcePin, map[string][]byte) {
	t.Helper()
	data, err := os.ReadFile("../../wordpress-source.json")
	if err != nil {
		t.Fatal(err)
	}
	var pin sourcePin
	if err := json.Unmarshal(data, &pin); err != nil {
		t.Fatal(err)
	}
	paths := []string{"license.txt", "wp-includes/version.php", "wp-includes/plugin.php", "wp-includes/class-wp-hook.php", "wp-includes/class-wp-filter-sentinel.php"}
	if pin.Schema != "gossip.wordpress-source/v1" || pin.Repository != "https://github.com/WordPress/WordPress" || !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(pin.Commit) || pin.WordPress == "" || len(pin.Files) != len(paths) {
		t.Fatal("invalid source pin")
	}
	files := make(map[string][]byte, len(paths))
	for _, path := range paths {
		data, err := os.ReadFile(filepath.Join("../../..", path))
		if err != nil {
			t.Fatal(err)
		}
		if hash := fmt.Sprintf("%x", sha256.Sum256(data)); hash != pin.Files[path] {
			t.Fatalf("WordPress source drift: %s got=%s expected=%s; review upstream changes before updating the pin", path, hash, pin.Files[path])
		}
		files[path] = data
	}
	return pin, files
}

func TestWordPressSourceBaseline(t *testing.T) {
	pin, _ := readPinnedSource(t)
	t.Logf("upstream_commit=%s wordpress=%s verified_files=%d", pin.Commit, pin.WordPress, len(pin.Files))
}

func stagePinnedSource(t *testing.T, files map[string][]byte) string {
	t.Helper()
	dir := t.TempDir()
	// The container runs as a different unprivileged UID and only receives
	// these public upstream sources, never the checkout's .git or config files.
	if err := os.Chmod(dir, 0755); err != nil {
		t.Fatal(err)
	}
	for path, data := range files {
		target := filepath.Join(dir, path)
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(target, data, 0644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}
