package repositories

import (
	model "github.com/alejmendez/goApiRest/app/models"
	"github.com/alejmendez/goApiRest/app/utils"
	"github.com/jinzhu/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		Db: db,
	}
}

type UserRepository interface {
	ListByWhere(userW *model.User) ([]*model.User, error)
	FindByWhere(userW *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Find(id string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(id string, userParam *model.User) (*model.User, error)
	Delete(id string) (bool, error)
}

type userRepository struct {
	Db *gorm.DB
}

func (uR *userRepository) ListByWhere(userW *model.User) ([]*model.User, error) {
	var list []*model.User
	if err := uR.Db.Where(&userW).Find(&list).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return list, nil
}

func (uR *userRepository) FindByWhere(userW *model.User) (*model.User, error) {
	var user model.User
	if err := uR.Db.Where(&userW).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (uR *userRepository) FindByEmail(email string) (*model.User, error) {
	return uR.FindByWhere(&model.User{Email: email})
}

func (uR *userRepository) Find(id string) (*model.User, error) {
	i, _ := utils.StringToUint(id)
	return uR.FindByWhere(&model.User{ID: i})
}

func (uR *userRepository) Create(user *model.User) (*model.User, error) {
	user.Password = utils.GenerateHash(user.Password)

	if err := uR.Db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (uR *userRepository) Update(id string, userParam *model.User) (*model.User, error) {
	user, err := uR.Find(id)

	if err != nil {
		return nil, err
	}

	user.Username = userParam.Username
	user.Email = userParam.Email

	if user.Password != "" {
		user.Password = utils.GenerateHash(user.Password)
	}

	if err := uR.Db.Save(&userParam).Error; err != nil {
		return nil, err
	}

	return userParam, nil
}

func (uR *userRepository) Delete(id string) (bool, error) {
	user, err := uR.Find(id)

	if err != nil {
		return false, err
	}

	if err := uR.Db.Unscoped().Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
