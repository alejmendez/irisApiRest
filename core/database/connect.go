package database

import (
	"fmt"

	"github.com/alejmendez/goApiRest/core/config"
	"github.com/jinzhu/gorm"
)

var (
	DBConn *gorm.DB
)

// ConnectDB connect to db
func ConnectDB() *gorm.DB {
	var err error

	// fmt.Println(GetConnectionString())
	db, err := gorm.Open("postgres", GetConnectionString())
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DBConn = db
	return db
}

func GetConnectionString() string {
	host := config.Conf.Db.Host
	port := config.Conf.Db.Port
	user := config.Conf.Db.Username
	password := config.Conf.Db.Password
	dbDatabase := config.Conf.Db.Database

	template := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

	return fmt.Sprintf(template, host, port, user, password, dbDatabase)
}
