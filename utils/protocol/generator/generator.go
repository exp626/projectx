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

func NewProtocolParser(cfg Config) (p *ProtocolParser) {
	return &ProtocolParser{
		cfg: cfg,
	}
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
		types := make([]string, 0, len(p.Manifest.Types))

		for _, item := range p.Manifest.Types {
			typeFmt, err := item.FormatType()
			if err != nil {
				return err
			}

			types = append(types, typeFmt)
		}

		baseTypesFile, err := os.OpenFile(
			fmt.Sprintf("%sbase_types.go", p.cfg.OutputDir),
			os.O_CREATE|os.O_RDWR,
			0644,
		)
		if err != nil {
			return err
		}

		defer baseTypesFile.Close()

		err = formatBaseTypes(baseTypesFile, p.Manifest.PackageName, types)
		if err != nil {
			return err
		}
	}

	return nil
}
