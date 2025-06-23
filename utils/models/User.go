package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"size:100;not null" json:"email"`
	Password string `gorm:"size:100;not null" json:"-"`
	Phone    uint   `json:"phone"`
	Otp      uint   `json:"otp"`
	Role     string `gorm:"size:20;default:'user'" json:"role"`
	Isadmin  bool   `json:"isadmin"`
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Phone    uint   `json:"phone" validate:"required"`
}

type LogInRequest struct {
	Phone    uint   `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
