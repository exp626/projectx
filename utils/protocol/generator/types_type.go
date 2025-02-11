package generator

import "fmt"

type Types []ProtocolType

func (t Types) FormatTypes() (formatted string, err error) {
	formatted += "\n"

	for _, item := range t {
		typeDeclaration, err := item.Format()
		if err != nil {
			return formatted, err
		}

		formatted += typeDeclaration
	}

	formatted += "\n"

	return formatted, nil
}

func (t Types) Format() (formatted string, err error) {
	typesDeclaration, err := t.FormatTypes()
	if err != nil {
		return formatted, err
	}

	formatted = fmt.Sprintf(baseTypesFileFmt, typesDeclaration)

	return formatted, nil
}
