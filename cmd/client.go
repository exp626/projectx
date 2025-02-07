package main

import (
	"context"
	game_protocol "github.com/exp626/projectx/pkg/api"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	udpServer, err := net.ResolveUDPAddr("udp", ":1053")

	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Listen failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()

	cli := game_protocol.NewClient(conn)

	t := time.NewTicker(time.Millisecond * 16)
	for _ = range t.C {
		err = cli.PlayerMove(context.Background(), game_protocol.PlayerMoveBody{
			EntityId: 1,
			Position: game_protocol.Vector{
				X: 1,
				Y: 1,
			},
			Direction: game_protocol.Vector{
				X: 2,
				Y: 2,
			},
		})
		if err != nil {
			log.Println(err)
		}
	}
}
