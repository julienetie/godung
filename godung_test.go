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

func TestGetOptions(t *testing.T) {

	t.Run("The error message should indicate that the Url is required", func(t *testing.T) {
		_, err := GetOptions(&GoDungConfig{})

		if !ErrorContains(err, "The config requires the `Url` field") {
			t.Errorf("Missing Url error %v", err)
		}
	})

	t.Run("The URL should start with http or https", func(t *testing.T) {
		_, err := GetOptions(&GoDungConfig{
			Url: "exmaple.com",
		})

		if !ErrorContains(err, "The `Url` should begin with `http://` or `http2://`") {
			t.Errorf("Missing invalid Url error")
		}
	})

	t.Run("The URL should return defaults for missing values", func(t *testing.T) {

		config, err := GetOptions(&GoDungConfig{
			Url: "https://exmaple.com",
		})

		if err != nil {
			t.Errorf("Url with https should not throw an error")
		}

		_, err2 := GetOptions(&GoDungConfig{
			Url: "http://exmaple.com",
		})

		if err2 != nil {
			t.Errorf("Url with http should not throw an error")
		}

		if config.PageLimit != int64(1000) {
			t.Errorf("PageLimit is not default")
		}

		if config.ExcludeSubDirectories != false {
			t.Errorf("ExcludeSubDirectories is not default")
		}
	})
}
