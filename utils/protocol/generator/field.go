package generator

type Field struct {
	Name      string
	Type      TypeName
	Offset    uint64
	EndOffset uint64

	IsBaseType bool
}
