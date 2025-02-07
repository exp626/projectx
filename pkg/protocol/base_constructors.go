package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	size8Byte = 8
	size4Byte = 4
	size2Byte = 2
	size1Byte = 1

	Sizeint     = size8Byte
	Sizeuint    = size8Byte
	Sizefloat64 = size8Byte
	Sizeint32   = size4Byte
	Sizebyte    = size1Byte
)

func New8ByteObject[T int | uint | int64 | uint64 | float64](raw [size8Byte]byte) (res T, err error) {
	err = binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &res)

	return res, err
}

func New8ByteObjectBytes[T int | uint | int64 | uint64 | float64](item T) (res [size8Byte]byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, item)

	return [8]byte(buf.Bytes()), err
}

func New4ByteObject[T int32 | uint32](raw [4]byte) (res T, err error) {
	err = binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &res)

	return res, err
}

func New4ByteObjectBytes[T int32 | uint32](item T) (res [4]byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, item)

	return [4]byte(buf.Bytes()), err
}

func Newint(raw [8]byte) (res int, err error) {
	return New8ByteObject[int](raw)
}

func NewintBytes(item int) (res [8]byte, err error) {
	return New8ByteObjectBytes(item)
}

func Newuint(raw [8]byte) (res uint, err error) {
	return New8ByteObject[uint](raw)
}

func NewuintBytes(item uint) (res [8]byte, err error) {
	return New8ByteObjectBytes(item)
}

func Newfloat64(raw [8]byte) (res float64, err error) {
	return New8ByteObject[float64](raw)
}

func Newfloat64Bytes(item float64) (res [8]byte, err error) {
	return New8ByteObjectBytes(item)
}

func Newbyte(raw [1]byte) (res byte, err error) {
	return raw[0], nil
}

func NewbyteBytes(item byte) (raw [1]byte, err error) {
	return [1]byte{item}, nil
}

func Newint32(raw [4]byte) (res int32, err error) {
	return New4ByteObject[int32](raw)
}

func Newint32Bytes(item int32) (res [4]byte, err error) {
	return New4ByteObjectBytes(item)
}
