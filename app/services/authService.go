package services

import (
	"errors"
	"time"

	"github.com/alejmendez/goApiRest/app/dto"
	"github.com/alejmendez/goApiRest/app/utils"
	"github.com/alejmendez/goApiRest/core/config"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(identity string, password string) (string, error) {
	var ud dto.UserDto
	email, err := GetUserByEmail(identity)
	if err != nil {
		return "", errors.New("error on email")
	}

	user, err := GetUserByUsername(identity)
	if err != nil {
		return "", errors.New("error on username")
	}

	if email == nil && user == nil {
		return "", errors.New("ser not found")
	}

	if email == nil {
		ud = dto.UserDto{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	} else {
		ud = dto.UserDto{
			ID:       email.ID,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	}

	if !utils.CheckPasswordHash(password, ud.Password) {
		return "", errors.New("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Conf.JwtSecret))
	if err != nil {
		return "", errors.New("error internal")
	}

	return t, nil
}
