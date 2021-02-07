package utils

import (
	"encoding/json"
)

// Message model
type Message struct {
	Message string `json:"message"`
}

// NewMessage returns new Message instance
func NewMessage(message string) *Message {
	return &Message{Message: message}
}

// NewResponse create new resonse with message
func NewResponse(message string) ([]byte, error) {
	msg := NewMessage(message)
	res, err := json.Marshal(msg)

	return res, err
}
