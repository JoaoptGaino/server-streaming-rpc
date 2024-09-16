package main

import (
	"github.com/joaoptgaino/server-streaming-rpc/client"
	"github.com/joaoptgaino/server-streaming-rpc/server"
)

func main() {
	go server.Run()

	client.Run()
}
