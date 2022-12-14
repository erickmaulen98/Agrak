package httpcalls

import (
	"net/http"
	presenter "product_service_composite/api/presenter/product"
	"product_service_composite/entity"
	"product_service_composite/pkg"
)

func (httpService *HandlerProduct) DeleteProduct() http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		// Check sku by vars in url
		sku, err := presenter.AdaptRequestToString(request)
		if err != nil {
			pkg.LogginHandleError(err)
			entity.JSONResponse(responseWriter, err, http.StatusBadRequest)
			return
		}

		// Send request to usecase
		countDelete, err := httpService.useCaseProduct.DeleteProduct(sku)
		if err != nil {
			pkg.LogginHandleError(err)
			entity.JSONResponse(responseWriter, err, http.StatusBadRequest)
			return
		}

		entity.JSONResponse(responseWriter, countDelete, http.StatusOK)

	})
}
