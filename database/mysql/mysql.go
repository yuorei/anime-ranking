package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	conn *gorm.DB
}

func NewMySQL() (*MySQL, error) {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_TZ := os.Getenv("DB_TZ")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%s",DB_USER,DB_PASS,DB_PORT,DB_NAME,DB_TZ)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return &MySQL{
		conn: db,
	}, err
}
