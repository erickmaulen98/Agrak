package handlergrpc

import (
	"context"
	"product_service/api/presenter"
	pb "product_service/service"
)

// CreateProduct creates a product
func (productService *Product) CreateProduct(ctx context.Context, productRequest *pb.Product) (*pb.Product, error) {
	// Call a presenter to convert the request to a product entity
	product := presenter.AdaptGrpcToEntity(productRequest)

	// Call the use case to create the product
	productCreated, err := productService.useCase.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	// Call a presenter to convert the product entity to a response
	productResponse := presenter.AdaptEntityToGrpc(productCreated)

	return productResponse, nil
}
