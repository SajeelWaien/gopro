package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

func InitDB() *gorm.DB {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	// fmt.Print(os.Getenv("HOST"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DBCon = database

	if err != nil {
		panic("Could not connect to database")
	}

	return database
}
