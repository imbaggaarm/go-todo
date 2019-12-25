package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email" gorm:"unique;not null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
}
