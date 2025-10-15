package user

import (
	"kornkk/entities"

	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(user *entities.User) (*entities.User, error)
	// GetUsers() ([]entities.User, error)
	GetUserByID(id uint) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
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

	user, _ = r.GetUserByID(user.ID)
	return user, nil
}

func (r *repo) GetUserByID(id uint) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", id).Preload("UserRole").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("username = ?", username).Preload("UserRole").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).Preload("UserRole").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
