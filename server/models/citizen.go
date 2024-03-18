package models

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type CitizenModel struct {
	gorm.Model
	FirstName   string    `gorm:"size:255; not null"`
	LastName    string    `gorm:"size:255; not null"`
	DateOfBirth time.Time `gorm:"not null"`
	Gender      string    `gorm:"size:255; not null"`
	Address     string    `gorm:"size:255; not null"`
	City        string    `gorm:"size:255; not null"`
	State       string    `gorm:"size:255; not null"`
	Pincode     string    `gorm:"size:255; not null"`
}

// func (citizen *CitizenModel) CreateCitizen() error {
// 	db := database.GetDB()
// 	err := db.Create(&citizen).Error
// 	return err
// }
