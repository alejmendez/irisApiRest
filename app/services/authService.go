package services

import (
	"errors"

	"github.com/alejmendez/goApiRest/app/repositories"
	"github.com/alejmendez/goApiRest/app/utils/jwt"
	"github.com/alejmendez/goApiRest/app/utils/password"
)

func NewAuthService(repo repositories.UserRepository) AuthServices {
	return &authServices{
		userRepository: repo,
	}
}

type AuthServices interface {
	GenerateToken(email string, pass string) (string, error)
}

type authServices struct {
	userRepository repositories.UserRepository
}

func (aS *authServices) GenerateToken(email string, pass string) (string, error) {
	user, _ := aS.userRepository.FindByEmail(email)
	if user == nil {
		return "", errors.New("user not found")
	}

	if !password.Verify(pass, user.Password) {
		return "", errors.New("invalid password")
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	return token, nil
}
