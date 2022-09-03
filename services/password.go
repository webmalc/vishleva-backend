package services

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the provided password.
func HashPassword(pwd []byte) ([]byte, error) {
	if len(pwd) == 0 {
		return []byte{}, errors.New("the password is empty")
	}

	return bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
}

// CheckPassword is a simple utility function to check the password given as raw
// against the user's hashed password.
func CheckPassword(hash []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(password)) == nil
}
