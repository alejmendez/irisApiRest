package services

import (
	"strconv"

	model "github.com/alejmendez/goApiRest/app/models"
	"github.com/alejmendez/goApiRest/app/utils"
	"github.com/alejmendez/goApiRest/core/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func ValidUser(id string, p string) bool {
	var user model.User
	database.DBConn.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !utils.CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

func ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func GetUserByEmail(e string) (*model.User, error) {
	var user model.User
	if err := database.DBConn.Where(&model.User{Email: e}).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(u string) (*model.User, error) {
	var user model.User
	if err := database.DBConn.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
