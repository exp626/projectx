// GENERATED CODE
// DO NOT EDIT

package game_protocol

import (
	"context"
	protocol "github.com/exp626/projectx/pkg/protocol"
)

const (
	CommandCodePlayerMove   byte = 21
	CommandCodeCreatePlayer byte = 19
	CommandCodeInput        byte = 20
)

type Service interface {
	PlayerMove(ctx context.Context, body PlayerMoveBody) (err error)
	CreatePlayer(ctx context.Context, body CreatePlayerBody) (err error)
	Input(ctx context.Context, body InputBody) (err error)
}

const SizePlayerMoveBody = 20

type PlayerMoveBody struct {
	EntityId  int32
	Position  Vector
	Direction Vector
}

func NewPlayerMoveBody(raw [SizePlayerMoveBody]byte) (body PlayerMoveBody, err error) {

	body.EntityId, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return body, err
	}

	body.Position, err = NewVector([SizeVector]byte(raw[4:12]))
	if err != nil {
		return body, err
	}

	body.Direction, err = NewVector([SizeVector]byte(raw[12:20]))
	if err != nil {
		return body, err
	}

	return body, nil
}

func NewPlayerMoveBodyBytes(body PlayerMoveBody) (raw [SizePlayerMoveBody]byte, err error) {

	EntityIdBytes, err := protocol.Newint32Bytes(body.EntityId)
	if err != nil {
		return raw, err
	}

	copy(raw[0:4], EntityIdBytes[:])

	PositionBytes, err := NewVectorBytes(body.Position)
	if err != nil {
		return raw, err
	}

	copy(raw[4:12], PositionBytes[:])

	DirectionBytes, err := NewVectorBytes(body.Direction)
	if err != nil {
		return raw, err
	}

	copy(raw[12:20], DirectionBytes[:])

	return raw, nil
}

const SizeCreatePlayerBody = 13

type CreatePlayerBody struct {
	EntityType EntityCode
	EntityId   int32
	Position   Vector
}

func NewCreatePlayerBody(raw [SizeCreatePlayerBody]byte) (body CreatePlayerBody, err error) {

	body.EntityType, err = NewEntityCode([SizeEntityCode]byte(raw[0:1]))
	if err != nil {
		return body, err
	}

	body.EntityId, err = protocol.Newint32([protocol.Sizeint32]byte(raw[1:5]))
	if err != nil {
		return body, err
	}

	body.Position, err = NewVector([SizeVector]byte(raw[5:13]))
	if err != nil {
		return body, err
	}

	return body, nil
}

func NewCreatePlayerBodyBytes(body CreatePlayerBody) (raw [SizeCreatePlayerBody]byte, err error) {

	EntityTypeBytes, err := NewEntityCodeBytes(body.EntityType)
	if err != nil {
		return raw, err
	}

	copy(raw[0:1], EntityTypeBytes[:])

	EntityIdBytes, err := protocol.Newint32Bytes(body.EntityId)
	if err != nil {
		return raw, err
	}

	copy(raw[1:5], EntityIdBytes[:])

	PositionBytes, err := NewVectorBytes(body.Position)
	if err != nil {
		return raw, err
	}

	copy(raw[5:13], PositionBytes[:])

	return raw, nil
}

const SizeInputBody = 8

type InputBody struct {
	Direction Vector
}

func NewInputBody(raw [SizeInputBody]byte) (body InputBody, err error) {

	body.Direction, err = NewVector([SizeVector]byte(raw[0:8]))
	if err != nil {
		return body, err
	}

	return body, nil
}

func NewInputBodyBytes(body InputBody) (raw [SizeInputBody]byte, err error) {

	DirectionBytes, err := NewVectorBytes(body.Direction)
	if err != nil {
		return raw, err
	}

	copy(raw[0:8], DirectionBytes[:])

	return raw, nil
}
