package grpcCalls

import (
	pb "product_service_composite/service"
)

// Interface grpc calls
type ProductCalls interface {
	pb.ProductServiceClient
}

type GrpcCallService struct {
	ProductClient ProductCalls
}

// NewServiceProduct create a new grpc call service
func NewServiceProduct(productClient ProductCalls) *GrpcCallService {
	return &GrpcCallService{
		ProductClient: productClient,
	}
}
