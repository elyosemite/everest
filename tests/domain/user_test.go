package domain

import (
	"testing"

	user "github.com/elyosemite/everest/internal/domain/user"
)

func TestUserCreationWithValidCPF(t *testing.T) {
	//address, err := address.NewAddress("Main St", "123", "São Paulo", "ST", "12345-678")

	user, err := user.NewUser("John Doe", "john.doe@mail.com", "12345678900")

	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	want := "John Doe"

	if user.Name() != want {
		t.Errorf("got %q, want %q", user.Name(), want)
	}
}
