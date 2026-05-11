package repository_test

import (
	"testing"

	"github.com/elyosemite/everest/internal/domain/user"
	"github.com/elyosemite/everest/internal/infra/repository"
)

func TestInMemoryUserRepository_Save_And_FindByID(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()

	u, err := user.NewUser("John Doe", "john@gmail.com", "12345678900")
	if err != nil {
		t.Fatalf("unexpected error saving user: %v", err)
	}
	u.SetID("user-1")

	if err := repo.Save(u); err != nil {
		t.Fatalf("unexpected error saving user: %v", err)
	}

	found, err := repo.FindByID("user-1")
	if err != nil {
		t.Fatalf("unexpected error finding user: %v", err)
	}
	if found.ID() != "user-1" {
		t.Fatalf("got ID %q, want %q", found.ID(), "user-1")
	}
	if found.Name() != "John Doe" {
		t.Fatalf("got ID %q, want %q", found.Name(), "John Doe")
	}
}

func TestInMemoryUserRepository_FindByID_NotFound(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()

	_, err := repo.FindByID("doesnotexist")
	if err == nil {
		t.Fatalf("expected error for non-existent user, got nil")
	}
}
