package account

import (
	"errors"
	"time"
)

type Account struct {
	id        string
	userID    string
	balance   float64
	createdAt time.Time
	active    bool
}

func NewAccount(userID string, initialBalance float64) (*Account, error) {
	if initialBalance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}

	return &Account{
		userID:    userID,
		balance:   initialBalance,
		createdAt: time.Now().UTC(),
		active:    true,
	}, nil
}

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
