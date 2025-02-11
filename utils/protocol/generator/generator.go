package generator

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
)

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

	err = p.Manifest.RenameAsLanguage(p.cfg.OutputLanguage)
	if err != nil {
		return err
	}

	err = p.Manifest.FillKnownTypes()
	if err != nil {
		return err
	}

	{
		baseTypesFile, err := p.OpenGeneratedFile("base_types.go")
		if err != nil {
			return err
		}

		defer baseTypesFile.Close()

		fileData, err := p.Manifest.Types.Format()
		if err != nil {
			return err
		}

		fileData, err = p.FormatFile(fileData)
		if err != nil {
			return err
		}

		_, err = baseTypesFile.WriteAt([]byte(fileData), 0)
		if err != nil {
			return err
		}
	}

	{
		commandsFile, err := p.OpenGeneratedFile("commands.go")
		if err != nil {
			return err
		}

		defer commandsFile.Close()

		fileData, err := p.Manifest.Commands.Format()
		if err != nil {
			return err
		}

		fileData, err = p.FormatFile(fileData)
		if err != nil {
			return err
		}

		_, err = commandsFile.WriteAt([]byte(fileData), 0)
		if err != nil {
			return err
		}
	}

	{
		serverFile, err := p.OpenGeneratedFile("server.go")
		if err != nil {
			return err
		}

		defer serverFile.Close()

		fileData, err := p.Manifest.Commands.FormatServer()
		if err != nil {
			return err
		}

		fileData, err = p.FormatFile(fileData)
		if err != nil {
			return err
		}

		_, err = serverFile.WriteAt([]byte(fileData), 0)
		if err != nil {
			return err
		}
	}

	{
		clientFile, err := p.OpenGeneratedFile("client.go")
		if err != nil {
			return err
		}

		defer clientFile.Close()

		fileData, err := p.Manifest.Commands.FormatClient()
		if err != nil {
			return err
		}

		fileData, err = p.FormatFile(fileData)
		if err != nil {
			return err
		}

		_, err = clientFile.WriteAt([]byte(fileData), 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *ProtocolParser) OpenGeneratedFile(name string) (file *os.File, err error) {
	file, err = os.OpenFile(
		fmt.Sprintf("%s%s", p.cfg.OutputDir, name),
		os.O_CREATE|os.O_TRUNC|os.O_RDWR,
		0644,
	)
	if err != nil {
		return file, err
	}

	return file, nil
}

func (p *ProtocolParser) FormatCode(src string) (data string, err error) {
	switch p.cfg.OutputLanguage {
	case GoLanguage:
		rawData, err := format.Source([]byte(src))
		if err != nil {
			return data, err
		}

		data = string(rawData)
	default:
		data = src
	}

	return data, nil
}

func (p *ProtocolParser) FormatFile(src string) (formatted string, err error) {
	formatted = fmt.Sprintf(fileFmt, p.Manifest.PackageName, src)

	formatted, err = p.FormatCode(formatted)
	if err != nil {
		return formatted, err
	}

	return formatted, nil
}
