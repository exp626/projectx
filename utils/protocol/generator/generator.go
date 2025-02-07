package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
)

type Config struct {
	Path            string
	OutputDir       string
	OutputLanguage  OutputLanguage
	ProtocolPackage string
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

	p.Manifest.ConvertNames(p.cfg.OutputLanguage)

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

		fileData := make([]byte, 0)
		buf := bytes.NewBuffer(fileData)

		err = p.Manifest.FormatBaseTypes(buf)
		if err != nil {
			return err
		}

		fileData, err = p.FormatCode(buf)
		if err != nil {
			return err
		}

		_, err = baseTypesFile.WriteAt(fileData, 0)
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

		fileData := make([]byte, 0)
		buf := bytes.NewBuffer(fileData)

		err = p.Manifest.FormatCommands(buf)
		if err != nil {
			return err
		}

		fileData, err = p.FormatCode(buf)
		if err != nil {
			return err
		}

		_, err = commandsFile.WriteAt(fileData, 0)
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

		fileData := make([]byte, 0)
		buf := bytes.NewBuffer(fileData)

		err = p.Manifest.FormatServer(buf)
		if err != nil {
			return err
		}

		fileData, err = p.FormatCode(buf)
		if err != nil {
			return err
		}

		_, err = serverFile.WriteAt(fileData, 0)
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

		fileData := make([]byte, 0)
		buf := bytes.NewBuffer(fileData)

		err = p.Manifest.FormatClient(buf)
		if err != nil {
			return err
		}

		fileData, err = p.FormatCode(buf)
		if err != nil {
			return err
		}

		_, err = clientFile.WriteAt(fileData, 0)
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

func (p *ProtocolParser) FormatCode(buf *bytes.Buffer) (data []byte, err error) {
	switch p.cfg.OutputLanguage {
	case GoLanguage:
		data, err = format.Source(buf.Bytes())
		if err != nil {
			return data, err
		}
	default:
		data = buf.Bytes()
	}

	return data, nil
}
