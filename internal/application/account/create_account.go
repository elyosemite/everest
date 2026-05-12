package account

import (
	domainaccount "github.com/elyosemite/everest/internal/domain/account"
	"github.com/elyosemite/everest/internal/domain/user"
)

type CreateAccountInput struct {
	UserId         string
	AccountType    domainaccount.AccountType
	InitialBalance float64
}

type CreatetAccountOutput struct {
	AccountID string
	Balance   float64
	Type      string
}

type CreateAccountUserCase struct {
	userRepo    user.Repository
	accountRepo domainaccount.Repository
}

func (uc *CreateAccountUserCase) Execute(input CreateAccountInput) (*CreatetAccountOutput, error)
