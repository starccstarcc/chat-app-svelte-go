package postgres

import "github.com/aletomasella/svelte-go-chat/pkg/domain"

func (q *postgresStore) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = 1

	return user, nil
}
