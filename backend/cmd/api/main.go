package main

import (
	"fmt"
	"net"
	"os"
)

const (
	Port = "3000"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())

	// Go All Format Types : https://pkg.go.dev/fmt

	for i := 0; i < 1000; i++ {
		conn.Write([]byte("Counting to 1000: " + fmt.Sprintf("%d", i) + "\n"))
	}

	conn.Write([]byte("Bye!\n"))

	fmt.Printf("Connection from %s closed\n", conn.RemoteAddr().String())
}

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

		go handleConnection(conn)
	}
}
