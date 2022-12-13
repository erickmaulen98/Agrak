package handlergrpc

import (
	"context"
	"product_service/api/presenter"
	pb "product_service/service"
)

// UpdateProduct updates a product
func (productService *Product) UpdateProduct(ctx context.Context, productRequest *pb.Product) (*pb.Product, error) {

	// Call presenter
	productUpdate := presenter.AdaptGrpcToEntity(productRequest)

	// Call the use case to update the product
	product, err := productService.useCase.UpdateProduct(productUpdate)
	if err != nil {
		return &pb.Product{}, err
	}

	// Adapt the entity to the gRPC response
	return presenter.AdaptEntityToGrpc(product), nil
}
