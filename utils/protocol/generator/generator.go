package generator

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Path      string
	OutputDir string
}

type ProtocolParser struct {
	Manifest ProtocolManifest
	cfg      Config
}

func (p *ProtocolParser) Parse() (err error) {
	file, err := os.Open(p.cfg.Path)
	if err != nil {
		return err
	}

	err = json.NewDecoder(file).Decode(&p.Manifest)
	if err != nil {
		return err
	}

	{

		types := make([]string, len(p.Manifest.Types))

		for _, item := range p.Manifest.Types {
		}

		baseTypesFile, err := os.OpenFile(
			fmt.Sprintf("%s/base_types.go", p.cfg.OutputDir),
			os.O_CREATE|os.O_RDWR,
			os.ModeAppend,
		)
		if err != nil {
			return err
		}

		formatBaseTypes(baseTypesFile, p.Manifest.PackageName, types)
	}

	return nil
}
