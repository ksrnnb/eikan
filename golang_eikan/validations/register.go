package validations

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/ksrnnb/eikan/utils"
)

// Register struct のvalidateタグがバリデーションルール
type Register struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,lte=255"`
	Birthday   string `json:"birthday" validate:"required,is-date"`
	GenderType int    `json:"genderType" validate:"required,lte=5"`
}

// ValidateRegister validates post parameter
func ValidateRegister(body []byte) error {
	validate := validator.New()
	err := validate.RegisterValidation("is-date", utils.IsDate)

	if err != nil {
		return err
	}

	var myValidator Register
	json.Unmarshal(body, &myValidator)

	err = validate.Struct(myValidator)

	return err
}
