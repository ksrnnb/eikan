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

// NewMessageResponse create new resonse with message
func NewMessageResponse(message string) ([]byte, error) {
	msg := NewMessage(message)
	res, err := json.Marshal(msg)

	return res, err
}
