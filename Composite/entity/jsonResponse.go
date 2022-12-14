package entity

import (
	"encoding/json"
	"errors"
	"net/http"
)

// JSONResponse return a json response with a status code and a response body
func JSONResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	//Return if no response is intended to be sent, nil
	if statusCode == 204 {
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(errors.New("Error encoding JSON"))
		return
	}
}
