package main

import (
	"fmt"
	game_protocol "github.com/exp626/projectx/pkg/api"
	"github.com/exp626/projectx/server"
	"os"
	"os/signal"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("myprogram.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Run your program here

	srv := server.SampleByteServer{
		Serv: game_protocol.NewServer(&server.SampleServer{}),
	}

	go func() {
		srv.Start()
	}()

	sig := make(chan os.Signal)

	signal.Notify(sig, os.Interrupt)

	<-sig
}
