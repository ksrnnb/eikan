package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ksrnnb/eikan/models"
	"github.com/ksrnnb/eikan/utils"
	"github.com/ksrnnb/eikan/validations"
)

// Home test用
// TODO:　削除
func Home(w http.ResponseWriter, r *http.Request) {
	message := utils.NewMessage("OK")
	res, err := json.Marshal(message)
	if err != nil {
		utils.ErrorLog("JSON ERROR: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// RegisterEmail メールアドレスの登録
func RegisterEmail(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	err := validations.ValidateEmail(body)

	if err != nil {
		utils.ErrorLog("Input validation error", err)
		utils.ReturnValidationError(w)
		return
	}

	// 認証用レコードを保存
	err = models.RegisterEmail(body)
	if err != nil {
		utils.ErrorLog("Verify code cannot be created", err)
		utils.ReturnInternalServerError(w)
		return
	}

	// 保存した後
	res, err := utils.NewMessageResponse("Verify code has been sent")
	if err != nil {
		utils.ErrorLog("Cannot marshal to json", err)
		utils.ReturnInternalServerError(w)
		return
	}

	utils.WriteResponse(w, res, http.StatusCreated)
}
