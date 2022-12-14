package presenter

import (
	"net/http"
	"product_service_composite/entity"
)

func AdaptRequestToUpdate(request *http.Request) (*entity.Product, error) {
	// Check body adapt to entity
	product, err := AdaptRequestToEntity(request)
	if err != nil {
		return nil, err
	}

	sku, err := AdaptRequestToString(request)
	if err != nil {
		return nil, err
	}

	// Set sku from url
	product.Sku = sku
	return product, nil
}
