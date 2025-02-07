// GENERATED CODE
// DO NOT EDIT

package game_protocol

import (
	"context"
	protocol "github.com/exp626/projectx/pkg/protocol"
)

type Service interface {
	// 21
	PlayerMove(ctx context.Context, body PlayerMoveBody) (err error)
	// 19
	CreatePlayer(ctx context.Context, body CreatePlayerBody) (err error)
	// 20
	Input(ctx context.Context, body InputBody) (err error)
}

const SizePlayerMoveBody int = 20

type PlayerMoveBody struct {
	entity_id int32
	position  Vector
	direction Vector
}

func NewPlayerMoveBody(raw [SizePlayerMoveBody]byte) (res PlayerMoveBody, err error) {

	res.entity_id, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.position, err = NewVector([SizeVector]byte(raw[4:12]))
	if err != nil {
		return res, err
	}

	res.direction, err = NewVector([SizeVector]byte(raw[12:20]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewPlayerMoveBodyBytes(item PlayerMoveBody) (res [SizePlayerMoveBody]byte, err error) {

	entity_idBytes, err := protocol.Newint32Bytes(item.entity_id)
	if err != nil {
		return res, err
	}

	copy(res[0:4], entity_idBytes[:])

	positionBytes, err := NewVectorBytes(item.position)
	if err != nil {
		return res, err
	}

	copy(res[4:12], positionBytes[:])

	directionBytes, err := NewVectorBytes(item.direction)
	if err != nil {
		return res, err
	}

	copy(res[12:20], directionBytes[:])

	return res, nil
}

const SizeCreatePlayerBody int = 13

type CreatePlayerBody struct {
	entity_type EntityCode
	entity_id   int32
	position    Vector
}

func NewCreatePlayerBody(raw [SizeCreatePlayerBody]byte) (res CreatePlayerBody, err error) {

	res.entity_type, err = NewEntityCode([SizeEntityCode]byte(raw[0:1]))
	if err != nil {
		return res, err
	}

	res.entity_id, err = protocol.Newint32([protocol.Sizeint32]byte(raw[1:5]))
	if err != nil {
		return res, err
	}

	res.position, err = NewVector([SizeVector]byte(raw[5:13]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewCreatePlayerBodyBytes(item CreatePlayerBody) (res [SizeCreatePlayerBody]byte, err error) {

	entity_typeBytes, err := NewEntityCodeBytes(item.entity_type)
	if err != nil {
		return res, err
	}

	copy(res[0:1], entity_typeBytes[:])

	entity_idBytes, err := protocol.Newint32Bytes(item.entity_id)
	if err != nil {
		return res, err
	}

	copy(res[1:5], entity_idBytes[:])

	positionBytes, err := NewVectorBytes(item.position)
	if err != nil {
		return res, err
	}

	copy(res[5:13], positionBytes[:])

	return res, nil
}

const SizeInputBody int = 8

type InputBody struct {
	direction Vector
}

func NewInputBody(raw [SizeInputBody]byte) (res InputBody, err error) {

	res.direction, err = NewVector([SizeVector]byte(raw[0:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewInputBodyBytes(item InputBody) (res [SizeInputBody]byte, err error) {

	directionBytes, err := NewVectorBytes(item.direction)
	if err != nil {
		return res, err
	}

	copy(res[0:8], directionBytes[:])

	return res, nil
}

const (

	// команда на движение игрока
	// 21
	CommandCodePlayerMove byte = 21

	// команда на создание игрока
	// 19
	CommandCodeCreatePlayer byte = 19

	// команда на ввод данных с джойстика
	// 20
	CommandCodeInput byte = 20
)
