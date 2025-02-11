package generator

import (
	"fmt"
)

type (
	EnumOptions struct {
		Type   TypeName           `json:"type"`
		Values []EnumOptionsValue `json:"values"`
	}
	EnumOptionsValue struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
)

func (e *EnumOptions) Format(typeName TypeName) (formatted string, err error) {
	formatted += "\n"

	for _, value := range e.Values {
		formatted += fmt.Sprintf(enumValue+"\n", typeName, value.Name, typeName, value.Value)
	}

	formatted += "\n"

	return formatted, nil
}
