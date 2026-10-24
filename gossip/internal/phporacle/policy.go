// Package phporacle gates opt-in local developer oracles, not production PHP.
package phporacle

import (
	"fmt"
	"regexp"
	"strconv"
)

const Image = "php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd"
const Minimum = "8.3.36"

// New digests require review before Docker starts. Version admission alone
// is not a compatibility or vulnerability guarantee.
func ValidateImage(image string) error {
	if image != Image {
		return fmt.Errorf("unapproved PHP oracle image: use the documented PHP %s digest", Minimum)
	}
	return nil
}

var stable83 = regexp.MustCompile(`^8\.3\.(0|[1-9][0-9]*)$`)

func ValidateVersion(version string) error {
	match := stable83.FindStringSubmatch(version)
	if len(match) == 2 {
		patch, err := strconv.ParseUint(match[1], 10, 32)
		if err == nil && patch >= 36 {
			return nil
		}
	}
	return fmt.Errorf("PHP oracle %q is not admitted: require stable 8.3 >=%s; other series require explicit validation", version, Minimum)
}

// Check the actual binary in the same PHP process BEFORE requiring any fixture.
// Paths are separate argv values, never interpolated into PHP or a shell.
const guard = `if (!preg_match('/^8[.]3[.](0|[1-9][0-9]*)$/D', PHP_VERSION) || PHP_VERSION_ID < 80336 || PHP_MAJOR_VERSION !== 8 || PHP_MINOR_VERSION !== 3) { fwrite(STDERR, 'PHP oracle requires stable 8.3 >=8.3.36'); exit(78); } `

func phpArgs(code string) []string {
	return []string{"-d", "display_errors=stderr", "-d", "log_errors=0", "-d", "error_reporting=-1", "-r", guard + code}
}

func VersionArgs() []string { return phpArgs("echo PHP_VERSION;") }

func ScriptArgs(path string) []string {
	return append(phpArgs("require $argv[1];"), "--", path)
}
