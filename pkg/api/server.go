// GENERATED CODE
// DO NOT EDIT

package game_protocol

import (
	"context"
	"errors"
)

type Server struct {
	service Service
}

func NewServer(service Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) HandleCommand(rawBody []byte) (err error) {
	if len(rawBody) < 2 {
		return errors.New("body is too short")
	}

	commandCode := rawBody[0]

	rawCommandBody := rawBody[1:]

	switch commandCode {

	case CommandCodePlayerMove:
		if len(rawCommandBody) < SizePlayerMoveBody {
			return errors.New("body is too short")
		}

		body, err := NewPlayerMoveBody([SizePlayerMoveBody]byte(rawCommandBody))
		if err != nil {
			return err
		}

		err = s.service.PlayerMove(context.Background(), body)
		if err != nil {
			return err
		}

	case CommandCodeCreatePlayer:
		if len(rawCommandBody) < SizeCreatePlayerBody {
			return errors.New("body is too short")
		}

		body, err := NewCreatePlayerBody([SizeCreatePlayerBody]byte(rawCommandBody))
		if err != nil {
			return err
		}

		err = s.service.CreatePlayer(context.Background(), body)
		if err != nil {
			return err
		}

	case CommandCodeInput:
		if len(rawCommandBody) < SizeInputBody {
			return errors.New("body is too short")
		}

		body, err := NewInputBody([SizeInputBody]byte(rawCommandBody))
		if err != nil {
			return err
		}

		err = s.service.Input(context.Background(), body)
		if err != nil {
			return err
		}

	default:
		return errors.New("unknown command code")
	}

	return nil
}
