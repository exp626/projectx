package main

import (
	"github.com/exp626/projectx/utils/protocol/generator"
	"log"
)

func main() {
	g := generator.NewProtocolParser(generator.Config{
		Path:           "./test_protocol.json",
		OutputDir:      "./pkg/api/",
		OutputLanguage: "go",
	})

	err := g.Parse()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
