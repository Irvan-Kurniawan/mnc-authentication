package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}