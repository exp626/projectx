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

import protocol "github.com/exp626/projectx/pkg/protocol"

{{ range $type := .Types }}
{{$type}}
{{ end }}
const (
{{ range $command := .Commands }}
// {{$command.Body.Description}}
// {{.CommandCode}}
	CommandCode{{ $command.Body.Name }} byte = {{ .CommandCode }}
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
	serverFmt = `// GENERATED CODE
// DO NOT EDIT

package {{.PackageName}}

import (
	"context"
	"errors"
)

type Service interface {
	{{ range $command := .Commands }}// {{.CommandCode}}
	{{ $command.Body.Name }}(ctx context.Context, body {{ $command.Body.Name }}) (err error)
	{{ end }}
}

type Server struct{	
	service Service
}

func (s *Server) HandleCommand(rawBody []byte) (err error){
	if len(rawBody) < 2 {
		return errors.New("body is too short")
	}

	commandCode := rawBody[0]

	rawCommandBody := rawBody[1:]

	switch commandCode{
	{{ range $command := .Commands }}
	case CommandCode{{ $command.Body.Name }}:
		if len(rawCommandBody) < Size{{$command.Body.Name}} {
			return errors.New("body is too short")
		}

		body, err := New{{$command.Body.Name}}([Size{{$command.Body.Name}}]byte(rawCommandBody))
		if err != nil {
			return err
		}

		err = s.service.{{ $command.Body.Name }}(context.Background(), body)
		if err != nil {
			return err
		}
	{{ end }}
	default:
		return errors.New("unknown command code")
	}

	return nil
}
`
)
