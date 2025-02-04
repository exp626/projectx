package generator

import (
	"errors"
	"text/template"
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

func (p *ProtocolType) FormatType() (typeStr string, err error) {
	wr := &StringWriter{}

	tmpl := template.New(p.Name)

	typeFmt := ""

	switch p.Type {
	case structType:
		typeFmt = structTypeFmt
	case enumType:
		typeFmt = enumTypeFmt
	case uint8Type,
		uint16Type,
		uint32Type,
		uint64Type,
		int8Type,
		int16Type,
		int32Type,
		int64Type,
		float32Type,
		float64Type,
		stringType,
		intType,
		uintType,
		uintptrType,
		byteType,
		runeType:
		typeFmt = baseTypeFmt
	default:
		return "", errors.New("unknown type")
	}

	tmpl, err = tmpl.Parse(typeFmt)
	if err != nil {
		return wr.s, err
	}

	err = tmpl.Execute(wr, *p)
	if err != nil {
		return wr.s, err
	}

	return wr.s, err
}
