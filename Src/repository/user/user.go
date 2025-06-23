package user

import (
	"errors"

	"gorm.io/gorm"

	"SNAPNIX/database"
	"SNAPNIX/utils/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	res := database.DB.Where("email=?", email).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func CheckUserByPhone(Phone uint) (*models.User, error) {
	var user models.User

	if err := database.DB.Where("phone=?", Phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return database.DB.Create(&user).Error
}
