package main

import (
	"flag"
	"fmt"
	"generator/generator"
	"os"
)

// фаил протокола получаем из параметров запуска

var fileName string

func main() {
	flag.StringVar(&fileName, "f", "protocol.json", "file name")

	if fileName == "" {
		fmt.Println("file name is required")
		os.Exit(1)
	}

	generator.GenerateProtocol(fileName)
}
