package main

import (
	"fmt"
	"net"
	"os"
)

const (
	Port = "3000"
)

func main() {
	ln, err := net.Listen("tcp", ":"+Port)

	if err != nil {
		fmt.Printf("ERROR: Trying to listen in port %s, but failed because of %s\n", Port, err)
		os.Exit(1)
	}

	fmt.Printf("Listening in port %s\n", Port)

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Printf("ERROR: Trying to accept connection, but failed because of %s\n", err)
			continue
		}

		fmt.Println("Connection accepted from", conn.RemoteAddr())
		conn.Write([]byte("Hello, world!\n"))
		conn.Close()
		//go handleConnection(conn)
	}
}
