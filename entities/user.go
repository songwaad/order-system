package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"uniqueIndex" json:"username" validate:"required"`
	Email      string `gorm:"uniqueIndex" json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	UserRoleID uint   `json:"user_role_id" validate:"required"`
	// Add relation field so GORM can preload the associated role
	UserRole UserRole `gorm:"foreignKey:UserRoleID" json:"user_role"`
}

type UserRole struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex" json:"name" validate:"required"`
}

type RegisterInput struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

type LoginInput struct {
	Identity string `json:"identity" validate:"required"` // can be username or email
	Password string `json:"password" validate:"required"`
}

type UpdateUserInput struct {
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	UserRoleID uint   `json:"user_role_id,omitempty"`
}
