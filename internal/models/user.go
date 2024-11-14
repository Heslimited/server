package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Tasks    []Task
}
