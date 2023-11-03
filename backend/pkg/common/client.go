package common

import (
	"net"

	"github.com/aletomasella/svelte-go-chat/pkg/domain"
)

func Client(conn net.Conn, messages chan domain.Message) {
	buffer := make([]byte, bufferSize)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			conn.Close()
			messages <- domain.Message{
				Type: domain.ClientDisconnected,
				Conn: conn,
				Body: "",
			}
			return
		}

		msg := string(buffer[:n])

		if n > 0 && len(msg) == n {
			messages <- domain.Message{
				Type: domain.MessageReceived,
				Body: msg,
				Conn: conn,
			}
		}
	}

}
