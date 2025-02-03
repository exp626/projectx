package generator

import "text/template"

const (
	structType = "struct"

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

type ProtocolType struct {
	Type        string
	Name        string
	Description string
	Options     interface{}
}

type Command struct {
	CommandCode byte
	Body        ProtocolType
}

type ProtocolManifest struct {
	PackageName string
	Commands    []Command
	Types       []ProtocolType
}

func (p *ProtocolType) FormatType() (typeStr string, err error) {
	wr := &StringWriter{}

	tmpl := template.New(p.Name)

	typeFmt := ""

	switch p.Type {
	case structType:
		typeFmt = structTypeFmt
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
	}

	tmpl, err = tmpl.Parse(typeFmt)
	if err != nil {
		return wr.s, err
	}

	err = tmpl.Execute(wr)
	if err != nil {
		return wr.s, err
	}

	return wr.s, err
}
