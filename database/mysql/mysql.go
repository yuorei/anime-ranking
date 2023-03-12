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

func NewMySQL() *MySQL {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_TZ := os.Getenv("DB_TZ")

	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=" + DB_TZ
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB接続失敗")
	}

	return &MySQL{
		Conn: db,
	}
}

func (d *MySQL) CreateTable() {
	if err := d.Conn.AutoMigrate(&table.User{}); err != nil {
		log.Fatalf("Database create table failed失敗")
	}
}
