package validations

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/ksrnnb/eikan/utils"
)

// VerifyCode struct のvalidateタグがバリデーションルール
type VerifyCode struct {
	Email string `json:"email" validate:"required,email,lte=255"`
	Code  string `json:"code" validate:"required,is-code"`
}

// ValidateVerifyCode validates input
func ValidateVerifyCode(body []byte) error {
	validate := validator.New()

	err := validate.RegisterValidation("is-code", utils.IsCode)

	if err != nil {
		return err
	}
	var myValidator VerifyCode
	json.Unmarshal(body, &myValidator)

	err = validate.Struct(myValidator)

	return err
}
