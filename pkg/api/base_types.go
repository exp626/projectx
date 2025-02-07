// GENERATED CODE
// DO NOT EDIT

package game_protocol

import protocol "github.com/exp626/projectx/pkg/protocol"

const SizeVector int = 8

type Vector struct {
	X int32
	Y int32
}

func NewVector(raw [SizeVector]byte) (res Vector, err error) {

	res.X, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.Y, err = protocol.Newint32([protocol.Sizeint32]byte(raw[4:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewVectorBytes(item Vector) (res [SizeVector]byte, err error) {

	XBytes, err := protocol.Newint32Bytes(item.X)
	if err != nil {
		return res, err
	}

	copy(res[0:4], XBytes[:])

	YBytes, err := protocol.Newint32Bytes(item.Y)
	if err != nil {
		return res, err
	}

	copy(res[4:8], YBytes[:])

	return res, nil
}

const SizePosition int = 8

type Position struct {
	X int32
	Y int32
}

func NewPosition(raw [SizePosition]byte) (res Position, err error) {

	res.X, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.Y, err = protocol.Newint32([protocol.Sizeint32]byte(raw[4:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewPositionBytes(item Position) (res [SizePosition]byte, err error) {

	XBytes, err := protocol.Newint32Bytes(item.X)
	if err != nil {
		return res, err
	}

	copy(res[0:4], XBytes[:])

	YBytes, err := protocol.Newint32Bytes(item.Y)
	if err != nil {
		return res, err
	}

	copy(res[4:8], YBytes[:])

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
