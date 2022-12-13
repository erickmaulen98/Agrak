package handlergrpc

import (
	"context"
	"product_service/api/presenter"
	pb "product_service/service"
)

// GetProduct get a product
func (productService *Product) GetProduct(ctx context.Context, productRequest *pb.ProductSKU) (*pb.Product, error) {
	// Call the use case to get the product
	product, err := productService.useCase.GetProduct(productRequest.Sku)
	if err != nil {
		return &pb.Product{}, err
	}

	// Adapt the entity to the gRPC response
	return presenter.AdaptEntityToGrpc(product), nil
}
