package utils

// Message model
type Message struct {
	Message string `json:"message"`
}

// NewMessage returns new Message instance
func NewMessage(message string) *Message {
	return &Message{Message: message}
}
