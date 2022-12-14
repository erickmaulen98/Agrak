package httpcalls

import (
	"net/http"
	presenter "product_service_composite/api/presenter/product"
	"product_service_composite/entity"
	"product_service_composite/pkg"
)

func (httpService *HandlerProduct) UpdateProduct() http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		// Check body adapt to entity
		product, err := presenter.AdaptRequestToUpdate(request)
		if err != nil {
			entity.JSONResponse(responseWriter, err, http.StatusBadRequest)
			return
		}

		// Send request to usecase
		productUseCase, err := httpService.useCaseProduct.UpdateProduct(product)
		if err != nil {
			pkg.LogginHandleError(err)
			entity.JSONResponse(responseWriter, err, http.StatusBadRequest)
			return
		}

		entity.JSONResponse(responseWriter, productUseCase, http.StatusCreated)

	})
}
