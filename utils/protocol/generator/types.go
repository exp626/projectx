package generator

type TypeName string

const (
	structType TypeName = "struct"
	enumType   TypeName = "enum"

	uint8Type   TypeName = "uint8"
	uint16Type  TypeName = "uint16"
	uint32Type  TypeName = "uint32"
	uint64Type  TypeName = "uint64"
	int8Type    TypeName = "int8"
	int16Type   TypeName = "int16"
	int32Type   TypeName = "int32"
	int64Type   TypeName = "int64"
	float32Type TypeName = "float32"
	float64Type TypeName = "float64"
	stringType  TypeName = "string"
	intType     TypeName = "int"
	uintType    TypeName = "uint"
	byteType    TypeName = "byte"
	runeType    TypeName = "rune"
)

var (
	knownTypes = map[TypeName]*ProtocolType{
		uint8Type: {
			Size: 1,
		},
		uint16Type: {
			Size: 2,
		},
		uint32Type: {
			Size: 4,
		},
		uint64Type: {
			Size: 8,
		},
		int8Type: {
			Size: 1,
		},
		int16Type: {
			Size: 2,
		},
		int32Type: {
			Size: 4,
		},
		int64Type: {
			Size: 8,
		},
		float32Type: {
			Size: 4,
		},
		float64Type: {
			Size: 8,
		},
		intType: {
			Size: 8,
		},
		uintType: {
			Size: 8,
		},
		byteType: {
			Size: 1,
		},
		runeType: {
			Size: 1,
		},
		stringType: {
			IsSizeDynamic: true,
		},
	}
)
