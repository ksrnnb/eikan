package utils

import (
	"fmt"
	"math/rand"
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

// GenerateVerifyCode makes 6 digit number
func GenerateVerifyCode() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(1000000)
	code := fmt.Sprintf("%06d", random)

	return code
}

// GenerateToken makes 32 digit random string
func GenerateToken() string {
	token := randomString(32)
	return token
}

func randomString(digit int) string {
	rand.Seed(time.Now().UnixNano())
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var result string
	for i := 0; i < digit; i++ {
		n := rand.Intn(len(str))
		result = result + string(str[n])
	}

	return result
}
