package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"ECOMERS/config"
)

type AuthUserClaims struct {
	UserId uint   `json:"user"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenarateJwt(userID uint, email string, role string) (string, error) {
	config, _ := config.LoadConfig()

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetTokenFromHeader(header string) string {
	if len(header) >= 7 && header[:7] == "Bearer " {
		return header[7:]
	}
	return header
}
