package generator

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
