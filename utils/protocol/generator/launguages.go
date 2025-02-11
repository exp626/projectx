package generator

import (
	"encoding/json"
	"errors"
	"slices"
)

type OutputLanguage string

const (
	GoLanguage OutputLanguage = "go"
)

var (
	supportedLanguages = []OutputLanguage{
		GoLanguage,
	}
)

func (o *OutputLanguage) UnmarshalJSON(bytes []byte) (err error) {
	var str string

	err = json.Unmarshal(bytes, &str)
	if err != nil {
		return err
	}

	if !slices.Contains(supportedLanguages, OutputLanguage(str)) {
		return errors.New("unsupported language")
	}

	return nil
}
