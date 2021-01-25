package main

import (
	"fmt"
	"net/http"

	"github.com/ksrnnb/eikan/configs"
	"github.com/ksrnnb/eikan/db"
)

var addr string
var port string

func init() {
	port = ":8000"
	cfg := configs.New()
	fmt.Println(cfg)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}

func main() {
	db.MongoInit()
	http.HandleFunc("/", hello)
	http.ListenAndServe(port, nil)
}
