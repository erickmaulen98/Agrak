package presenter

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// AdaptRequestToString adapt request to string for get sku
func AdaptRequestToString(request *http.Request) (string, error) {
	sku := mux.Vars(request)["sku"]
	if sku == "" {
		return "", errors.New("sku is required")
	}
	return sku, nil
}
