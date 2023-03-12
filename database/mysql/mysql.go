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

	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=" + DB_TZ
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return &MySQL{
		conn: db,
	}, err
}
