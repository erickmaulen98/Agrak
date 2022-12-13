package product

import "product_service/entity"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Create a new product
func (serviceInstance *Service) CreateProduct(product *entity.Product) (*entity.Product, error) {
	return serviceInstance.repo.CreateProduct(product)
}

// Update a product
func (serviceInstance *Service) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	return serviceInstance.repo.UpdateProduct(product)
}

// Delete a product
func (serviceInstance *Service) DeleteProduct(sku string) (int64, error) {
	return serviceInstance.repo.DeleteProduct(sku)
}

// Get a product
func (serviceInstance *Service) GetProduct(sku string) (*entity.Product, error) {
	return serviceInstance.repo.GetProduct(sku)
}

// Get all products
func (serviceInstance *Service) GetProducts() ([]*entity.Product, error) {
	return serviceInstance.repo.GetProducts()
}
