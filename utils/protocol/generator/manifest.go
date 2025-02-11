package generator

import (
	"errors"
)

type CommandDirection string

const (
	CommandDirectionServerToClient = "server_to_client"
	CommandDirectionClientToServer = "client_to_server"
)

type ProtocolManifest struct {
	PackageName string   `json:"packageName"`
	Commands    Commands `json:"commands"`
	Types       Types    `json:"types"`
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

func (m *ProtocolManifest) RenameAsLanguage(lang OutputLanguage) (err error) {
	for i := 0; i < len(m.Types); i++ {
		err = m.Types[i].RenameAsLanguage(lang)
		if err != nil {
			return err
		}
	}

	for i := 0; i < len(m.Commands); i++ {
		err = m.Commands[i].RenameAsLanguage(lang)
		if err != nil {
			return err
		}
	}

	return nil
}
