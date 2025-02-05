// GENERATED CODE
// DO NOT EDIT

package game_protocol

import protocol "github.com/exp626/projectx/pkg/protocol"

const Sizevector_clone int = 8

type vector_clone struct {
	x vector
}

func Newvector_clone(raw [Sizevector_clone]byte) (res vector_clone, err error) {

	res.x, err = Newvector([Sizevector]byte(raw[0:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func Newvector_cloneBytes(item vector_clone) (res [Sizevector_clone]byte, err error) {

	xBytes, err := NewvectorBytes(item.x)
	if err != nil {
		return res, err
	}

	copy(res[0:8], xBytes[:])

	return res, nil
}

const Sizevector int = 8

type vector struct {
	x int32
	y int32
}

func Newvector(raw [Sizevector]byte) (res vector, err error) {

	res.x, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.y, err = protocol.Newint32([protocol.Sizeint32]byte(raw[4:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewvectorBytes(item vector) (res [Sizevector]byte, err error) {

	xBytes, err := protocol.Newint32Bytes(item.x)
	if err != nil {
		return res, err
	}

	copy(res[0:4], xBytes[:])

	yBytes, err := protocol.Newint32Bytes(item.y)
	if err != nil {
		return res, err
	}

	copy(res[4:8], yBytes[:])

	return res, nil
}

const Sizeposition int = 8

type position struct {
	x int32
	y int32
}

func Newposition(raw [Sizeposition]byte) (res position, err error) {

	res.x, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.y, err = protocol.Newint32([protocol.Sizeint32]byte(raw[4:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewpositionBytes(item position) (res [Sizeposition]byte, err error) {

	xBytes, err := protocol.Newint32Bytes(item.x)
	if err != nil {
		return res, err
	}

	copy(res[0:4], xBytes[:])

	yBytes, err := protocol.Newint32Bytes(item.y)
	if err != nil {
		return res, err
	}

	copy(res[4:8], yBytes[:])

	return res, nil
}

const Sizeentity_code = protocol.Sizebyte

type entity_code byte

const (
	entity_codeplayer entity_code = 0
	entity_codeenemy  entity_code = 1
)

func Newentity_code(raw [Sizeentity_code]byte) (res entity_code, err error) {
	baseRes, err := protocol.Newbyte(raw)
	if err != nil {
		return res, err
	}

	res = entity_code(baseRes)

	return res, nil
}

func Newentity_codeBytes(item entity_code) (res [Sizeentity_code]byte, err error) {
	res, err = protocol.NewbyteBytes(byte(item))
	if err != nil {
		return res, err
	}

	return res, nil
}
