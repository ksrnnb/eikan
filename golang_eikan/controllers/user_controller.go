package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ksrnnb/eikan/models"
	"github.com/ksrnnb/eikan/utils"
	"github.com/ksrnnb/eikan/validations"
)

// test用
func Home(w http.ResponseWriter, r *http.Request) {
	message := utils.NewMessage("OK")
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("JSON ERROR: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ユーザー登録
func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: もう少しModelに処理をもっていく
	body, _ := ioutil.ReadAll(r.Body)

	err := validations.ValidateRegister(body)
	if err != nil {
		log.Printf("Input validation error: %v\n", err)
		utils.ReturnValidationError(w)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		log.Printf("JSON unmarshal error: %v\n", err)

		utils.ReturnValidationError(w)
		return
	}

	// userをdbに保存
	user.Create()

	// 成功したあと
	message := utils.NewMessage("user has created.")
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Cannot create message: %v", err)
		utils.ReturnInternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
