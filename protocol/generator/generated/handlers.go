package generated

import "errors"

const (
	CreatePlayer = 20
	PlayerMove   = 21
)

type Protocol interface {
	PlayerMoveHandler(body PlayerMoveBody)
}

type PlayerMoveBody struct {
	Position Position
}

type Position struct {
	X int32
	Y int32
}

type Server struct {
	protocol Protocol
}

func (s *Server) SetProtocol(protocol Protocol) {
	s.protocol = protocol
}

func (s *Server) HandleCommand(payload []byte) error {
	if s.protocol == nil {
		return errors.New("protocol not set")
	}

	if len(payload) < 1 {
		return errors.New("payload too short")
	}

	commandCode := payload[0]

	if len(payload[1:]) < 3 {
		return errors.New("payload too short")
	}

	body := payload[1:]

	switch commandCode {
	case PlayerMove:
		body := PlayerMoveBody{
			Position: Position{X: int32(body[1]), Y: int32(body[2])},
		}

		s.protocol.PlayerMoveHandler(body)
	}

	return nil
}
