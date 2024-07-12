package services

import (
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id string) (models.User, error)
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetUsers() ([]models.User, error) {
	return s.repo.GetUsers()
}

func (s *UserServiceImpl) GetUser(id string) (models.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserServiceImpl) CreateUser(user models.User) (models.User, error) {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
		if err != nil {
			return models.User{}, err
		}

		user.Password = string(hashedPassword)
	}
	return s.repo.CreateUser(user)
}

func (s *UserServiceImpl) UpdateUser(user models.User) (models.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *UserServiceImpl) DeleteUser(id string) (models.User, error) {
	return s.repo.DeleteUser(id)
}
