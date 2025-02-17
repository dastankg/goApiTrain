package auth

import (
	"apiProject/internal/user"
	"errors"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewUserService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(name, email, password string) (string, error) {
	exist, _ := service.UserRepository.GetByEmail(email)
	if exist != nil {
		return "", errors.New(ErrUserExists)
	}
	user := &user.User{
		Name:     name,
		Email:    email,
		Password: "",
	}
	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
