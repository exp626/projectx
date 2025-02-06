package generator

import (
	"errors"
	"io"
	"text/template"
)

type CommandDirection string

const (
	CommandDirectionServerToClient = "server_to_client"
	CommandDirectionClientToServer = "client_to_server"
)

type Command struct {
	CommandCode byte         `json:"command_code"`
	Direction   string       `json:"direction"`
	Body        ProtocolType `json:"body"`
}

type ProtocolManifest struct {
	PackageName string         `json:"packageName"`
	Commands    []Command      `json:"commands"`
	Types       []ProtocolType `json:"types"`
}

func (m *ProtocolManifest) FillKnownTypes() (err error) {
	for i := 0; i < len(m.Types); i++ {
		knownTypes[m.Types[i].Name] = &m.Types[i]
	}

	for i := 0; i < len(m.Types); i++ {
		knownType, ok := knownTypes[m.Types[i].Name]
		if !ok {
			return errors.New("unknown type")
		}

		err = knownType.CalculateSize()
		if err != nil {
			return err
		}
	}

	for i := 0; i < len(m.Commands); i++ {
		err = m.Commands[i].Body.CalculateSize()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ProtocolManifest) FormatBaseTypes(wr io.Writer) (err error) {
	types := make([]string, 0, len(m.Types))

	for _, item := range m.Types {
		typeFmt, err := item.FormatType()
		if err != nil {
			return err
		}

		types = append(types, typeFmt)
	}

	tmpl := template.New("baseTypes")

	tmpl, err = tmpl.Parse(baseTypesFmt)
	if err != nil {
		return err
	}

	err = tmpl.Execute(wr, struct {
		PackageName string
		Types       []string
	}{
		PackageName: m.PackageName,
		Types:       types,
	})
	if err != nil {
		return err
	}

	return err
}

func (m *ProtocolManifest) FormatCommands(wr io.Writer) (err error) {
	types := make([]string, 0, len(m.Commands))

	for _, item := range m.Commands {
		typeFmt, err := item.Body.FormatType()
		if err != nil {
			return err
		}

		types = append(types, typeFmt)
	}

	tmpl := template.New("commandsTypes")

	tmpl, err = tmpl.Parse(commandsFmt)
	if err != nil {
		return err
	}

	err = tmpl.Execute(wr, struct {
		PackageName string
		Types       []string
		Commands    []Command
	}{
		PackageName: m.PackageName,
		Types:       types,
		Commands:    m.Commands,
	})
	if err != nil {
		return err
	}

	return err
}

func (m *ProtocolManifest) FormatServer(wr io.Writer) (err error) {
	tmpl := template.New("server")

	tmpl, err = tmpl.Parse(serverFmt)
	if err != nil {
		return err
	}

	err = tmpl.Execute(wr, struct {
		PackageName string
		Commands    []Command
	}{
		PackageName: m.PackageName,
		Commands:    m.Commands,
	})
	if err != nil {
		return err
	}

	return err
}
