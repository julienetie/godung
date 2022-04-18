package godung

import (
	"errors"
	"fmt"
	"strings"
	//	"github.com/PuerkitoBio/goquery"
)

const pageLimitDefault = 1000

type GoDungConfig struct {
	Url                   string
	PageLimit             int64
	ExcludeSubDirectories bool
}

func GetOptions(options *GoDungConfig) (config *GoDungConfig, err error) {
	fmt.Println(options)
	if *options == (GoDungConfig{}) {
		err = errors.New("The config requires the `Url` field")
		return options, err
	}

	// Set defaults
	if options.PageLimit == int64(0) {
		options.PageLimit = pageLimitDefault
	}

	// ExcludeSubDirectories is false by default

	config = options
	hasHttp := strings.HasPrefix(options.Url, "http")
	hasHttps := strings.HasPrefix(options.Url, "https")

	if hasHttp || hasHttps {
		err = nil
	} else {
		err = errors.New("The `Url` should begin with `http://` or `http2://`")
	}
	return
}

func GoDung(options *GoDungConfig) {
	fmt.Println("Go dung")
	GetOptions(options)
}
