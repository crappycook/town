package main

import (
	"bys/bootstrap"
	"bys/dal"
	"bys/rpc"
	"bys/server/handlers"
)

func main() {
	config := bootstrap.LoadConfig()

	// rpc caller init
	rpc.InitClient(config)

	// dal init
	dal.Init(config)

	server := handlers.NewHttpServer(
		handlers.WithBootConfig(config),
	)
	go server.Run()

	select {}
}
