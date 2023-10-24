package common

import "testing"

var (
	password = "password"
)

func TestPasswordHash(t *testing.T) {

	hashedPassword, err := PasswordHash(password)

	if err != nil {
		t.Fatal(err)
	}

	if len(hashedPassword) == 0 {
		t.Fatalf("Password was not hashed")
	}

	if hashedPassword == password {
		t.Fatalf("Password was not hashed")
	}
}

func TestPasswordHashCompare(t *testing.T) {

	hashedPassword, err := PasswordHash(password)

	if err != nil {
		t.Fatal(err)
	}

	if len(hashedPassword) == 0 {
		t.Fatalf("Password was not hashed")
	}

	if hashedPassword == password {
		t.Fatalf("Password was not hashed")
	}

	err = PasswordHashCompare(password, hashedPassword)

	if err != nil {
		t.Fatal(err)
	}
}
