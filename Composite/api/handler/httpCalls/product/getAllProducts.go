package httpcalls

import (
	"net/http"
	"product_service_composite/entity"
	"product_service_composite/pkg"
)

func (httpService *HandlerProduct) GetAllProduct() http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		// Send request to usecase
		products, err := httpService.useCaseProduct.GetAllProduct()
		if err != nil {
			pkg.LogginHandleError(err)
			entity.JSONResponse(responseWriter, err, http.StatusBadRequest)
			return
		}

		entity.JSONResponse(responseWriter, products, http.StatusOK)

	})
}
