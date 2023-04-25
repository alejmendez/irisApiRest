package services

import (
	"errors"

	"github.com/alejmendez/goApiRest/app/repositories"
	"github.com/alejmendez/goApiRest/app/utils"
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

	if !utils.VerifyHash(pass, user.Password) {
		return "", errors.New("invalid password")
	}

	token := utils.JwtGenerate(&utils.TokenPayload{
		ID: user.ID,
	})

	return token, nil
}
