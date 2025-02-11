// GENERATED CODE
// DO NOT EDIT

package game_protocol

import protocol "github.com/exp626/projectx/pkg/protocol"

const SizeVector = 8

type Vector struct {
	X int32
	Y int32
}

func NewVector(raw [SizeVector]byte) (body Vector, err error) {

	body.X, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return body, err
	}

	body.Y, err = protocol.Newint32([protocol.Sizeint32]byte(raw[4:8]))
	if err != nil {
		return body, err
	}

	return body, nil
}

func NewVectorBytes(body Vector) (raw [SizeVector]byte, err error) {

	XBytes, err := protocol.Newint32Bytes(body.X)
	if err != nil {
		return raw, err
	}

	copy(raw[0:4], XBytes[:])

	YBytes, err := protocol.Newint32Bytes(body.Y)
	if err != nil {
		return raw, err
	}

	copy(raw[4:8], YBytes[:])

	return raw, nil
}

const SizePosition = 8

type Position struct {
	X int32
	Y int32
}

func NewPosition(raw [SizePosition]byte) (body Position, err error) {

	body.X, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return body, err
	}

	body.Y, err = protocol.Newint32([protocol.Sizeint32]byte(raw[4:8]))
	if err != nil {
		return body, err
	}

	return body, nil
}

func NewPositionBytes(body Position) (raw [SizePosition]byte, err error) {

	XBytes, err := protocol.Newint32Bytes(body.X)
	if err != nil {
		return raw, err
	}

	copy(raw[0:4], XBytes[:])

	YBytes, err := protocol.Newint32Bytes(body.Y)
	if err != nil {
		return raw, err
	}

	copy(raw[4:8], YBytes[:])

	return raw, nil
}

const SizeEntityCode = 1

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
