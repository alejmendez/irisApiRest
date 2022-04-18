package database

import (
	"fmt"

	"github.com/alejmendez/goApiRest/core/config"
	"github.com/jinzhu/gorm"
)

// ConnectDB connect to db
func ConnectDB() *gorm.DB {
	var err error

	fmt.Println(GetConnectionString())
	DB, err = gorm.Open("postgres", GetConnectionString())
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	return DB
}

func GetConnectionString() string {
	host := config.Get("DB_HOST")
	port := config.Get("DB_PORT")
	user := config.Get("DB_USERNAME")
	password := config.Get("DB_PASSWORD")
	dbDatabase := config.Get("DB_DATABASE")

	template := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

	return fmt.Sprintf(template, host, port, user, password, dbDatabase)
}
