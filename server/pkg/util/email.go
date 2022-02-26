package util

import (
	"net/mail"
)

// ValidEmail checks if email address is valid.
func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}
