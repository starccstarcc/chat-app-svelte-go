package postgres

import (
	"testing"

	"github.com/aletomasella/svelte-go-chat/pkg/domain"
)

func TestCreateUser(t *testing.T) {

	pStre := NewPostgresStore(nil)

	user := &domain.User{
		Email:    "test@test.com",
		Password: "password",
		Name:     "Test User",
	}

	createdUser, err := pStre.CreateUser(user)

	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == 0 {
		t.Fatal("Expected user ID to be set")
	}

	if createdUser.Email != user.Email {
		t.Fatalf("Expected email to be %s, got %s", user.Email, createdUser.Email)
	}

	if createdUser.Password == user.Password {
		t.Fatalf("Password was not HASHED")
	}

	if createdUser.Name != user.Name {
		t.Fatalf("Expected name to be %s, got %s", user.Name, createdUser.Name)
	}
}
