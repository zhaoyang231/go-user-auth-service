package service

import (
	"go-user-auth-service/internal/config"
	"go-user-auth-service/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (s *UserService) CreateUser(user *model.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}
func (s *UserService) GetAllUsers() ([]model.User, error) {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
