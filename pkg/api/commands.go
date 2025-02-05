// GENERATED CODE
// DO NOT EDIT

package game_protocol

import protocol "github.com/exp626/projectx/pkg/protocol"

const Sizeplayer_move int = 20

type player_move struct {
	entity_id int32
	position  vector
	direction vector
}

func Newplayer_move(raw [Sizeplayer_move]byte) (res player_move, err error) {

	res.entity_id, err = protocol.Newint32([protocol.Sizeint32]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.position, err = Newvector([Sizevector]byte(raw[4:12]))
	if err != nil {
		return res, err
	}

	res.direction, err = Newvector([Sizevector]byte(raw[12:20]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func Newplayer_moveBytes(item player_move) (res [Sizeplayer_move]byte, err error) {

	entity_idBytes, err := protocol.Newint32Bytes(item.entity_id)
	if err != nil {
		return res, err
	}

	copy(res[0:4], entity_idBytes[:])

	positionBytes, err := NewvectorBytes(item.position)
	if err != nil {
		return res, err
	}

	copy(res[4:12], positionBytes[:])

	directionBytes, err := NewvectorBytes(item.direction)
	if err != nil {
		return res, err
	}

	copy(res[12:20], directionBytes[:])

	return res, nil
}

const Sizecreate_player int = 13

type create_player struct {
	entity_type entity_code
	entity_id   int32
	position    vector
}

func Newcreate_player(raw [Sizecreate_player]byte) (res create_player, err error) {

	res.entity_type, err = Newentity_code([Sizeentity_code]byte(raw[0:1]))
	if err != nil {
		return res, err
	}

	res.entity_id, err = protocol.Newint32([protocol.Sizeint32]byte(raw[1:5]))
	if err != nil {
		return res, err
	}

	res.position, err = Newvector([Sizevector]byte(raw[5:13]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func Newcreate_playerBytes(item create_player) (res [Sizecreate_player]byte, err error) {

	entity_typeBytes, err := Newentity_codeBytes(item.entity_type)
	if err != nil {
		return res, err
	}

	copy(res[0:1], entity_typeBytes[:])

	entity_idBytes, err := protocol.Newint32Bytes(item.entity_id)
	if err != nil {
		return res, err
	}

	copy(res[1:5], entity_idBytes[:])

	positionBytes, err := NewvectorBytes(item.position)
	if err != nil {
		return res, err
	}

	copy(res[5:13], positionBytes[:])

	return res, nil
}

const Sizeinput int = 8

type input struct {
	direction vector
}

func Newinput(raw [Sizeinput]byte) (res input, err error) {

	res.direction, err = Newvector([Sizevector]byte(raw[0:8]))
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewinputBytes(item input) (res [Sizeinput]byte, err error) {

	directionBytes, err := NewvectorBytes(item.direction)
	if err != nil {
		return res, err
	}

	copy(res[0:8], directionBytes[:])

	return res, nil
}

// команда на движение игрока
// 21
const (
	CommandCodeplayer_move = 21
)

// команда на создание игрока
// 20
const (
	CommandCodecreate_player = 20
)

// команда на ввод данных с джойстика
// 20
const (
	CommandCodeinput = 20
)
