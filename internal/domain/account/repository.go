package account

type Repository interface {
	Save(account *Account) error
	FindByID(id string) (*Account, error)
	FindByUserID(userID string) ([]*Account, error)
}
