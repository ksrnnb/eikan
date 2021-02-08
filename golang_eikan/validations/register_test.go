package validations

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestValidPattern(t *testing.T) {
	pattern := Register{
		Password:   "password",
		Birthday:   "1999-11-11",
		GenderType: 1,
	}

	data, _ := json.Marshal(pattern)

	if ValidateRegister(data) != nil {
		t.Errorf("Pattern is valid but ValidateRegister function was incorrect.")
	}
}

func TestTooLongStr(t *testing.T) {
	tooLongStr := strings.Repeat("a", 256)

	pattern := Register{
		Password:   tooLongStr,
		Birthday:   "1999-11-11",
		GenderType: 1,
	}

	data, _ := json.Marshal(pattern)

	if ValidateRegister(data) == nil {
		t.Errorf("password is too long but validation function didn't validate")
	}
}

func TestInvalidDate(t *testing.T) {
	pattern := Register{
		Password:   "password",
		Birthday:   "1999-02-29",
		GenderType: 1,
	}

	data, _ := json.Marshal(pattern)

	if ValidateRegister(data) == nil {
		t.Errorf("birthday is incorrect but validation function didn't validate")
	}
}

func TestInvalidGenderType(t *testing.T) {
	pattern := Register{
		Password:   "password",
		Birthday:   "1999-11-11",
		GenderType: 10,
	}

	data, _ := json.Marshal(pattern)
	if ValidateRegister(data) == nil {
		t.Errorf("GenderType is incorrect but validation function didn't validate")
	}
}
