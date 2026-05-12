package account

import (
	"fmt"

	domainaccount "github.com/elyosemite/everest/internal/domain/account"
	"github.com/elyosemite/everest/internal/domain/user"
)

type CreateAccountInput struct {
	UserID         string
	AccountType    domainaccount.AccountType
	InitialBalance float64
}

type CreateAccountOutput struct {
	AccountID string
	Balance   float64
	Type      string
}

type CreateAccountUserCase struct {
	userRepo    user.Repository
	accountRepo domainaccount.Repository
}

func (uc *CreateAccountUserCase) Execute(input CreateAccountInput) (*CreateAccountOutput, error) {
	u, err := uc.userRepo.FindByID(input.UserID)
	if err != nil {
		return nil, fmt.Errorf("creaet account: %w", err)
	}

	a, err := domainaccount.NewAccount(u.ID(), input.AccountType, input.InitialBalance)
	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}

	if err := uc.accountRepo.Save(a); err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}

	u.AddAccount(a)

	if err := uc.userRepo.Save(u); err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}

	return &CreateAccountOutput{
		AccountID: a.ID(),
		Balance:   a.Balance(),
		Type:      a.Type().String(),
	}, nil
}
