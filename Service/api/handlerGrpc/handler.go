package handlergrpc

import (
	useCase "product_service/usecase/product"
)

type Product struct {
	useCase useCase.UseCase
}

func NewProductImplementation(useCase useCase.UseCase) *Product {
	return &Product{
		useCase: useCase,
	}
}
