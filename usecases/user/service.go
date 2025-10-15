package user

import (
	"kornkk/entities"
	"kornkk/middleware"
)

type Service interface {
	Register(user *entities.RegisterInput) (*entities.User, error)
	GetUserByID(id uint) (*entities.User, error)
	DeleteUser(id uint) error
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo: repo}
}

func (s *service) Register(user *entities.RegisterInput) (*entities.User, error) {
	passwordHash, err := middleware.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	createdUser, err := s.repo.CreateUser(&entities.User{
		Username:   user.Username,
		Email:      user.Email,
		Password:   passwordHash,
		UserRoleID: user.RoleID,
	})
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *service) GetUserByID(id uint) (*entities.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *service) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
