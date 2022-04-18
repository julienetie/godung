package godung

import (
	"strings"
	"testing"
)

func ErrorContains(got error, want string) bool {
	if got == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(got.Error(), want)
}

func TestGoDung(t *testing.T) {
	t.Skip()
}

// Should ensure URL exists
// Should apply defaults for fields that don't exist
func TestGetOptions(t *testing.T) {

	t.Run("The error message should indicate that the Url is required", func(t *testing.T) {
		_, err := GetOptions(GoDungConfig{})

		if ErrorContains(err, "The config requires the `Url` field") {
			t.Errorf("Missing Url error: %v", err)
		}
	})

	t.Run("The URL should start with http or https", func(t *testing.T) {
		_, err := GetOptions(GoDungConfig{
			Url: "exmaple.com",
		})

		if ErrorContains(err, "The `Url` should begin with `http://` or `http2://`") {
			t.Errorf("Missing invalid Url error: %v", err)
		}
	})

	t.Run("The URL should return defaults for missing values", func(t *testing.T) {
		config, err := GetOptions(GoDungConfig{
			Url: "https://exmaple.com",
		})

		_, err2 := GetOptions(GoDungConfig{
			Url: "http://exmaple.com",
		})

		if err != nil {
			t.Errorf("Url with https should not throw an error")
		}

		if err2 != nil {
			t.Errorf("Url with http should not throw an error")
		}

		if config.PageLimit != "1000" {
			t.Errorf("PageLimit is not default")
		}

		if config.IncludeSubDirectories != true {
			t.Errorf("IncludeSubDirectories is not default")
		}
	})
}
