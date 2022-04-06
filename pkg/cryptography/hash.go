package cryptography

import "golang.org/x/crypto/bcrypt"

const PasswordCost = 14

// HashPassword generates a bcrypt hash of the password.
// Uses a work factor of PasswordCost.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)

	return string(bytes), err
}

// CheckPasswordHash compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns bool if the password match, false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
