package main

import (
	"net/http"

	"github.com/ksrnnb/eikan/utils"

	"github.com/gorilla/mux"
	"github.com/ksrnnb/eikan/controllers"
	"github.com/ksrnnb/eikan/db"
)

var host string
var version string
var port string

func init() {
	db.Connect()
	host = "0.0.0.0"
	version = "/v1"
	port = ":8000"

	utils.LoggingSetUp("/var/log/app.log")
}

func main() {

	r := mux.NewRouter()
	s := r.PathPrefix(version).Subrouter()
	s.HandleFunc("", controllers.Home)
	s.HandleFunc("/email", controllers.RegisterEmail)
	s.HandleFunc("/code", controllers.VerifyCode)
	s.HandleFunc("/register", controllers.Register)
	http.Handle("/", s)
	http.ListenAndServe(port, nil)

	db.Close()
}
