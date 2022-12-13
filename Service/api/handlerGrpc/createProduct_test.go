package handlergrpc

import (
	"context"
	"product_service/entity"
	pb "product_service/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProduct_CreateProduct(t *testing.T) {
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

	productGrpc := pb.Product{
		Id:             "1",
		Name:           "Product 1",
		Brand:          "Brand 1",
		Price:          100.23,
		Size:           "M",
		Images:         []string{"image1", "image2"},
		PrincipalImage: "image1",
	}

	t.Run("CreateProduct", func(t *testing.T) {
		productUseCase.EXPECT().CreateProduct(gomock.Any()).Return(&product, nil)
		_, err := productImplementation.CreateProduct(context.TODO(), &productGrpc)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})

	t.Run("CreateProductError", func(t *testing.T) {
		productUseCase.EXPECT().CreateProduct(gomock.Any()).Return(nil, entity.ErrProductNotFound)
		_, err := productImplementation.CreateProduct(context.TODO(), &productGrpc)

		assert.Equal(t, entity.ErrProductNotFound, err)
	})
}
