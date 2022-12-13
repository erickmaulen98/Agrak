package handlergrpc

import (
	"context"
	"testing"

	"product_service/entity"
	pb "product_service/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProduct_DeleteProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	productUseCase := NewMockUseCase(controller)
	productImplementation := NewProductImplementation(productUseCase)

	productSku := pb.ProductSKU{
		Sku: "123",
	}

	t.Run("DeleteProduct", func(t *testing.T) {
		productUseCase.EXPECT().DeleteProduct("123").Return(int64(1), nil)
		_, err := productImplementation.DeleteProduct(context.TODO(), &productSku)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("DeleteProductError", func(t *testing.T) {
		productUseCase.EXPECT().DeleteProduct("123").Return(int64(0), entity.ErrDeleteProduct)
		_, err := productImplementation.DeleteProduct(context.TODO(), &productSku)
		assert.Equal(t, err, entity.ErrDeleteProduct)
	})
}
