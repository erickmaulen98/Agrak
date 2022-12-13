package product

import "product_service/entity"

type Reader interface {
	GetProduct(sku string) (*entity.Product, error)
	GetProducts() ([]*entity.Product, error)
}

type Writer interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(sku string) (int64, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(sku string) (int64, error)
	GetProduct(sku string) (*entity.Product, error)
	GetProducts() ([]*entity.Product, error)
}
