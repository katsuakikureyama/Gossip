package phporacle

import "testing"

func TestOracleVersionAdmission(t *testing.T) {
	for _, version := range []string{"8.3.36", "8.3.37", "8.3.100"} {
		if err := ValidateVersion(version); err != nil {
			t.Errorf("admit %s: %v", version, err)
		}
	}
	for _, version := range []string{"", "7.4.33", "8.2.36", "8.3.9", "8.3.13", "8.3.30", "8.3.35", "8.4.0", "9.3.36", "8.3", "8.3.036", "8.3.+36", "8.3.-36", "8.3.36RC1", "8.3.36-dev", "8.3.36\n", " 8.3.36", "8.3.4294967296"} {
		if err := ValidateVersion(version); err == nil {
			t.Errorf("unexpected admission: %q", version)
		}
	}
}

func TestOracleImageAdmission(t *testing.T) {
	if err := ValidateImage(Image); err != nil {
		t.Fatal(err)
	}
	for _, image := range []string{"", "php:8.3.36-cli", "php:latest", "wordpress@sha256:04be81de853887fca707742785b32bf93fb4a692246dd1bc90b3790105071a9a", Image + "x"} {
		if err := ValidateImage(image); err == nil {
			t.Errorf("unexpected admission: %q", image)
		}
	}
}
