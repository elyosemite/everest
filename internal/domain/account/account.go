package account

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	id          string
	userID      string
	accountType AccountType
	balance     float64
	createdAt   time.Time
	active      bool
}

func NewAccount(userID string, accountType AccountType, initialBalance float64) (*Account, error) {
	if initialBalance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}

	return &Account{
		id:          uuid.New().String(),
		userID:      userID,
		accountType: accountType,
		balance:     initialBalance,
		createdAt:   time.Now().UTC(),
		active:      true,
	}, nil
}

func (a *Account) Type() AccountType { return a.accountType }

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if a.balance < amount {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deactivate() {
	a.active = false
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) UserID() string {
	return a.userID
}
