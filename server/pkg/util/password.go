package util

import (
	"golang.org/x/crypto/bcrypt"
)

const passwordCost = 14

// HashPassword hashes given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)

	return string(bytes), err
}

// CheckPasswordHash compares raw password with it's hashed value.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
