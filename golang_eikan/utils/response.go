package utils

import (
	"net/http"
)

// WriteResponse creates json response
func WriteResponse(w http.ResponseWriter, jsonResponse []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}

// ReturnInternalServerError returns error response
func ReturnInternalServerError(w http.ResponseWriter) {
	res, _ := NewMessageResponse("Internal Server Error")

	WriteResponse(w, res, http.StatusInternalServerError)
}

// ReturnValidationError returns error response
func ReturnValidationError(w http.ResponseWriter) {
	res, _ := NewMessageResponse("Validation Error")

	WriteResponse(w, res, http.StatusBadRequest)
}
