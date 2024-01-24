package utils

import "golang.org/x/crypto/bcrypt"

// ? This function is used to convert the simple string password into the HashPassword
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
