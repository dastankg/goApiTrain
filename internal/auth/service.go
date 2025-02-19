package auth

import (
	"apiProject/internal/user"
	"errors"
	"golang.org/x/crypto/bcrypt"
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
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Name:     name,
		Email:    email,
		Password: string(hashPass),
	}
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	exist, _ := service.UserRepository.GetByEmail(email)
	if exist == nil {
		return "", errors.New(ErrNoExists)
	}
	err := bcrypt.CompareHashAndPassword([]byte(exist.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrNoExists)
	}
	return exist.Email, nil

}
