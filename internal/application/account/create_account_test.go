package account_test

import (
	"testing"

	usecase "github.com/elyosemite/everest/internal/application/account"
	domainaccount "github.com/elyosemite/everest/internal/domain/account"
	"github.com/elyosemite/everest/internal/domain/user"
	"github.com/elyosemite/everest/internal/infra/repository"
)

func TestCreateAccount_Success_Checking(t *testing.T) {
	userRepo := repository.NewInMemoryUserRepository()
	accountRepo := repository.NewInMemoryAccountRepository()

	u, err := user.NewUser("John Doe", "johndoe@mail.com", "12345678901")
	if err != nil {
		t.Fatalf("unexpected error creating user: %v", err)
	}
	if err := userRepo.Save(u); err != nil {
		t.Fatalf("unexpected error saving user: %v", err)
	}

	uc := usecase.NewCreateAccountUseCase(userRepo, accountRepo)

	input := usecase.CreateAccountInput{
		UserID:         u.ID(),
		AccountType:    domainaccount.Checking,
		InitialBalance: 500.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		t.Fatalf("unexpected no error, got: %v", err)
	}
	if output == nil {
		t.Fatalf("expected output, got nil")
	}
}
