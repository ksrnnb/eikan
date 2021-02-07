package main

import (
	"net/http"

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
}

// TODO: logのセットアップ
func main() {

	r := mux.NewRouter()
	s := r.PathPrefix(version).Subrouter()
	s.HandleFunc("", controllers.Home)
	s.HandleFunc("/register", controllers.Register)
	http.Handle("/", s)
	http.ListenAndServe(port, nil)

	db.Close()
}
