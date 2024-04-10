package main

import (
	"flag"

	"liars-dice/internal/client"
	"liars-dice/internal/server"
)

var serverFlag = flag.Bool("server", false, "starts the program as the server")

func main() {
	flag.Parse()

	if *serverFlag {
		server.Run()
	} else {
		client.Run()
	}
}
