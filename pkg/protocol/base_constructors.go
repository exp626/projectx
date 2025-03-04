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
	Sizeint8    = size1Byte
	Sizeuint8   = size1Byte
	Sizeuint32  = size4Byte
	Sizeint16   = size2Byte
)

func New8ByteObject[T int | uint | int64 | uint64 | float64](raw [size8Byte]byte) (res T, err error) {
	err = binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &res)

	return res, err
}

func New8ByteObjectBytes[T int | uint | int64 | uint64 | float64](item T) (res [size8Byte]byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, item)

	return [size8Byte]byte(buf.Bytes()), err
}

func New4ByteObject[T int32 | uint32 | float32](raw [size4Byte]byte) (res T, err error) {
	err = binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &res)

	return res, err
}

func New4ByteObjectBytes[T int32 | uint32 | float32](item T) (res [size4Byte]byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, item)

	return [size4Byte]byte(buf.Bytes()), err
}

func New2ByteObject[T int16 | uint16](raw [size2Byte]byte) (res T, err error) {
	err = binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &res)

	return res, err
}

func New2ByteObjectBytes[T int16 | uint16](item T) (res [size2Byte]byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, item)

	return [size2Byte]byte(buf.Bytes()), err
}

func New1ByteObject[T int8 | uint8](raw [size1Byte]byte) (res T, err error) {
	err = binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &res)

	return res, err
}

func New1ByteObjectBytes[T int8 | uint8](item T) (res [size1Byte]byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, item)

	return [size1Byte]byte(buf.Bytes()), err
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

func Newfloat32(raw [4]byte) (res float32, err error) {
	return New4ByteObject[float32](raw)
}

func Newfloat32Bytes(item float32) (res [4]byte, err error) {
	return New4ByteObjectBytes(item)
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

func Newint8(raw [1]byte) (res int8, err error) {
	return New1ByteObject[int8](raw)
}

func Newint8Bytes(item int8) (res [1]byte, err error) {
	return New1ByteObjectBytes(item)
}

func Newuint8(raw [1]byte) (res uint8, err error) {
	return New1ByteObject[uint8](raw)
}

func Newuint8Bytes(item uint8) (res [1]byte, err error) {
	return New1ByteObjectBytes(item)
}

func Newuint32(raw [Sizeuint32]byte) (res uint32, err error) {
	return New4ByteObject[uint32](raw)
}

func Newuint32Bytes(item uint32) (res [Sizeuint32]byte, err error) {
	return New4ByteObjectBytes(item)
}

func Newint16(raw [Sizeint16]byte) (res int16, err error) {
	return New2ByteObject[int16](raw)
}

func Newint16Bytes(item int16) (res [Sizeint16]byte, err error) {
	return New2ByteObjectBytes(item)
}
