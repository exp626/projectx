package protocol

import (
	"bytes"
	"encoding/binary"
)

func New8ByteObject[T int | uint | int64 | uint64 | float64](raw [8]byte) (res T, err error) {
	err = binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &res)

	return res, err
}

func New8ByteObjectBytes[T int | uint | int64 | uint64 | float64](item T) (res [8]byte, err error) {
	buf := bytes.NewBuffer(res[:])

	err = binary.Write(buf, binary.LittleEndian, item)

	return [8]byte(buf.Bytes()), err
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
