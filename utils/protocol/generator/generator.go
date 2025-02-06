package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	format "go/format"
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

	err = p.Manifest.FillKnownTypes()
	if err != nil {
		return err
	}

	{
		baseTypesFile, err := os.OpenFile(
			fmt.Sprintf("%sbase_types.go", p.cfg.OutputDir),
			os.O_CREATE|os.O_TRUNC|os.O_RDWR,
			0644,
		)
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

		switch p.cfg.OutputLanguage {
		case GoLanguage:
			fileData, err = format.Source(buf.Bytes())
			if err != nil {
				return err
			}
		default:
			fileData = buf.Bytes()
		}

		_, err = baseTypesFile.WriteAt(fileData, 0)
		if err != nil {
			return err
		}
	}

	{
		commandsFile, err := os.OpenFile(
			fmt.Sprintf("%scommands.go", p.cfg.OutputDir),
			os.O_CREATE|os.O_TRUNC|os.O_RDWR,
			0644,
		)
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

		switch p.cfg.OutputLanguage {
		case GoLanguage:
			fileData, err = format.Source(buf.Bytes())
			if err != nil {
				return err
			}
		default:
			fileData = buf.Bytes()
		}

		_, err = commandsFile.WriteAt(fileData, 0)
		if err != nil {
			return err
		}
	}

	{
		serverFile, err := os.OpenFile(
			fmt.Sprintf("%sserver.go", p.cfg.OutputDir),
			os.O_CREATE|os.O_TRUNC|os.O_RDWR,
			0644,
		)
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

		switch p.cfg.OutputLanguage {
		case GoLanguage:
			fileData, err = format.Source(buf.Bytes())
			if err != nil {
				return err
			}
		default:
			fileData = buf.Bytes()
		}

		_, err = serverFile.WriteAt(fileData, 0)
		if err != nil {
			return err
		}
	}

	{
		clientFile, err := os.OpenFile(
			fmt.Sprintf("%sclient.go", p.cfg.OutputDir),
			os.O_CREATE|os.O_TRUNC|os.O_RDWR,
			0644,
		)
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

		switch p.cfg.OutputLanguage {
		//case GoLanguage:
		//	fileData, err = format.Source(buf.Bytes())
		//	if err != nil {
		//		return err
		//	}
		default:
			fileData = buf.Bytes()
		}

		_, err = clientFile.WriteAt(fileData, 0)
		if err != nil {
			return err
		}
	}

	return nil
}
