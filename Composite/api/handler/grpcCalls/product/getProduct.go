package grpcCalls

import (
	"context"
	presenter "product_service_composite/api/presenter/product"
	"product_service_composite/entity"
	"product_service_composite/pkg"
	"time"
)

// GetProduc get a product from gRPC service
func (service *GrpcCallService) GetProduc(sku string) (*entity.Product, error) {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call presenter to adapt string to grpc
	skuGrpc := presenter.AdaptStringToGrpc(sku)

	// Call gRPC service
	productGrpc, err := service.ProductClient.GetProduct(ctx, &skuGrpc)
	if err != nil {
		pkg.LogginHandleError(err)
		return nil, err
	}

	// Call presenter to adapt gRPC response to entity
	productEntity := presenter.AdaptGrpcToEntity(productGrpc)

	return productEntity, nil
}
