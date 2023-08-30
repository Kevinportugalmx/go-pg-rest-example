package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name" required:"true"`
	LastName  string `gorm:"not null" json:"last_name" required:"true"`
	Email     string `gorm:"not null;uniqueIndex" json:"email" required:"true"`
	Password  string `gorm:"not null" json:"password" required:"true"`
}
