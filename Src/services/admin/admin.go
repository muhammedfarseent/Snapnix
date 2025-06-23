package services

import (
	"errors"

	repository "ECOMERS/Src/repository/admin"
	"ECOMERS/utils/helper"
	"ECOMERS/utils/models"
)

type AuthService struct {
	Repo *repository.AdminRepository
}

func (s *AuthService) SignupAdmin(name, email, password string) (*models.User, error) {
	hasPassword, _ := helper.HashPassword(password)
	user := models.User{Name: name, Email: email, Password: string(hasPassword), Isadmin: true}
	return &user, nil
}

func (s *AuthService) LoginAdmin(email, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil || !user.Isadmin {
		return "", errors.New("admin not found or not autherized")
	}
	if !helper.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid")
	}
	token, err := helper.GenarateAdminJwt(user.ID, user.Isadmin)
	return token, err
}
