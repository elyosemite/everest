package user

import (
	"errors"
	"regexp"
	"strings"

	account "github.com/elyosemite/everest/internal/domain/account"
	address "github.com/elyosemite/everest/internal/domain/address"
)

type User struct {
	id       string
	name     string
	email    string
	identity string
	address  []*address.Address
	accounts []*account.Account
}

func NewUser(name, email, document string) (*User, error) {
	if err := validateUser(name, email, document); err != nil {
		return nil, err
	}

	return &User{
		name:     strings.TrimSpace(name),
		email:    strings.ToLower(strings.TrimSpace(email)),
		identity: strings.TrimSpace(document),
	}, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) SetID(id string) {
	u.id = id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func cleanDocument(doc string) string {
	return strings.ReplaceAll(doc, "[^0-9]", "")
}

func validateUser(name, email, document string) error {
	if len(name) < 5 {
		return errors.New("name must be at least 5 characters")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	if len(cleanDocument(document)) != 11 && len(cleanDocument(document)) != 14 {
		return errors.New("document must be CPF (11 digits) or CNPJ (14 digits)")
	}

	return nil
}

func (u *User) AddAccount(account *account.Account) error {
	if account == nil {
		return errors.New("account cannot be nil")
	}
	u.accounts = append(u.accounts, account)
	return nil
}

func (u *User) Accounts() []account.Account {
	accounts := make([]account.Account, len(u.accounts))
	for i, acc := range u.accounts {
		accounts[i] = *acc
	}
	return accounts
}
