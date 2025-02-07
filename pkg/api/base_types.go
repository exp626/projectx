// GENERATED CODE
// DO NOT EDIT

package game_protocol

import protocol "github.com/exp626/projectx/pkg/protocol"

const SizeVectorClone int = 8

type VectorClone struct {
	x Vector
}

func NewVectorClone(raw [SizeVectorClone]byte) (res VectorClone, err error) {

	res.x, err = NewVector([SizeVector]byte(raw[0:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewVectorCloneBytes(item VectorClone) (res [SizeVectorClone]byte, err error) {

	xBytes, err := NewVectorBytes(item.x)
	if err != nil {
		return res, err
	}

	copy(res[0:8], xBytes[:])

	return res, nil
}

const SizeVector int = 8

type Vector struct {
	x int32
	y int32
}

func NewVector(raw [SizeVector]byte) (res Vector, err error) {

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

func NewVectorBytes(item Vector) (res [SizeVector]byte, err error) {

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

const SizePosition int = 8

type Position struct {
	x int32
	y int32
}

func NewPosition(raw [SizePosition]byte) (res Position, err error) {

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

func NewPositionBytes(item Position) (res [SizePosition]byte, err error) {

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

const SizeEntityCode = protocol.Sizebyte

type EntityCode byte

const (
	EntityCodePlayer EntityCode = 0
	EntityCodeEnemy  EntityCode = 1
)

func NewEntityCode(raw [SizeEntityCode]byte) (res EntityCode, err error) {
	baseRes, err := protocol.Newbyte(raw)
	if err != nil {
		return res, err
	}

	res = EntityCode(baseRes)

	return res, nil
}

func NewEntityCodeBytes(item EntityCode) (res [SizeEntityCode]byte, err error) {
	res, err = protocol.NewbyteBytes(byte(item))
	if err != nil {
		return res, err
	}

	return res, nil
}
