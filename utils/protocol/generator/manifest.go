package generator

import (
	"errors"
	"log"
)

type Command struct {
	CommandCode byte         `json:"command_code"`
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

	log.Println(knownTypes)

	return nil
}
