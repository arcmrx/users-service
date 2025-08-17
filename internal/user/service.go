package user

import (
	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

type Service interface {
	CreateUser(email string) (User, error)
	GetUser(userId uuid.UUID) (User, error)
	ListUsers() ([]User, error)
	UpdateUser(userId uuid.UUID, email string) (User, error)
	DeleteUser(userId uuid.UUID) error
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateUser(email string) (User, error) {
	user := User{
		Id:    uuid.New(),
		Email: email,
	}

	err := s.repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) GetUser(userId uuid.UUID) (User, error) {
	return s.repo.GetUser(userId)
}

func (s *service) ListUsers() ([]User, error) {
	return s.repo.ListUsers()
}

func (s *service) UpdateUser(userId uuid.UUID, email string) (User, error) {
	if err := s.repo.UpdateUser(userId, email); err != nil {
		return User{}, err
	}
	return s.repo.GetUser(userId)
}

func (s *service) DeleteUser(userId uuid.UUID) error {
	return s.repo.DeleteUser(userId)
}
