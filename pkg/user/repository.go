package user

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *User) error {
	// Database operation for creating a user.
}

func (r *UserRepository) GetUser(userID int) (*User, error) {
	// Database operation for getting a user.
}
