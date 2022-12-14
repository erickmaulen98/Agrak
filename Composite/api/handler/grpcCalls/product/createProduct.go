package grpcCalls

import (
	"context"
	presenter "product_service_composite/api/presenter/product"
	"product_service_composite/entity"
	"product_service_composite/pkg"
	"time"
)

// CreateProduct create a product
func (service *GrpcCallService) CreateProduct(product *entity.Product) (*entity.Product, error) {

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Call presenter
	productAdapter := presenter.AdaptEntityToGrpc(product)

	// Call gRPC service
	productGrpc, err := service.ProductClient.CreateProduct(ctx, productAdapter)
	if err != nil {
		pkg.LogginHandleError(err)
		return nil, err
	}

	// Call presenter to adapt gRPC response to entity
	productEntity := presenter.AdaptGrpcToEntity(productGrpc)

	return productEntity, nil
}
