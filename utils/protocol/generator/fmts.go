package generator

import (
	"io"
	"text/template"
)

const (
	baseTypesFmt = `// GENERATED CODE
// DO NOT EDIT

package {{.PackageName}}
import (
	"bytes"
	"encoding/binary"
)

  {{ range $type := .Types }}
	{{$type}}
  {{ end }}
`
	commandsFmt = `// GENERATED CODE
// DO NOT EDIT

package {{.PackageName}}
{{ range $type := .Types }}
{{$type}}
{{ end }}
{{ range $command := .Commands }}
// {{$command.Body.Description}}
// {{.CommandCode}}
const (
	CommandCode{{ $command.Body.Name }} = {{ .CommandCode }}
)
{{ end }}
`

	baseTypeFmt = `// {{.Description}}
type {{.Name}} {{.Type}}`
	structTypeFmt = `
const {{.Name}}Size int = {{.Size}}
type {{.Name}} struct {
    {{ range $field := .Options.Fields }} {{ $field.Name }} {{ $field.Type }}
{{ end }}}

func New{{.Name}}(raw [{{.Name}}Size]byte) (res {{.Name}}, err error){
	{{.Options.FieldsConstruct}}
}

func New{{.Name}}Bytes(item {{.Name}}) (res [{{.Name}}Size]byte, err error) {
}
`
	enumTypeFmt = `type {{ .Name }} {{ .Options.Type }}
	{{$typeName := .Name}}
	const (
		{{ range $value := .Options.Values }}{{$typeName}}{{$value.Name}} {{$typeName}} = {{ $value.Value }}
{{ end }})
`
	structFieldsConstructFmt = `
    {{ range $field := . }} 
	res.{{$field.Name}}, err = New{{$field.Type}}( [{{$field.Type}}Size]byte (raw[{{$field.Offset}}:{{$field.EndOffset}}]))
	if err != nil {
		return res, err
	}
	{{end}}
`
)

func (m *ProtocolManifest) FormatBaseTypes(wr io.Writer) (err error) {
	types := make([]string, 0, len(m.Types))

	for _, item := range m.Types {
		typeFmt, err := item.FormatType()
		if err != nil {
			return err
		}

		types = append(types, typeFmt)
	}

	tmpl := template.New("baseTypes")

	tmpl, err = tmpl.Parse(baseTypesFmt)
	if err != nil {
		return err
	}

	err = tmpl.Execute(wr, struct {
		PackageName string
		Types       []string
	}{
		PackageName: m.PackageName,
		Types:       types,
	})
	if err != nil {
		return err
	}

	return err
}

func (m *ProtocolManifest) FormatCommands(wr io.Writer) (err error) {
	types := make([]string, 0, len(m.Commands))

	for _, item := range m.Commands {
		typeFmt, err := item.Body.FormatType()
		if err != nil {
			return err
		}

		types = append(types, typeFmt)
	}

	tmpl := template.New("commandsTypes")

	tmpl, err = tmpl.Parse(commandsFmt)
	if err != nil {
		return err
	}

	err = tmpl.Execute(wr, struct {
		PackageName string
		Types       []string
		Commands    []Command
	}{
		PackageName: m.PackageName,
		Types:       types,
		Commands:    m.Commands,
	})
	if err != nil {
		return err
	}

	return err
}
