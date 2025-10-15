package user

import (
	infra "kornkk/Infra"
	"kornkk/entities"

	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUserByID(id uint) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(id uint, user *entities.User) (*entities.User, error)
	DeleteUser(id uint) error
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

func (r *repo) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Preload("UserRole").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
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

func (r *repo) UpdateUser(id uint, updateUser *entities.User) (*entities.User, error) {
	user, err := r.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	if updateUser.Username != "" {
		user.Username = updateUser.Username
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	if updateUser.Password != "" {
		user.Password, err = infra.HashPassword(updateUser.Password)
		if err != nil {
			return nil, err
		}
	}

	if updateUser.UserRoleID != 0 {
		user.UserRoleID = updateUser.UserRoleID
	}

	if err := r.db.Omit("UserRole").Save(&user).Error; err != nil {
		return nil, err
	}

	user, _ = r.GetUserByID(user.ID)
	return user, nil
}

func (r *repo) DeleteUser(id uint) error {
	return r.db.Delete(&entities.User{}, id).Error
}
