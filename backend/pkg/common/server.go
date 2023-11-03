package common

import (
	"fmt"
	"net"
	"time"

	"github.com/aletomasella/svelte-go-chat/pkg/domain"
)

const (
	SafeMode           = false
	ClientDisconnected = "Client %s Disconnected\n"
	ClientConnected    = "Client %s Connected\n"
	DefaultCase        = "ERROR : Message type not found for Message %s\n"
)

func safeAdress(addr net.Conn) string {
	if SafeMode {
		return "[SAFEMODE]"
	}
	return addr.RemoteAddr().String()
}

func handleMessages(msg domain.Message, clients map[string]*domain.Client) {

	address := msg.Conn.RemoteAddr().String()
	switch msg.Type {
	case domain.ClientConnected:
		clients[address] = &domain.Client{
			Conn:        msg.Conn,
			LastMessage: time.Now(),
		}
		fmt.Printf(ClientConnected, safeAdress(msg.Conn))

	case domain.ClientDisconnected:
		delete(clients, address)
		msg.Conn.Close()
		fmt.Printf(ClientDisconnected, safeAdress(msg.Conn))

	case domain.MessageReceived:
		fmt.Printf(MessageReceived, msg.Body)
		// Send message to all clients except the one who sent it
		for _, client := range clients {
			if client.Conn.RemoteAddr().String() != address {
				if client.Conn != nil {
					client.Conn.Write([]byte(msg.Body))

				}
			}
		}
	default:
		fmt.Printf(DefaultCase, msg.Body)
	}
}

func ChannelServer(msgChannel chan domain.Message) {
	clients := make(map[string]*domain.Client)

	// Infinite loop
	for {
		msg := <-msgChannel
		handleMessages(msg, clients)
	}

}
