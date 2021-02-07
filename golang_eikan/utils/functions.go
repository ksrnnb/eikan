package utils

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CurrentTime returns current time based on RFC3339
func CurrentTime() time.Time {
	return time.Now()
}

// GenerateHashedPassword generates hashed password
func GenerateHashedPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash)
}
