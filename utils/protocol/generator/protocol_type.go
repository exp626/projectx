package generator

import "encoding/json"

const (
	structType = "struct"
	enumType   = "enum"

	uint8Type   = "uint8"
	uint16Type  = "uint16"
	uint32Type  = "uint32"
	uint64Type  = "uint64"
	int8Type    = "int8"
	int16Type   = "int16"
	int32Type   = "int32"
	int64Type   = "int64"
	float32Type = "float32"
	float64Type = "float64"
	stringType  = "string"
	intType     = "int"
	uintType    = "uint"
	uintptrType = "uintptr"
	byteType    = "byte"
	runeType    = "rune"
)

type (
	ProtocolTypeMainFields struct {
		Type        string          `json:"type"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		RawOptions  json.RawMessage `json:"options"`
	}
	ProtocolType struct {
		ProtocolTypeMainFields
		Options interface{}
	}
)

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

type (
	EnumOptions struct {
		Type   string             `json:"type"`
		Values []EnumOptionsValue `json:"values"`
	}
	EnumOptionsValue struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	StructOptions struct {
		Fields []ProtocolType `json:"fields"`
	}
)
