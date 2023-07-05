package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connection database
func DatabaseInit() {
	var err error

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	// // if you want use postgresql u can use it
	// dsn := fmt.Sprintf("host=%s user=%s  dbname=%s port=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT)
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// fmt.Println("hello world")

	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database")
}
