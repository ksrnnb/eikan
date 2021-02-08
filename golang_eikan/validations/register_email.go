package validations

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// RegisterEmail struct のvalidateタグがバリデーションルール
type RegisterEmail struct {
	Email string `json:"email" validate:"required,email,lte=255"`
}

// ValidateEmail validates post parameter
func ValidateEmail(body []byte) error {
	validate := validator.New()

	var myValidator RegisterEmail
	json.Unmarshal(body, &myValidator)

	err := validate.Struct(myValidator)

	return err
}
