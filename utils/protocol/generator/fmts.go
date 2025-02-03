package generator

import (
	"io"
	"text/template"
)

const (
	baseTypesFmt = `package {{.package_name}}
  {{ range $type := .types }}
    type {{$type}}

  {{ end }}
`
	baseTypeFmt   = `{{.name}} {{.type}}`
	structTypeFmt = `{{.name}} struct {
    {{ range $field := .options.fields }}
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
