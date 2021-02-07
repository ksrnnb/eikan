package utils

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// IsDate checks argument is valid date or not
func IsDate(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	_, err := time.Parse(time.RFC3339[:len(date)], date)

	return err == nil
}
