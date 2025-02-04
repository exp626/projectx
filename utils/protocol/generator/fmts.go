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
    type {{$type}}
  {{ end }}
`
	baseTypeFmt   = `{{.Name}} {{.Type}}`
	structTypeFmt = `{{.Name}} struct {
    {{ range $field := .Options.Fields }} {{ $field.Name }} {{ $field.Type }}
{{ end }}}`
	enumTypeFmt = `{{ .Name }} {{ .Options.Type }}
	{{$typeName := .Name}}
	const (
		{{ range $value := .Options.Values }}{{$typeName}}{{$value.Name}} {{$typeName}} = {{ $value.Value }}
{{ end }})
`
)

func formatBaseTypes(wr io.Writer, packageName string, types []string) (err error) {
	tmpl := template.New("baseTypes")

	tmpl, err = tmpl.Parse(baseTypesFmt)
	if err != nil {
		return err
	}

	err = tmpl.Execute(wr, struct {
		PackageName string
		Types       []string
	}{
		PackageName: packageName,
		Types:       types,
	})
	if err != nil {
		return err
	}

	return err
}
