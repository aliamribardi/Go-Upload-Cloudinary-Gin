package Models

import (
	"gorm.io/gorm"
)

type User struct {
	Name	string		`json:"name" gorm:"required"`
	Image	string		`json:"image" gorm:"required"`
	gorm.Model
}