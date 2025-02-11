package generator

import (
	"errors"
	"flag"
	"os"
	"slices"
)

type Config struct {
	Path           string
	OutputDir      string
	OutputLanguage OutputLanguage
}

func (c *Config) UnmarshalFlags() (err error) {
	if len(os.Args) < 2 {
		return errors.New("invalid protocol manifest path")
	}

	lang := ""

	flag.Parse()

	flag.StringVar(&c.OutputDir, "out", "./", "path to output directory")
	flag.StringVar(&lang, "lang", "go", "output language (only go is available)")

	c.OutputLanguage = OutputLanguage(lang)

	if !slices.Contains(supportedLanguages, c.OutputLanguage) {
		return errors.New("unsupported language")
	}

	c.Path = os.Args[len(os.Args)-1]

	return nil
}
