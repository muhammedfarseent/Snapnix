package models

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	UserID uint       `gorm:"unique;notnull"`
	Name   string     `gorm:"name"`
	Email  string     `gorm:"size:100" json:"email"`
	Phone  string     `gorm:"size:20" json:"phone"`
	DOB    *time.Time `json:"dob,omitempty"`
	Gender string     `gorm:"size:20" json:"gender"`
}

type CreateUSerPrifileRequest struct {
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
	DOB    string `json:"dob"`
	Gender string `json:"gender"`
}
