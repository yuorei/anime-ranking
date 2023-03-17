package auth

import (
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
	if userPassword==inputPassword{
		return true
	}
	return false
}

// Generate a JWT token
func GenerateToken(userID uint) (string, error) {
	return "", nil
}
