package godung

import (
	"fmt"
	//	"github.com/PuerkitoBio/goquery"
)

type GoDungConfig struct {
	Url       string
	PageLimit int64
}

func GetOptions(options GoDungConfig) {
	fmt.Println(options)
}

func GoDung(options GoDungConfig) {
	fmt.Println("Go dung")
	GetOptions(options)
}
