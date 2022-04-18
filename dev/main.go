package main

import (
	"fmt"
	"github.com/julienetie/go-dung"
)

func main() {
	fmt.Println("Hello")

	options := godung.GoDungConfig{
		Url:       "https://go.dev",
		PageLimit: 100,
	}
	godung.GoDung(options)
}
