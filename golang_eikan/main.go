package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ksrnnb/eikan/db"
)

var addr string
var port string

type Message struct {
	Message string `json:"message"`
}

func init() {
	port = ":8000"
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}

func home(w http.ResponseWriter, r *http.Request) {
	message := Message{Message: "OK"}
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("JSON ERROR: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	db.Connect()

	http.HandleFunc("/", hello)
	http.HandleFunc("/v1", home)
	http.ListenAndServe(port, nil)

}
