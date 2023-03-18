package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/yuorei/anime-ranking/database/mysql"
	"github.com/yuorei/anime-ranking/database/table"
)

// Retrieve the user from the database
func GetUserByName(name string) (table.User, error) {
	user, err := mysql.GetUserByName(name)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Check if the password is correct
func VerifyPassword(userPassword, inputPassword string) bool {
	if userPassword == inputPassword {
		return true
	}
	return false
}

// Generate a JWT token
func GenerateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = strconv.Itoa(int(userID))
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	// TODO secret-keyをハードコーディングしているので環境変数からにする
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}
	return tokenString,nil
}
