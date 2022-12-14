package productUsecase

func (s *Service) DeleteProduct(sku string) (int64, error) {
	// Send request to grpc handler
	return s.GrpcHandler.DeleteProduct(sku)
}
