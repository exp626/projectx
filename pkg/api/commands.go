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
	EntityId  int32
	Position  Vector
	Direction Vector
}

func NewPlayerMoveBody(raw [SizePlayerMoveBody]byte) (res PlayerMoveBody, err error) {

	res.EntityId, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.Position, err = NewVector([SizeVector]byte(raw[4:12]))
	if err != nil {
		return res, err
	}

	res.Direction, err = NewVector([SizeVector]byte(raw[12:20]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewPlayerMoveBodyBytes(item PlayerMoveBody) (res [SizePlayerMoveBody]byte, err error) {

	EntityIdBytes, err := protocol.Newint32Bytes(item.EntityId)
	if err != nil {
		return res, err
	}

	copy(res[0:4], EntityIdBytes[:])

	PositionBytes, err := NewVectorBytes(item.Position)
	if err != nil {
		return res, err
	}

	copy(res[4:12], PositionBytes[:])

	DirectionBytes, err := NewVectorBytes(item.Direction)
	if err != nil {
		return res, err
	}

	copy(res[12:20], DirectionBytes[:])

	return res, nil
}

const SizeCreatePlayerBody int = 13

type CreatePlayerBody struct {
	EntityType EntityCode
	EntityId   int32
	Position   Vector
}

func NewCreatePlayerBody(raw [SizeCreatePlayerBody]byte) (res CreatePlayerBody, err error) {

	res.EntityType, err = NewEntityCode([SizeEntityCode]byte(raw[0:1]))
	if err != nil {
		return res, err
	}

	res.EntityId, err = protocol.Newint32([protocol.Sizeint32]byte(raw[1:5]))
	if err != nil {
		return res, err
	}

	res.Position, err = NewVector([SizeVector]byte(raw[5:13]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewCreatePlayerBodyBytes(item CreatePlayerBody) (res [SizeCreatePlayerBody]byte, err error) {

	EntityTypeBytes, err := NewEntityCodeBytes(item.EntityType)
	if err != nil {
		return res, err
	}

	copy(res[0:1], EntityTypeBytes[:])

	EntityIdBytes, err := protocol.Newint32Bytes(item.EntityId)
	if err != nil {
		return res, err
	}

	copy(res[1:5], EntityIdBytes[:])

	PositionBytes, err := NewVectorBytes(item.Position)
	if err != nil {
		return res, err
	}

	copy(res[5:13], PositionBytes[:])

	return res, nil
}

const SizeInputBody int = 8

type InputBody struct {
	Direction Vector
}

func NewInputBody(raw [SizeInputBody]byte) (res InputBody, err error) {

	res.Direction, err = NewVector([SizeVector]byte(raw[0:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewInputBodyBytes(item InputBody) (res [SizeInputBody]byte, err error) {

	DirectionBytes, err := NewVectorBytes(item.Direction)
	if err != nil {
		return res, err
	}

	copy(res[0:8], DirectionBytes[:])

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
