package repository

import (
	"errors"

	"github.com/elyosemite/everest/internal/domain/user"
)

type InMemoryUserRepository struct {
	store map[string]*user.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		store: make(map[string]*user.User),
	}
}

func (r *InMemoryUserRepository) Save(u *user.User) error {
	r.store[u.ID()] = u
	return nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*user.User, error) {
	u, ok := r.store[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}
