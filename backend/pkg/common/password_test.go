package common

import "testing"

func TestPasswordHash(t *testing.T) {

	password := "password"

	hashedPassword, err := PasswordHash(password)

	if err != nil {
		t.Fatal(err)
	}

	if hashedPassword == password {
		t.Fatalf("Password was not HASHED")
	}
}
