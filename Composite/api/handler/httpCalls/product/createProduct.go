package httpcalls

import (
	"net/http"
	presenter "product_service_composite/api/presenter/product"
	"product_service_composite/entity"
	"product_service_composite/pkg"
)

// CreateProduct  create a product
func (httpService *HandlerProduct) CreateProduct() http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		// Check body adapt to entity
		product, err := presenter.AdaptRequestToEntity(request)
		if err != nil {
			entity.JSONResponse(responseWriter, err, http.StatusBadRequest)
			return
		}

		// Send request to usecase
		productUseCase, err := httpService.useCaseProduct.CreateProduct(product)
		if err != nil {
			pkg.LogginHandleError(err)
			entity.JSONResponse(responseWriter, err, http.StatusBadRequest)
			return
		}

		entity.JSONResponse(responseWriter, productUseCase, http.StatusCreated)

	})
}
