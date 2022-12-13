package handlergrpc

import (
	"context"
	"product_service/api/presenter"
	pb "product_service/service"
)

// ListProducts get all products
func (productService *Product) ListProducts(ctx context.Context, productRequest *pb.ProductListRequest) (*pb.ProductList, error) {
	// Call the use case to get the products
	products, err := productService.useCase.GetProducts()
	if err != nil {
		return &pb.ProductList{}, err
	}

	// Adapt the entity to the gRPC response

	productsGrpc := presenter.AdaptEntityArrayToGrpc(products)

	return productsGrpc, nil
}
