package services

import (
	"errors"
	"time"

	"github.com/alejmendez/goApiRest/core/config"
	dtoUser "github.com/alejmendez/goApiRest/modules/users/dto"
	userService "github.com/alejmendez/goApiRest/modules/users/services"
	"github.com/alejmendez/goApiRest/modules/users/utils"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(identity, password string) (string, error) {
	var ud dtoUser.UserDto
	email, err := userService.GetUserByEmail(identity)
	if err != nil {
		return "", errors.New("Error on email")
	}

	user, err := userService.GetUserByUsername(identity)
	if err != nil {
		return "", errors.New("Error on username")
	}

	if email == nil && user == nil {
		return "", errors.New("User not found")
	}

	if email == nil {
		ud = dtoUser.UserDto{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	} else {
		ud = dtoUser.UserDto{
			ID:       email.ID,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	}

	if !utils.CheckPasswordHash(password, ud.Password) {
		return "", errors.New("Invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Get("JWT_SECRET")))
	if err != nil {
		return "", errors.New("Error internal")
	}

	return t, nil
}
