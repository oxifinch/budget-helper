package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type Budget struct {
	gorm.Model
	StartDate string
	EndDate   string
	UserID    uint
	User      User
}
