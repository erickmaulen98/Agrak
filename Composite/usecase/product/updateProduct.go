package productUsecase

import "product_service_composite/entity"

func (s *Service) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	// Send request to grpc handler
	return s.GrpcHandler.UpdateProduct(product)
}
