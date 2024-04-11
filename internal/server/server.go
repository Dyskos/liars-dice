package server

import (
	"fmt"
	"net"
	"os"
)

func Run() {
	port := "10150"
	// TODO: use configurable port
	ln, err := net.Listen("tcp", ":" + port)
	if err != nil {
		fmt.Printf("! %s\r\n", err)
		os.Exit(1)
	}
	fmt.Printf("[*] Listening on: %v\r\n", ln.Addr())

	for {
		cn, err := ln.Accept()
		if err != nil {
			fmt.Printf("[!] %s\r\n", err)
		}
		go handleConnection(cn)
	}
}

func handleConnection(cn net.Conn) {
	fmt.Printf("[*] connection established: %v\r\n", cn)
	return
}
