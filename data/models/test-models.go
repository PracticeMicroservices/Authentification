package models

import (
	"authentification/data/entities"
	"authentification/data/repository"
	"database/sql"
	"time"
)

type TestRepository struct {
	Conn *sql.DB
}

func NewTestRepository(db *sql.DB) repository.Repository {
	return &TestRepository{
		Conn: db,
	}
}

func (u *TestRepository) GetAll() ([]*entities.User, error) {
	var users []*entities.User

	return users, nil
}

// GetByEmail returns one user by email
func (u *TestRepository) GetByEmail(email string) (*entities.User, error) {
	user := entities.User{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "me@here.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *TestRepository) GetOne(id int) (*entities.User, error) {
	user := entities.User{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "me@here.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *TestRepository) Update(user entities.User) error {
	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *TestRepository) DeleteByID(id int) error {
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *TestRepository) Insert(user entities.User) (int, error) {
	return 2, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *TestRepository) ResetPassword(password string, user entities.User) error {
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *TestRepository) PasswordMatches(plainText string, user entities.User) (bool, error) {
	return true, nil
}
