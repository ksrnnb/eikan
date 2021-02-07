package utils

import "golang.org/x/crypto/bcrypt"

// Message model
type Message struct {
	Message string `json:"message"`
}

// NewMessage returns new Message instance
func NewMessage(message string) *Message {
	return &Message{Message: message}
}

// GenerateHashedPassword generates hashed password
func GenerateHashedPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash)
}
