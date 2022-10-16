package services

import (
	model "github.com/alejmendez/goApiRest/app/models"
	"github.com/alejmendez/goApiRest/app/repositories"
	"github.com/alejmendez/goApiRest/app/utils/password"
)

type UserService interface {
	Valid(id string, pass string) bool
	Get(id string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(id string, userParam *model.User) (*model.User, error)
	Delete(id string) (bool, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repository: repo,
	}
}

func (uS *userService) Valid(id string, pass string) bool {
	user, err := uS.repository.Find(id)
	if err != nil || user == nil {
		return false
	}

	return password.Verify(pass, user.Password)
}

func (uS *userService) Get(id string) (*model.User, error) {
	return uS.repository.Find(id)
}

func (uS *userService) Create(user *model.User) (*model.User, error) {
	return uS.repository.Create(user)
}

func (uS *userService) Update(id string, userParam *model.User) (*model.User, error) {
	return uS.repository.Update(id, userParam)
}

func (uS *userService) Delete(id string) (bool, error) {
	return uS.repository.Delete(id)
}
