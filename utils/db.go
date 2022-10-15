package utils

import (
	"UserManagementAPI/static"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	WriteLog("AppLog.txt")
	log.Println("=============================================")
	// dsn := "postgres://postgres:postgrespw@localhost:49153" // docker
	dsn := "postgres://" + static.DB_ACCOUNT + ":" + static.DB_PASSWORD + "@" + static.DB_HOST + ":" + static.DB_PORT +
		"/" + static.DB_NAME + "?sslmode=disable" // psql
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Panic("Init failed: Database connection.")
	} else {
		log.Println("Init successful: Database Connection.")
	}
	return db
}
