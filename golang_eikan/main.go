package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ksrnnb/eikan/db"
	"github.com/ksrnnb/eikan/models"
	"github.com/ksrnnb/eikan/utils"
)

var host string
var version string
var port string

func init() {
	db.Connect()
	host = "0.0.0.0"
	version = "/v1"
	port = ":8000"
}

func home(w http.ResponseWriter, r *http.Request) {
	message := utils.NewMessage("OK")
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("JSON ERROR: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func register(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var user models.User
	err := json.Unmarshal(body, &user)

	if err != nil {
		log.Printf("JSON unmarshal error: %v\n", err)
		// bad request???
		// あとで
		// 40x系のエラーを返すように。
		return
	}

	// TODO: validation check

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

func main() {

	r := mux.NewRouter()
	s := r.PathPrefix(version).Subrouter()
	s.HandleFunc("", home)
	s.HandleFunc("/register", register)
	http.Handle("/", s)
	http.ListenAndServe(port, nil)

	db.Close()
}
