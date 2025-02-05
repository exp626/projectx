package generator

import (
	"io"
	"text/template"
)

const (
	baseTypesFmt = `// GENERATED CODE
// DO NOT EDIT

package {{.PackageName}}

import protocol "github.com/exp626/projectx/pkg/protocol"

  {{ range $type := .Types }}
	{{$type}}
  {{ end }}
`
	commandsFmt = `// GENERATED CODE
// DO NOT EDIT

package {{.PackageName}}

import protocol "github.com/exp626/projectx/pkg/protocol"

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
const Size{{.Name}} int = {{.Size}}
type {{.Name}} struct {
    {{ range $field := .Options.Fields }} {{ $field.Name }} {{ $field.Type }}
{{ end }}}

func New{{.Name}}(raw [Size{{.Name}}]byte) (res {{.Name}}, err error){
	{{.Options.FieldsConstruct}}

	return res, nil
}

func New{{.Name}}Bytes(item {{.Name}}) (res [Size{{.Name}}]byte, err error) {
	{{.Options.BytesConstruct}}

	return res, nil
}
`
	enumTypeFmt = `
	const Size{{.Name}} = protocol.Size{{.Options.Type}}
	type {{ .Name }} {{ .Options.Type }}
	{{$typeName := .Name}}
	const (
		{{ range $value := .Options.Values }}{{$typeName}}{{$value.Name}} {{$typeName}} = {{ $value.Value }}
	{{ end }})
	
	func New{{.Name}}(raw [Size{{.Name}}]byte) (res {{.Name}}, err error){
		baseRes, err := protocol.New{{.Options.Type}}(raw)
		if err != nil {
			return res, err
		}

		res = {{.Name}}(baseRes)
		
		return res, nil
	}
	
	func New{{.Name}}Bytes(item {{.Name}}) (res [Size{{.Name}}]byte, err error) {
		res, err = protocol.New{{.Options.Type}}Bytes({{.Options.Type}}(item))
		if err != nil {
			return res, err
		}

		return res, nil
	}
`
	structFieldsConstructFmt = `
    {{ range $field := . }} 
	res.{{$field.Name}}, err = {{if $field.IsBaseType}}protocol.{{end}}New{{$field.Type}}( [{{if $field.IsBaseType}}protocol.{{end}}Size{{$field.Type}}]byte (raw[{{$field.Offset}}:{{$field.EndOffset}}]))
	if err != nil {
		return res, err
	}
	{{end}}
`
	structBytesConstructFmt = `
	{{ range $field := . }}
	{{$field.Name}}Bytes, err := {{if $field.IsBaseType}}protocol.{{end}}New{{$field.Type}}Bytes(item.{{$field.Name}})
	if err != nil {
		return res, err
	}

	copy(res[{{$field.Offset}}:{{$field.EndOffset}}], {{$field.Name}}Bytes[:])
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
