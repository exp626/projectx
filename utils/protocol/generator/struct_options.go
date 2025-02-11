package generator

import (
	"fmt"
)

type StructOptions struct {
	Fields []Field `json:"fields"`

	FieldsConstruct string
	BytesConstruct  string

	// map field name: size
	CalculatedSize map[string]uint64
}

func (s *StructOptions) Format() (fieldsFormat string, err error) {
	fieldsFormat += "\n"

	for _, field := range s.Fields {
		fieldsFormat += fmt.Sprintf(structFieldFmt+"\n", field.Name, field.Type)
	}

	fieldsFormat += "\n"

	return fieldsFormat, nil
}

func (s *StructOptions) FormatFieldsConstruct() (formatted string, err error) {
	formatted += "\n"

	for _, field := range s.Fields {
		formatted += fmt.Sprintf(
			structFieldConstruct,
			field.Name,
			ReturnPackage(field.Type),
			field.Type,
			ReturnPackage(field.Type),
			field.Type,
			field.Offset,
			field.EndOffset,
		)
	}

	formatted += "\n"

	return formatted, nil
}

func (s *StructOptions) FormatFieldsBytesConstruct() (formatted string, err error) {
	formatted += "\n"

	for _, field := range s.Fields {
		formatted += fmt.Sprintf(
			structFieldBytesConstruct,
			field.Name,
			ReturnPackage(field.Type),
			field.Type,
			field.Name,
			field.Offset,
			field.EndOffset,
			field.Name,
		)
	}

	formatted += "\n"

	return formatted, nil
}
