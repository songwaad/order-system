package auth

import (
	"errors"
	infra "kornkk/Infra"
	"kornkk/entities"
	"kornkk/usecases/user"
)

type Service interface {
	Login(identity, password string) (string, error)
}

type service struct {
	repo user.Repo
}

func NewService(repo user.Repo) Service {
	return &service{repo: repo}
}

func (s *service) Login(identity, password string) (string, error) {
	var user *entities.User
	var err error
	if infra.IsEmail(identity) {
		user, err = s.repo.GetUserByEmail(identity)
	} else {
		user, err = s.repo.GetUserByUsername(identity)
	}

	if err != nil {
		return "", err
	}

	if !infra.CheckPasswordHash(password, user.Password) {
		return "", errors.New("Invalid identity or password")
	}

	t, err := infra.JwtClaims(user)
	if err != nil {
		return "", err
	}

	return t, nil
}
