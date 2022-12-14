package presenter

import (
	"encoding/json"
	"net/http"
	"product_service_composite/entity"
)

func AdaptRequestToEntity(request *http.Request) (*entity.Product, error) {
	// Check body adapt to entity
	product := &entity.Product{}
	err := json.NewDecoder(request.Body).Decode(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
