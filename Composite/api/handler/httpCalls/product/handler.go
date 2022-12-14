package httpcalls

import (
	useCaseProduct "product_service_composite/usecase/product"
)

type HandlerProduct struct {
	useCaseProduct useCaseProduct.Service
}

func NewHandlerProductImplementation(useCaseProduct useCaseProduct.Service) *HandlerProduct {
	return &HandlerProduct{
		useCaseProduct: useCaseProduct,
	}
}
