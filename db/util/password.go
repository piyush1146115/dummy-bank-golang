package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the given password
func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash of the password: %w", err)
	}

	return string(hashedPass), nil
}

// CheckPassword verifies if the provided password is correct or not
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
