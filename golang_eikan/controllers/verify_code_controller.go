package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ksrnnb/eikan/models"
	"github.com/ksrnnb/eikan/utils"
	"github.com/ksrnnb/eikan/validations"
)

// VerifyCode ...
func VerifyCode(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	err := validations.ValidateVerifyCode(body)

	if err != nil {
		utils.ErrorLog("Input validation error", err)
		utils.ReturnValidationError(w)
		return
	}

	// レコードを認証、usersテーブルに保存。
	userToken, err := models.VerifyCode(body)
	if err != nil {
		utils.ErrorLog("No verify code found", err)
		utils.ReturnInternalServerError(w)
		return
	}

	responseMap := map[string]string{
		"message":   "Email has been verified",
		"userToken": userToken,
	}

	res, err := json.Marshal(responseMap)
	if err != nil {
		utils.ErrorLog("Cannot marshal to json", err)
		utils.ReturnInternalServerError(w)
		return
	}

	utils.WriteResponse(w, res, http.StatusCreated)
}
