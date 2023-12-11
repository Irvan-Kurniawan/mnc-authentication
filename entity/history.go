package entity

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	Username   string `json:"username"`
	Action     string `json:"action"`
	Nominal    int `json:"balance"`
	ActionDate time.Time  `json:"actiondate"`
}
