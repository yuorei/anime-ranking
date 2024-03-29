package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yuorei/anime-ranking/database/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db.Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("DB接続失敗: " + err.Error())
	}

}

func CreateTable() {
	if err := db.Conn.AutoMigrate(&table.User{}); err != nil {
		log.Fatalf("Database create table failed")
	}

	if err := db.Conn.AutoMigrate(&table.AnimeRanking{}); err != nil {
		log.Fatalf("Database create table failed")
	}
}

func InsertUser(user table.User) (table.User, error) {
	if err := db.Conn.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByName(name string) (table.User, error) {
	var user table.User
	db.Conn.Where("name = ?", name).First(&user)
	return user, nil
}

func InsertAnimeRanking(anime table.AnimeRanking) (table.AnimeRanking, error) {
	if err := db.Conn.Create(&anime).Error; err != nil {
		return anime, err
	}

	return anime, nil
}

func GetAllUsers() ([]*table.User, error) {
	var users []*table.User
	if err := db.Conn.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetAnimeRankingByID(id int) (table.AnimeRanking, error) {
	var animeRanking table.AnimeRanking
	if err := db.Conn.First(&animeRanking, id).Error; err != nil {
		return table.AnimeRanking{}, err
	}
	return animeRanking, nil
}

func GetUserByID(id int) (table.User, error) {
	var user table.User
	if err := db.Conn.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetHaveAnimeByUserID(id int) ([]table.AnimeRanking, error) {
	var animes []table.AnimeRanking
	if err := db.Conn.Where("user_id = ?", id).Order("`rank` ASC").Find(&animes).Error; err != nil {
		return animes, err
	}
	return animes, nil
}

func UpdateAnimeRanking(id int, anime table.AnimeRanking) (table.AnimeRanking, error) {
	anime.UpdatedAt = time.Now()
	if err := db.Conn.Save(&anime).Error; err != nil {
		return anime, err
	}
	return anime, nil
}

func UpdateUser(user table.User) (table.User, error) {
	user.UpdatedAt = time.Now()
	if err := db.Conn.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DeleteAnimeRanking(anime table.AnimeRanking) error {
	if err := db.Conn.Delete(&anime).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user table.User) error {
	if err := db.Conn.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserAllAnimeRanking(id int) error {
	var anime table.AnimeRanking
	if err := db.Conn.Where("user_id = ?", id).Delete(&anime).Error; err != nil {
		return err
	}
	return nil
}
