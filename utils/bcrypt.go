package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(password string, input_password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(input_password)) == nil
}
