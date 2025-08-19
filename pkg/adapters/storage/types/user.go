package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `gorm:"unique" json:"phone,omitempty"`
	Password  string `json:"password,omitempty"`
}
