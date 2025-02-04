package generator

import (
	"encoding/json"
	"errors"
	"text/template"
)

type (
	ProtocolTypeMainFields struct {
		Type        TypeName        `json:"type"`
		Name        TypeName        `json:"name"`
		Description string          `json:"description"`
		RawOptions  json.RawMessage `json:"options"`
	}
	ProtocolType struct {
		ProtocolTypeMainFields
		Options interface{}

		IsSizeDynamic bool
		Size          uint64
	}
)

func (p *ProtocolType) UnmarshalJSON(bytes []byte) (err error) {
	err = json.Unmarshal(bytes, &p.ProtocolTypeMainFields)
	if err != nil {
		return err
	}

	switch p.Type {
	case structType:
		p.Options = &StructOptions{
			CalculatedSize: make(map[string]uint64),
		}

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
		Type   TypeName           `json:"type"`
		Values []EnumOptionsValue `json:"values"`
	}
	EnumOptionsValue struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	StructOptions struct {
		Fields []Field `json:"fields"`

		FieldsConstruct string

		// map field name: size
		CalculatedSize map[string]uint64
	}
)

func (p *ProtocolType) FormatType() (typeStr string, err error) {
	wr := &StringWriter{}

	tmpl := template.New(string(p.Name))

	typeFmt := ""

	switch p.Type {
	case structType:
		typeFmt = structTypeFmt

		opt, ok := p.Options.(*StructOptions)
		if !ok {
			return typeStr, errors.New("struct options is not set")
		}

		err = opt.EnrichConstructFormat()
		if err != nil {
			return typeStr, err
		}
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

func (p *ProtocolType) CalculateSize() (err error) {
	switch p.Type {
	case enumType:
		opts, ok := p.Options.(*EnumOptions)
		if !ok {
			return errors.New("enum options is not set")
		}

		p.Size = knownTypes[opts.Type].Size
	case structType:
		opts, ok := p.Options.(*StructOptions)
		if !ok {
			return errors.New("enum options is not set")
		}

		for _, field := range opts.Fields {
			fieldType, ok := knownTypes[field.Type]
			if !ok {
				return errors.New("unknown struct field type")
			}

			if fieldType.IsSizeDynamic {
				p.IsSizeDynamic = true
			} else if fieldType.Size != 0 {
				opts.CalculatedSize[field.Name] = fieldType.Size
				p.Size += fieldType.Size
			} else {
				// possibly infinite recursion
				// TODO: fix

				err = fieldType.CalculateSize()
				if err != nil {
					return err
				}

				opts.CalculatedSize[field.Name] = fieldType.Size
				p.Size += fieldType.Size
			}
		}
	default:
		return nil
	}

	return nil
}

func (o *StructOptions) EnrichConstructFormat() (err error) {
	var offset uint64

	for i := 0; i < len(o.Fields); i++ {
		o.Fields[i].Offset = offset
		o.Fields[i].EndOffset = offset + knownTypes[o.Fields[i].Type].Size

		offset = o.Fields[i].EndOffset
	}

	tmpl := template.New("fields")

	tmpl, err = tmpl.Parse(structFieldsConstructFmt)
	if err != nil {
		return err
	}

	strWriter := &StringWriter{}

	err = tmpl.Execute(strWriter, o.Fields)
	if err != nil {
		return err
	}

	o.FieldsConstruct = strWriter.s

	return nil
}
