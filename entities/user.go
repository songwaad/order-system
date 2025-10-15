package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"username" validate:"required"`
	Email    string `gorm:"uniqueIndex" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

type Role struct {
	gorm.Model
	Name string `gorm:"uniqueIndex" json:"name" validate:"required"`
}

type RegisterInput struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   uint   `json:"role_id" validate:"required"`
}
