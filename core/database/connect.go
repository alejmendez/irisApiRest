package database

import (
	"fmt"
	"log"

	"github.com/alejmendez/goApiRest/core/config"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

// ConnectDB connect to db
func ConnectDB() *gorm.DB {
	// log.Println(GetConnectionString())
	db, err := gorm.Open("postgres", GetConnectionString())
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connection Opened to Database")
	DB = db
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
