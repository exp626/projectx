package protocol

import "errors"

var (
	ErrBodyIsTooShort              = errors.New("body is too short")
	ErrAllInformationWasNotWritten = errors.New("all information was not written")
)
