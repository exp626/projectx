package server

import (
	"context"
	"github.com/exp626/projectx/pkg/api"
	"log"
	"net"
)

type SampleServer struct{}

func (s *SampleServer) PlayerMove(ctx context.Context, body game_protocol.PlayerMoveBody) (err error) {
	return nil
}

func (s *SampleServer) CreatePlayer(ctx context.Context, body game_protocol.CreatePlayerBody) (err error) {
	return nil
}

func (s *SampleServer) Input(ctx context.Context, body game_protocol.InputBody) (err error) {
	return nil
}

type SampleByteServer struct {
	Serv *game_protocol.Server
}

func (s *SampleByteServer) Start() {
	udpServer, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		_, _, err := udpServer.ReadFrom(buf)
		if err != nil {
			log.Println(err.Error())
		}

		err = s.Serv.HandleCommand(buf)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
