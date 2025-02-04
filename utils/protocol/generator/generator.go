package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Config struct {
	Path      string
	OutputDir string
	OutputLanguage
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
	defer file.Close()

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
			os.O_CREATE|os.O_TRUNC|os.O_RDWR,
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

	{

		//types := make([]string, 0, len(p.Manifest.Types))
		//
		//for _, item := range p.Manifest.Types {
		//	typeFmt, err := item.FormatType()
		//	if err != nil {
		//		return err
		//	}
		//
		//	types = append(types, typeFmt)
		//}
		//
		//baseTypesFile, err := os.OpenFile(
		//	fmt.Sprintf("%scommands.go", p.cfg.OutputDir),
		//	os.O_CREATE|os.O_TRUNC|os.O_RDWR,
		//	0644,
		//)
		//if err != nil {
		//	return err
		//}
		//
		//defer baseTypesFile.Close()
		//
		//err = formatBaseTypes(baseTypesFile, p.Manifest.PackageName, types)
		//if err != nil {
		//	return err
		//}
	}

	switch p.cfg.OutputLanguage {
	case GoLanguage:
		cmd := exec.Command(fmt.Sprintf("go fmt %s*.go", p.cfg.OutputDir))
		if err = cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
