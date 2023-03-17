package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/yuorei/anime-ranking/database/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Conn *gorm.DB
}

var db MySQL

func NewMySQL() {
	var err error

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_TZ := os.Getenv("DB_TZ")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%s", DB_USER, DB_PASS, DB_PORT, DB_NAME, DB_TZ)

	db.Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB接続失敗")
	}

}

func CreateTable() {
	if err := db.Conn.AutoMigrate(&table.User{}); err != nil {
		log.Fatalf("Database create table failed失敗")
	}
}

func InsertUser(user table.User) (table.User, error) {
	if err := db.Conn.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
