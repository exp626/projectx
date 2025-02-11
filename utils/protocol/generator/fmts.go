package generator

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

import (
	"context"
	protocol "github.com/exp626/projectx/pkg/protocol"
)

type Service interface {
	{{ range $command := .Commands }}// {{.CommandCode}}
	{{ $command.Name }}(ctx context.Context, body {{ $command.Body.Name }}) (err error)
	{{ end }}
}

{{ range $type := .Types }}
{{$type}}
{{ end }}
const (
{{ range $command := .Commands }}
// {{$command.Body.Description}}
// {{.CommandCode}}
	CommandCode{{ $command.Name }} byte = {{ .CommandCode }}
{{ end }}
)
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
)

// new gen fmt
const (
	fileFmt = `// GENERATED CODE
// DO NOT EDIT

package %s

%s
`
	structFmt = `
	%s

	type %s struct {%s}
	
	%s
`
	structFieldFmt = "%s %s"
	enumFmt        = `
	%s

	type %s %s
	const (%s)

	%s
`
	enumValue = `%s%s %s = %d`
	baseFmt   = "%s %s"

	typeSizeFmt          = "const Size%s = %d"
	structFieldConstruct = `
	body.%s, err = %sNew%s( [%sSize%s]byte (raw[%d:%d]))
	if err != nil {
		return body, err
	}
`
	structFieldBytesConstruct = `
	%sBytes, err := %sNew%sBytes(body.%s)
	if err != nil {
		return raw, err
	}

	copy(raw[%d:%d], %sBytes[:])
`
	structConstructorsFmt = `
	func New%s(raw [Size%s]byte) (body %s, err error) {
	%s
	return body, nil
}

	func New%sBytes(body %s) (raw [Size%s]byte, err error) {
	%s
	return raw, nil
}
`
	enumAndBaseConstructorsFmt = `
	func New%s(raw [Size%s]byte) (res %s, err error){
		baseRes, err := protocol.New%s(raw)
		if err != nil {
			return res, err
		}

		res = %s(baseRes)
		
		return res, nil
	}
	
	func New%sBytes(item %s) (res [Size%s]byte, err error) {
		res, err = protocol.New%sBytes(%s(item))
		if err != nil {
			return res, err
		}

		return res, nil
	}
`

	baseTypesFileFmt = `

import protocol "github.com/exp626/projectx/pkg/protocol"

%s
`
)
