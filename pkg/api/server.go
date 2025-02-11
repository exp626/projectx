// GENERATED CODE
// DO NOT EDIT

package game_protocol

import (
	"context"
	"errors"
	protocol "github.com/exp626/projectx/pkg/protocol"
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
		return protocol.ErrBodyIsTooShort
	}

	commandCode := rawBody[0]

	rawCommandBody := rawBody[1:]

	switch commandCode {
	case CommandCodePlayerMove:
		{
			if len(rawCommandBody) < SizePlayerMoveBody {
				return protocol.ErrBodyIsTooShort
			}

			body, err := NewPlayerMoveBody([SizePlayerMoveBody]byte(rawCommandBody))
			if err != nil {
				return err
			}

			err = s.service.PlayerMove(context.Background(), body)
			if err != nil {
				return err

			}
		}
	case CommandCodeCreatePlayer:
		{
			if len(rawCommandBody) < SizeCreatePlayerBody {
				return protocol.ErrBodyIsTooShort
			}

			body, err := NewCreatePlayerBody([SizeCreatePlayerBody]byte(rawCommandBody))
			if err != nil {
				return err
			}

			err = s.service.CreatePlayer(context.Background(), body)
			if err != nil {
				return err

			}
		}
	case CommandCodeInput:
		{
			if len(rawCommandBody) < SizeInputBody {
				return protocol.ErrBodyIsTooShort
			}

			body, err := NewInputBody([SizeInputBody]byte(rawCommandBody))
			if err != nil {
				return err
			}

			err = s.service.Input(context.Background(), body)
			if err != nil {
				return err

			}
		}

	default:
		return errors.New("unknown command code")
	}

	return nil
}
