package utils

import (
	"encoding/json"
	"net/http"
)

// ReturnInternalServerError returns error response
func ReturnInternalServerError(w http.ResponseWriter) {
	msg := Message{Message: "Internal Server Error"}
	res, _ := json.Marshal(msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(res)
}

// ReturnValidationError returns error response
func ReturnValidationError(w http.ResponseWriter) {
	msg := Message{Message: "Validation Error"}
	res, _ := json.Marshal(msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
}
