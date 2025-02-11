package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"slices"
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
		opt := &StructOptions{
			CalculatedSize: make(map[string]uint64),
		}

		err = json.Unmarshal(p.RawOptions, opt)
		if err != nil {
			return err
		}

		p.Options = opt
	case enumType:
		p.Options = &EnumOptions{}

		err = json.Unmarshal(p.RawOptions, p.Options)
		if err != nil {
			return err
		}
	}

	return nil
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
				_, ok = opts.CalculatedSize[field.Name]
				if !ok {
					opts.CalculatedSize[field.Name] = fieldType.Size
					p.Size += fieldType.Size
				}
			} else {
				// possibly infinite recursion
				// TODO: fix

				err = fieldType.CalculateSize()
				if err != nil {
					return err
				}

				_, ok = opts.CalculatedSize[field.Name]
				if !ok {
					opts.CalculatedSize[field.Name] = fieldType.Size
					p.Size += fieldType.Size
				}
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

		o.Fields[i].IsBaseType = slices.Contains(baseTypes, o.Fields[i].Type)
	}

	return nil
}

func (p *ProtocolType) Format() (formatted string, err error) {
	switch p.Type {
	case structType:
		opt, ok := p.Options.(*StructOptions)
		if !ok {
			return formatted, errors.New("struct options is not set")
		}

		err = opt.EnrichConstructFormat()
		if err != nil {
			return formatted, err
		}

		structFields, err := opt.Format()
		if err != nil {
			return structFields, err
		}

		structConstructors, err := p.FormatConstructors()
		if err != nil {
			return formatted, err
		}

		sizeDeclaration, err := p.FormatSize()
		if err != nil {
			return formatted, err
		}

		formatted = fmt.Sprintf(structFmt, sizeDeclaration, p.Name, structFields, structConstructors)
	case enumType:
		opt, ok := p.Options.(*EnumOptions)
		if !ok {
			return formatted, errors.New("enum options is not set")
		}

		enumValues, err := opt.Format(p.Name)
		if err != nil {
			return formatted, err
		}

		constructorsDeclaration, err := p.FormatConstructors()
		if err != nil {
			return formatted, err
		}

		sizeDeclaration, err := p.FormatSize()
		if err != nil {
			return formatted, err
		}

		formatted = fmt.Sprintf(enumFmt, sizeDeclaration, p.Name, opt.Type, enumValues, constructorsDeclaration)
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
		formatted = fmt.Sprintf(baseFmt, p.Name, p.Type)
	default:
		return formatted, errors.New("unknown type")
	}

	formatted += "\n"

	return formatted, nil
}

func (p *ProtocolType) FormatConstructors() (formatted string, err error) {
	switch p.Type {
	case structType:
		opt, ok := p.Options.(*StructOptions)
		if !ok {
			return formatted, errors.New("struct options is not set")
		}

		fieldsConstructFormatted, err := opt.FormatFieldsConstruct()
		if err != nil {
			return formatted, err
		}

		fieldsBytesFormatted, err := opt.FormatFieldsBytesConstruct()
		if err != nil {
			return formatted, err
		}

		err = opt.EnrichConstructFormat()
		if err != nil {
			return formatted, err
		}

		formatted = fmt.Sprintf(
			structConstructorsFmt,
			p.Name,
			p.Name,
			p.Name,
			fieldsConstructFormatted,
			p.Name,
			p.Name,
			p.Name,
			fieldsBytesFormatted,
		)
	case enumType:
		opt, ok := p.Options.(*EnumOptions)
		if !ok {
			return formatted, errors.New("enum options is not set")
		}

		formatted = fmt.Sprintf(
			enumAndBaseConstructorsFmt,
			p.Name,
			p.Name,
			p.Name,
			opt.Type,
			p.Name,

			p.Name,
			p.Name,
			p.Name,
			opt.Type,
			opt.Type,
		)
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

		formatted = fmt.Sprintf(
			enumAndBaseConstructorsFmt,
			p.Name,
			p.Name,
			p.Name,
			p.Type,
			p.Name,

			p.Name,
			p.Name,
			p.Name,
			p.Type,
			p.Type,
		)
	default:
		return formatted, errors.New("unknown type")
	}

	return formatted, nil
}

func (p *ProtocolType) FormatSize() (formatted string, err error) {
	formatted = fmt.Sprintf(typeSizeFmt+"\n", p.Name, p.Size)

	return formatted, nil
}

func (p *ProtocolType) RenameAsLanguage(lang OutputLanguage) (err error) {
	switch lang {
	case GoLanguage:
		p.Name.ToCamel()
		p.Type.ToCamel()
	}

	switch p.Type {
	case structType:
		opts, ok := p.Options.(*StructOptions)
		if !ok {
			return errors.New("invalid struct options")
		}

		for i := 0; i < len(opts.Fields); i++ {
			switch lang {
			case GoLanguage:
				opts.Fields[i].Type.ToCamel()
				opts.Fields[i].Name = strcase.ToCamel(opts.Fields[i].Name)
			}
		}
	case enumType:
		opts, ok := p.Options.(*EnumOptions)
		if !ok {
			return errors.New("invalid struct options")
		}

		for i := 0; i < len(opts.Values); i++ {
			switch lang {
			case GoLanguage:
				opts.Values[i].Name = strcase.ToCamel(opts.Values[i].Name)
			}
		}
	}
	return nil
}
