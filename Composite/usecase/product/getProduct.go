package productUsecase

import "product_service_composite/entity"

func (s *Service) GetProduct(sku string) (*entity.Product, error) {
	// Send request to grpc handler
	return s.GrpcHandler.GetProduc(sku)
}
