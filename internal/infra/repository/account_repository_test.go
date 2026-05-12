package repository_test

import (
	"testing"

	"github.com/elyosemite/everest/internal/domain/account"
	"github.com/elyosemite/everest/internal/infra/repository"
)

func TestInMemoryAccountRepository_Save_And_FindByID(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()

	a, err := account.NewAccount("user-1", account.Checking, 100)
	if err != nil {
		t.Fatalf("unexpected error creating account: %v", err)
	}

	if err := repo.Save(a); err != nil {
		t.Fatalf("unexpected error saving account: %v", err)
	}

	found, err := repo.FindByID(a.ID())
	if err != nil {
		t.Fatalf("unexpected error finding account: %v", err)
	}
	if found.ID() != a.ID() {
		t.Errorf("got ID %q, want %q", found.ID(), a.ID())
	}
	if found.Balance() != 100 {
		t.Errorf("got balance %v, want %v", found.Balance(), 100)
	}
}

func TestInMemoryAccountRepository_FindByID_NotFound(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()

	_, err := repo.FindByID("nao-existe")
	if err == nil {
		t.Fatal("expected error for non-existent account, got nil")
	}
}

func TestInMemoryAccountRepository_FindAccountByUserID(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()

	a1, _ := account.NewAccount("user-1", account.Checking, 100)
	a2, _ := account.NewAccount("user-1", account.Savings, 200)
	a3, _ := account.NewAccount("user-2", account.Checking, 50)

	repo.Save(a1)
	repo.Save(a2)
	repo.Save(a3)

	results, err := repo.FindByUserID("user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("got %d accounts, want 2", len(results))
	}
}

func TestInMemoryAccountRepository_FindByUserID_NotFound(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()

	results, err := repo.FindByUserID("nao-existe")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 0 {
		t.Errorf("got %d accounts, want 0", len(results))
	}
}

func TestInmemoryRepository_TrySaveExistingAccount(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()

	newAccount, _ := account.NewAccount("user-1", account.Savings, 100_000)
	repo.Save(newAccount)

	err := repo.Save(newAccount)
	if err == nil {
		t.Fatal("expected error when saving duplicate account, got nil")
	}
}

func TestInMemoryAccountRepository_FindByID_ReturnsCorrectData(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()

	a, _ := account.NewAccount("user-1", account.Checking, 250.0)
	repo.Save(a)

	found, err := repo.FindByID(a.ID())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if found.ID() != a.ID() {
		t.Fatalf("ID: got %q, want %q", found.ID(), a.ID())
	}
	if found.Type() != a.Type() {
		t.Errorf("Type: got %v, want %v", found.Type(), account.Savings)
	}
	if found.Balance() != 250.0 {
		t.Errorf("Balance: got %v, want %v", found.Balance(), 250.0)
	}
}
