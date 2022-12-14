package grpcCalls

import (
	"context"
	presenter "product_service_composite/api/presenter/product"
	"product_service_composite/entity"
	pb "product_service_composite/service"
	"time"
)

// GetAllProduct get all products from gRPC service
func (service *GrpcCallService) GetAllProduct() ([]*entity.Product, error) {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call gRPC service
	productsGrpc, err := service.ProductClient.ListProducts(ctx, &pb.ProductListRequest{})
	if err != nil {
		return nil, err
	}

	// Call presenter to adapt gRPC response to entity
	productsEntity, err := presenter.AdaptGrpcToEntityArray(productsGrpc)
	if err != nil {
		return nil, err
	}

	return productsEntity, nil
}
