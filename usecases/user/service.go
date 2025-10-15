package user

import (
	infra "kornkk/Infra"
	"kornkk/entities"
)

type Service interface {
	Register(user *entities.RegisterInput) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUserByID(id uint) (*entities.User, error)
	UpdateUser(id uint, user *entities.UpdateUserInput) (*entities.User, error)
	DeleteUser(id uint) error
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo: repo}
}

func (s *service) Register(user *entities.RegisterInput) (*entities.User, error) {
	passwordHash, err := infra.HashPassword(user.Password)
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

func (s *service) GetAllUsers() ([]entities.User, error) {
	return s.repo.GetAllUsers()
}

func (s *service) GetUserByID(id uint) (*entities.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *service) UpdateUser(id uint, updateUser *entities.UpdateUserInput) (*entities.User, error) {
	return s.repo.UpdateUser(id, &entities.User{
		Username:   updateUser.Username,
		Email:      updateUser.Email,
		Password:   updateUser.Password,
		UserRoleID: updateUser.UserRoleID,
	})
}

func (s *service) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
