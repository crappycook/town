package main

import (
	"bysrpc/server"
)

func main() {
	s := server.NewHostStatusServer()
	s.Run()

	select {}
}
