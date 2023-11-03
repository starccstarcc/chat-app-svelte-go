package domain

import (
	"net"
	"time"
)

// This is the way tp declare enums in Go (Starts in 0)
// IOTA means auto increment
const (
	ClientConnected MessageType = iota + 1
	ClientDisconnected
	MessageReceived
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Client struct {
	Conn        net.Conn
	LastMessage time.Time
}

type MessageType int

type Message struct {
	Type MessageType
	Body string
	Conn net.Conn
}
