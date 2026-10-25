package gohp_test

import (
	"testing"

	"github.com/katsuakikureyama/Gossip/gossip/internal/phporacle"
)

func requirePHPOracle(t *testing.T, image string) string {
	t.Helper()
	if err := phporacle.ValidateImage(image); err != nil {
		t.Fatal(err)
	}
	args := append([]string{"--entrypoint=php", image}, phporacle.VersionArgs()...)
	version, stderr := container(t, "php_version", args...)
	if len(stderr) != 0 {
		t.Fatalf("unexpected PHP version stderr: %s", stderr)
	}
	if err := phporacle.ValidateVersion(string(version)); err != nil {
		t.Fatal(err)
	}
	return string(version)
}

func runPHPOracle(t *testing.T, stage, image, input, script string) ([]byte, []byte) {
	t.Helper()
	if err := phporacle.ValidateImage(image); err != nil {
		t.Fatal(err)
	}
	args := []string{"--mount", "type=bind,source=" + input + ",target=/input,readonly", "--env=LC_ALL=C", "--entrypoint=php", image}
	return container(t, stage, append(args, phporacle.ScriptArgs(script)...)...)
}
