package handlergrpc

import (
	"context"
	"product_service/entity"
	pb "product_service/service"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProduct_GetProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	productUseCase := NewMockUseCase(controller)
	productImplementation := NewProductImplementation(productUseCase)
	product := entity.Product{
		ID:             "1",
		Name:           "Product 1",
		Brand:          "Brand 1",
		Price:          100.23,
		Size:           "M",
		Images:         []string{"image1", "image2"},
		PrincipalImage: "image1",
	}

	sku := pb.ProductSKU{
		Sku: "123",
	}
	t.Run("GetProduct", func(t *testing.T) {
		productUseCase.EXPECT().GetProduct("123").Return(&product, nil)
		_, err := productImplementation.GetProduct(context.TODO(), &sku)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
