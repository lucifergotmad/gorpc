package user

import (
	"fmt"

	"gorm.io/gorm"
)

type GetUserBatchRequest struct {
	page int
	size int
}

type Repository interface {
	CreateUser(user *User) error
	GetUser(userID string) (*User, error)
	GetUserBatch(query GetUserBatchRequest) ([]*User, error)
}

type UserRepository struct {
	Repository
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {
			return fmt.Errorf("could not create user: %v", result.Error)
	}
	return nil
}

func (r *UserRepository) GetUser(userID string) (*User, error) {
	var user User
	
	err := r.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

