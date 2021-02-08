package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/ksrnnb/eikan/models"
	"github.com/ksrnnb/eikan/utils"
	"github.com/ksrnnb/eikan/validations"
)

// Register ユーザー登録
func Register(w http.ResponseWriter, r *http.Request) {
	// validation
	body, _ := ioutil.ReadAll(r.Body)
	err := validations.ValidateRegister(body)

	if err != nil {
		utils.ErrorLog("Input validation error", err)
		utils.ReturnValidationError(w)
		return
	}

	// userをdbに保存
	err = models.RegisterUser(body)
	if err != nil {
		utils.ErrorLog("User cannot be created", err)
		utils.ReturnInternalServerError(w)
		return
	}

	res, err := utils.NewMessageResponse("User has been created")
	if err != nil {
		utils.ErrorLog("Cannot marshal to json", err)
		utils.ReturnInternalServerError(w)
		return
	}

	utils.WriteResponse(w, res, http.StatusCreated)
}
