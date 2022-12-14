package productUsecase

import "product_service_composite/entity"

func (s *Service) GetAllProduct() ([]*entity.Product, error) {
	// Send request to grpc handler
	return s.GrpcHandler.GetAllProduct()
}
