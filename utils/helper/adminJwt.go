package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"SNAPNIX/config"
)

func GenarateAdminJwt(userID uint, Isadmin bool) (string, error) {
	config, _ := config.LoadConfig()

	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": Isadmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	jwtSecret := []byte(config.Key)
	return token.SignedString(jwtSecret)
}
func ParseToken(tokenstr string) (jwt.MapClaims, error) {
	config, _ := config.LoadConfig()
	jwtSecret := []byte(config.Key)

	token, err := jwt.ParseWithClaims(tokenstr, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected siging method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// take the clamims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
