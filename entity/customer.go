package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password gorm:"not null"`
	Balance  int `json:"balance" gorm:"default:0"`
	IsLogin bool `json:"isLogin" gorm:"default:false"`
}

func (customer *Customer) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	customer.Password = string(bytes)
	return nil
}
func (customer *Customer) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}