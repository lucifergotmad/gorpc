package user

import (
	"log"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByID(userID uint64) (*User, error)
	CreateUser(user *User) (error)
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
	log.Println("Hallo dari repos")
	user := &User{}
	if err := r.db.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user *User) (error) {
	err := r.db.Create(&user)

	if err != nil {
		return err.Error
	}

	return nil
}