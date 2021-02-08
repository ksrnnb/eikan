package utils

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// IsDate checks argument is valid date or not
func IsDate(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	_, err := time.Parse(time.RFC3339[:len(date)], date)

	return err == nil
}

// IsCode checks its arguments is 6 digit nubmer or not
func IsCode(fl validator.FieldLevel) bool {
	code := fl.Field().String()
	matched, _ := regexp.Match(`\d{6}`, []byte(code))

	return matched
}
