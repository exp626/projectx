package main

import (
	"github.com/exp626/projectx/utils/protocol/generator"
	"log"
)

func main() {
	cfg := generator.Config{}
	err := cfg.UnmarshalFlags()
	if err != nil {
		log.Fatalf("Parsing args error: %v", err)
	}

	g := generator.NewProtocolParser(cfg)

	err = g.Parse()
	if err != nil {
		log.Fatalf("Generating error: %v", err)
	}
}
