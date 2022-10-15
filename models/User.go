package models

import "gorm.io/gorm"

type User struct {

	gorm.Model

	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null; unique"`
	Task 	[]Task `json:"task"`
}