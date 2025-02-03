package main

import (
	"flag"
	"fmt"
	"generator/generator"
	"os"
)

// фаил протокола получаем из параметров запуска

var fileName string
var templatePath string
var outputPath string

func main() {
	flag.StringVar(&fileName, "f", "protocol.json", "file name")
	flag.StringVar(&templatePath, "t", "templates", "template path")
	flag.StringVar(&outputPath, "o", "generated", "output path")

	flag.Parse()

	if fileName == "" {
		fmt.Println("file name is required")
		os.Exit(1)
	}

	if outputPath == "" {
		fmt.Println("output path is required")
		os.Exit(1)
	}

	if templatePath == "" {
		fmt.Println("template path is required")
		os.Exit(1)
	}

	generator.GenerateProtocol(fileName, templatePath, outputPath)
}
