package user

import (
	"kornkk/entities"

	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(user *entities.User) (*entities.User, error)
	// GetUsers() ([]entities.User, error)
	// GetUserByID(id uint) (*entities.User, error)
	// UpdateUser(user *entities.User) (*entities.User, error)
	// DeleteUser(id uint) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) CreateUser(user *entities.User) (*entities.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
