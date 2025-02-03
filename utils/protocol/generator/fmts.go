package generator

import (
	"io"
	"text/template"
)

const (
	baseTypesFmt = `package {{.PackageName}}
  {{ range $type := .Types }}
    type {{$type}}

  {{ end }}
`
	baseTypeFmt   = `{{.Name}} {{.Type}}`
	structTypeFmt = `{{.Name}} struct {
    {{ range $field := .Options.Fields }}
      {{ $field.Name }} {{ $field.Type }}
    {{ end }}
  }`
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
