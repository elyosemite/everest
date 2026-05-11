package repository

import (
	"errors"

	"github.com/elyosemite/everest/internal/domain/account"
)

type InMemoryAccountRepository struct {
	store map[string]*account.Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		store: make(map[string]*account.Account),
	}
}

func (r *InMemoryAccountRepository) Save(a *account.Account) error {
	r.store[a.ID()] = a
	return nil
}

func (r *InMemoryAccountRepository) FindByID(id string) (*account.Account, error) {
	a, ok := r.store[id]
	if !ok {
		return nil, errors.New("account not found")
	}
	return a, nil
}

func (r *InMemoryAccountRepository) FindByUserID(userID string) ([]*account.Account, error) {
	var result []*account.Account
	for _, a := range r.store {
		if a.UserID() == userID {
			result = append(result, a)
		}
	}
	return result, nil
}
