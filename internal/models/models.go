package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}

type Space struct {
	gorm.Model
	Name        string
	Description string
	OwnerID     uint
	Owner       User `gorm:"foreignKey:OwnerID"`
}

type Interest struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

type UserInterest struct {
	gorm.Model
	UserID     uint
	InterestID uint
	User       User     `gorm:"foreignKey:UserID"`
	Interest   Interest `gorm:"foreignKey:InterestID"`
}

type SpaceMember struct {
	gorm.Model
	UserID  uint
	SpaceID uint
	User    User  `gorm:"foreignKey:UserID"`
	Space   Space `gorm:"foreignKey:SpaceID"`
}