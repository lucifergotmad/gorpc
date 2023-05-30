package user

import (
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByID(userID uint64) (*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(userID uint64) (*User, error) {
	user := &User{}
	if err := r.db.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

