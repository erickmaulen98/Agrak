package productUsecase

type Service struct {
	GrpcHandler GrpcCalls
}

//NewService create new use case service
func NewService(grpcHandler GrpcCalls) *Service {
	return &Service{
		GrpcHandler: grpcHandler,
	}
}
