package services

import (
	"errors"

	"event-reporting/app/database/pgsql/repository"
	"event-reporting/app/models"
	"event-reporting/app/utils/hashing"
)

type UserService interface {
	Register(name, email, password string) (*models.User, error)
	Verify(email, password string) (*models.User, error)
}

type userService struct{ repo repository.UserRepository }

func NewUserService(r repository.UserRepository) UserService { return &userService{repo: r} }

func (s *userService) Register(name, email, password string) (*models.User, error) {
	hash, err := hashing.Hash(password)
	if err != nil { return nil, err }
	u := &models.User{Name: name, Email: email, PasswordHash: hash}
	if err := s.repo.Create(u); err != nil { return nil, err }
	return u, nil
}

func (s *userService) Verify(email, password string) (*models.User, error) {
	u, err := s.repo.ByEmail(email)
	if err != nil { return nil, err }
	if err := hashing.Compare(u.PasswordHash, password); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return u, nil
}