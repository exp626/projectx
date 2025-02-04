package generator

import (
	"encoding/json"
	"errors"
	"text/template"
)

const (
	structType = "struct"
	enumType   = "enum"

	uint8Type      = "uint8"
	uint16Type     = "uint16"
	uint32Type     = "uint32"
	uint64Type     = "uint64"
	int8Type       = "int8"
	int16Type      = "int16"
	int32Type      = "int32"
	int64Type      = "int64"
	float32Type    = "float32"
	float64Type    = "float64"
	complex64Type  = "complex64"
	complex128Type = "complex128"
	stringType     = "string"
	intType        = "int"
	uintType       = "uint"
	uintptrType    = "uintptr"
	byteType       = "byte"
	runeType       = "rune"
)

type ProtocolTypeMainFields struct {
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	RawOptions  json.RawMessage `json:"options"`
}

type StructOptions struct {
	Fields []ProtocolType `json:"fields"`
}

type ProtocolType struct {
	ProtocolTypeMainFields
	Options interface{}
}

func (p *ProtocolType) UnmarshalJSON(bytes []byte) (err error) {
	err = json.Unmarshal(bytes, &p.ProtocolTypeMainFields)
	if err != nil {
		return err
	}

	switch p.Type {
	case structType:
		p.Options = &StructOptions{}

		err = json.Unmarshal(p.RawOptions, p.Options)
		if err != nil {
			return err
		}
	case enumType:
		p.Options = &EnumOptions{}

		err = json.Unmarshal(p.RawOptions, p.Options)
		if err != nil {
			return err
		}
	}

	return nil
}

type Command struct {
	CommandCode byte         `json:"command_code"`
	Body        ProtocolType `json:"body"`
}

type ProtocolManifest struct {
	PackageName string         `json:"packageName"`
	Commands    []Command      `json:"commands"`
	Types       []ProtocolType `json:"types"`
}

type EnumOptions struct {
	Type   string             `json:"type"`
	Values []EnumOptionsValue `json:"values"`
}

type EnumOptionsValue struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
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
		complex64Type,
		complex128Type,
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
