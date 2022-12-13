package main

import (
	"log"

	"product_service_composite/config"

	pb "product_service_composite/service"

	"google.golang.org/grpc"
)

func main() {

	opts := []grpc.DialOption{grpc.WithInsecure()}

	productConnection, err := grpc.Dial(config.PRODUCT_CLIENT_ADDRESS, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer productConnection.Close()

	// Set a Product Service with its gRPC client attached
	productClient := pb.NewProductServiceClient(productConnection)

}
