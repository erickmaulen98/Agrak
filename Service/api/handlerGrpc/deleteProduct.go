package handlergrpc

import (
	"context"
	pb "product_service/service"
)

// DeleteProduct deletes a product
func (productService *Product) DeleteProduct(ctx context.Context, productRequest *pb.ProductSKU) (*pb.DeleteResponse, error) {

	// Call the use case to delete the product
	deletedProducts, err := productService.useCase.DeleteProduct(productRequest.Sku)
	if err != nil {
		return &pb.DeleteResponse{
			Count: 0,
		}, err
	}

	return &pb.DeleteResponse{
		Count: deletedProducts,
	}, nil
}
