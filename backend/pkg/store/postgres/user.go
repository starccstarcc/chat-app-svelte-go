package postgres

import (
	"github.com/aletomasella/svelte-go-chat/pkg/common"
	"github.com/aletomasella/svelte-go-chat/pkg/domain"
)

// SQL queries

const (
	createUserQuery = `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
)

func (q *postgresStore) CreateUser(user *domain.User) (*domain.User, error) {

	hashedPassword, err := common.PasswordHash(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	err = q.db.QueryRow(createUserQuery, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
