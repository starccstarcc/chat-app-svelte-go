package domain

import "time"

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
