package productUsecase

import "product_service_composite/entity"

type GrpcCalls interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(sku string) (int64, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	GetProduc(sku string) (*entity.Product, error)
	GetAllProduct() ([]*entity.Product, error)
}
