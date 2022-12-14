package productUsecase

import "product_service_composite/entity"

func (s *Service) CreateProduct(product *entity.Product) (*entity.Product, error) {
	// Send request to grpc handler
	return s.GrpcHandler.CreateProduct(product)
}
