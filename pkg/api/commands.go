// GENERATED CODE
// DO NOT EDIT

package game_protocol

const player_moveSize int = 0

type player_move struct {
	entity_id int32
	position  vector
	direction vector
}

func Newplayer_move(raw [player_moveSize]byte) (res player_move, err error) {

	res.entity_id, err = Newint32([int32Size]byte(raw[0:4]))
	if err != nil {
		return res, err
	}

	res.position, err = Newvector([vectorSize]byte(raw[4:20]))
	if err != nil {
		return res, err
	}

	res.direction, err = Newvector([vectorSize]byte(raw[20:36]))
	if err != nil {
		return res, err
	}

}

func Newplayer_moveBytes(item player_move) (res [player_moveSize]byte, err error) {
}

const create_playerSize int = 0

type create_player struct {
	entity_type entity_code
	entity_id   int32
	position    vector
}

func Newcreate_player(raw [create_playerSize]byte) (res create_player, err error) {

	res.entity_type, err = Newentity_code([entity_codeSize]byte(raw[0:1]))
	if err != nil {
		return res, err
	}

	res.entity_id, err = Newint32([int32Size]byte(raw[1:5]))
	if err != nil {
		return res, err
	}

	res.position, err = Newvector([vectorSize]byte(raw[5:21]))
	if err != nil {
		return res, err
	}

}

func Newcreate_playerBytes(item create_player) (res [create_playerSize]byte, err error) {
}

const inputSize int = 0

type input struct {
	direction vector
}

func Newinput(raw [inputSize]byte) (res input, err error) {

	res.direction, err = Newvector([vectorSize]byte(raw[0:16]))
	if err != nil {
		return res, err
	}

}

func NewinputBytes(item input) (res [inputSize]byte, err error) {
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
