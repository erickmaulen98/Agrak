package grpcCalls

import (
	"context"
	presenter "product_service_composite/api/presenter/product"
	"product_service_composite/pkg"
	"time"
)

// DeleteProduct delete a product
func (service *GrpcCallService) DeleteProduct(sku string) (int64, error) {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call presenter to adapt string to grpc
	skuGrpc := presenter.AdaptStringToGrpc(sku)

	// Call gRPC service
	countDelete, err := service.ProductClient.DeleteProduct(ctx, &skuGrpc)
	if err != nil {
		pkg.LogginHandleError(err)
		return 0, err
	}

	// Call presenter to adapt gRPC response to entity
	countDeleteEntity := presenter.AdaptGrpcToEntityInt64(countDelete)

	return countDeleteEntity, nil
}
