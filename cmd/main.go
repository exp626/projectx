package main

import (
	"log"
	"projectx/utils/protocol/generator"
)

func main() {
	g := generator.NewProtocolParser(generator.Config{
		Path:           "./protocol/protocol.json",
		OutputDir:      "./",
		OutputLanguage: "go",
	})

	err := g.Parse()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
