package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Email   string
	State   string
	Phone   string
	Address string
}
