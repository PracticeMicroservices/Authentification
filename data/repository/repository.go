package repository

import "authentification/data/entities"

type Repository interface {
	GetAll() ([]*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetOne(id int) (*entities.User, error)
	Update(user entities.User) error
	DeleteByID(id int) error
	Insert(user entities.User) (int, error)
	ResetPassword(password string, user entities.User) error
	PasswordMatches(plainText string, user entities.User) (bool, error)
}
