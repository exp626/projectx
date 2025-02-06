// GENERATED CODE
// DO NOT EDIT

package game_protocol

import (
	"context"
	"errors"
)

type Service interface {
	// 21
	player_move(ctx context.Context, body player_move) (err error)
	// 19
	create_player(ctx context.Context, body create_player) (err error)
	// 20
	input(ctx context.Context, body input) (err error)
	
}

type Server struct{	
	service Service
}

func (s *Server) HandleCommand(rawBody []byte) (err error){
	if len(rawBody) < 2 {
		return errors.New("body is too short")
	}

	commandCode := rawBody[0]

	rawCommandBody := rawBody[1:]

	switch commandCode{
	
	case CommandCodeplayer_move:
		if len(rawCommandBody) < Sizeplayer_move {
			return errors.New("body is too short")
		}

		body, err := Newplayer_move([Sizeplayer_move]byte(rawCommandBody))
		if err != nil {
			return err
		}

		err = s.service.player_move(context.Background(), body)
		if err != nil {
			return err
		}
	
	case CommandCodecreate_player:
		if len(rawCommandBody) < Sizecreate_player {
			return errors.New("body is too short")
		}

		body, err := Newcreate_player([Sizecreate_player]byte(rawCommandBody))
		if err != nil {
			return err
		}

		err = s.service.create_player(context.Background(), body)
		if err != nil {
			return err
		}
	
	case CommandCodeinput:
		if len(rawCommandBody) < Sizeinput {
			return errors.New("body is too short")
		}

		body, err := Newinput([Sizeinput]byte(rawCommandBody))
		if err != nil {
			return err
		}

		err = s.service.input(context.Background(), body)
		if err != nil {
			return err
		}
	
	default:
		return errors.New("unknown command code")
	}

	return nil
}
