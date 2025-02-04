// GENERATED CODE
// DO NOT EDIT

package game_protocol

type player_move struct {
	entity_id int32
	position  vector
	direction vector
}

type create_player struct {
	entity_type entity_code
	entity_id   int32
	position    vector
}

type input struct {
	direction vector
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
