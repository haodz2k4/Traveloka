package models

import (
	"time"
)

type Users struct {
	UserID    string    `gorm:"primaryKey;type:char(4);not null" json:"user_id"`
	FirstName string    `gorm:"type:varchar(50);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(50);not null" json:"last_name"`
	Email     string    `gorm:"type:varchar(50);unique;not null" json:"email"`
	Phone     string    `gorm:"type:varchar(50);not null" json:"phone"`
	Status    string    `gorm:"type:enum('active','inactive');not null" json:"status"`
	Deleted   bool      `gorm:"type:boolean;default:false" json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
