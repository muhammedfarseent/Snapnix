package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	repository "ECOMERS/Src/repository/user"
	"ECOMERS/utils/helper"
	"ECOMERS/utils/models"
)

func UserSignUp(req models.SignUpRequest) error {
	existsEmail, _ := repository.GetUserByEmail(req.Email)
	if existsEmail != nil {
		return errors.New("user already exists")
	}
	existsPhone, _ := repository.CheckUserByPhone(req.Phone)
	if existsPhone != nil {
		return errors.New("user already exists")
	}
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Phone:    req.Phone,
	}
	return repository.CreateUser(&user)

}

func Login(req models.LogInRequest) (string, error) {
	user, err := repository.CheckUserByPhone(req.Phone)
	if err != nil {
		return "", errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid credential")
	}
	token, err := helper.GenarateJwt(user.ID, user.Email, "user")
	if err != nil {
		return "", errors.New("couldn't genarate the token")
	}
	return token, nil
}
