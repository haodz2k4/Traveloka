package models

import "time"

type Users struct {
	UserId    string `gorm:"primary_key;type:char(4)"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
	Email     string `gorm:"type:varchar(50);unique;not null"`
	Phone     string `gorm:"type:varchar(50);not null"`
	status    string `gorm:"type:enum('active','inactive');not null"`
	createdAt time.Time
	updatedAt time.Time
}
