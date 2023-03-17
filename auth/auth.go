package auth

import "github.com/yuorei/anime-ranking/database/table"

// Retrieve the user from the database
func GetUserByName(name string) (table.User, error) {
	return table.User{}, nil
}

// Check if the password is correct
func VerifyPassword(userPassword, inputPassword string) bool {
	return true
}

// Generate a JWT token
func GenerateToken(userID uint) (string, error) {
	return "", nil
}
