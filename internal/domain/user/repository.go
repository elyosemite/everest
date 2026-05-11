package user

type Repository interface {
	Save(user *User) error
	FindIdID(id string) (*User, error)
}
