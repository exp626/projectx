package generator

type StringWriter struct {
	s string
}

func (s *StringWriter) Write(p []byte) (n int, err error) {
	s.s += string(p)

	return len(p), nil
}
