package handlergrpc

import (
	"context"
	pb "product_service/service"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProduct_ListProducts(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	productUseCase := NewMockUseCase(controller)
	productImplementation := NewProductImplementation(productUseCase)

	products := pb.ProductList{
		Products: []*pb.Product{
			{
				Sku:   "123",
				Name:  "Product 1",
				Brand: "Brand 1",
				Price: 100,
				Size:  "m",
			},
			{
				Sku:   "456",
				Name:  "Product 2",
				Brand: "Brand 2",
				Price: 200,
				Size:  "l",
			},
		},
	}

	t.Run("ListProducts", func(t *testing.T) {
		productUseCase.EXPECT().GetProducts().Return(products, nil)
		_, err := productImplementation.ListProducts(context.TODO(), &pb.ProductListRequest{})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
