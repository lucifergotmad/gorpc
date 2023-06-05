package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id        uint64 `gorm:"primary_key"`
	name 			string `gorm:"not null"`
	username  string `gorm:"unique;not null"`
	email     string `gorm:"unique;not null"`
	password  string `gorm:"not null"`
	created_at time.Time
	updated_at time.Time
	deleted_at time.Time
}
