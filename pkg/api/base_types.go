// GENERATED CODE
// DO NOT EDIT

package game_protocol

import (
	"bytes"
	"encoding/binary"
)

const vector_cloneSize int = 8

type vector_clone struct {
	x vector
}

func Newvector_clone(raw [vector_cloneSize]byte) (res vector_clone, err error) {

	res.x, err = Newvector([vectorSize]byte(raw[0:16]))
	if err != nil {
		return res, err
	}

}

func Newvector_cloneBytes(item vector_clone) (res [vector_cloneSize]byte, err error) {
}

const vectorSize int = 16

type vector struct {
	x int32
	y int32
}

func Newvector(raw [vectorSize]byte) (res vector, err error) {

	res.x, err = Newint32([int32Size]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.y, err = Newint32([int32Size]byte(raw[4:8]))
	if err != nil {
		return res, err
	}

}

func NewvectorBytes(item vector) (res [vectorSize]byte, err error) {
}

const positionSize int = 8

type position struct {
	x int32
	y int32
}

func Newposition(raw [positionSize]byte) (res position, err error) {

	res.x, err = Newint32([int32Size]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.y, err = Newint32([int32Size]byte(raw[4:8]))
	if err != nil {
		return res, err
	}

}

func NewpositionBytes(item position) (res [positionSize]byte, err error) {
}

type entity_code byte

const (
	entity_codeplayer entity_code = 0
	entity_codeenemy  entity_code = 1
)
