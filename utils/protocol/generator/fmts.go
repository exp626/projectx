package generator

import (
	"io"
	"text/template"
)

const (
	baseTypesFmt = `// GENERATED CODE
// DO NOT EDIT

package {{.PackageName}}
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
	structTypeFmt = `type {{.Name}} struct {
    {{ range $field := .Options.Fields }} {{ $field.Name }} {{ $field.Type }}
{{ end }}}`
	enumTypeFmt = `type {{ .Name }} {{ .Options.Type }}
	{{$typeName := .Name}}
	const (
		{{ range $value := .Options.Values }}{{$typeName}}{{$value.Name}} {{$typeName}} = {{ $value.Value }}
{{ end }})
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
