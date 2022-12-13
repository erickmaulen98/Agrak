package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"product_service/api/middleware"
	"product_service/config"
	"product_service/infrastructure/repository"
	"product_service/usecase/product"

	pb "product_service/service"

	handler "product_service/api/handlerGrpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	//User - Password - Host - DB
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&tlsInsecure=true", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_NAME)
	fmt.Println(mongoURI)

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Test pinging the database
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Connect client to Product Repository
	productRepo := repository.NewProductCollection(client)
	productUseCase := product.NewService(productRepo)

	// Configure the net listener to the api port defined
	port := config.API_PORT
	if port == "" {
		port = "8080"
	}

	grpcEndpoint := fmt.Sprintf(":%s", port)
	log.Printf("gRPC endpoint [%s]", grpcEndpoint)

	netListener, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Setting new Logger
	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	// Creates grpc server
	grpcServerGC := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.LoggingServerInterceptor),
	)
	log.Printf("Server listening at %v", netListener.Addr())

	productImplementation := handler.NewProductImplementation(productUseCase)

	pb.RegisterProductServiceServer(grpcServerGC, productImplementation)

	// Starting the server
	if err := grpcServerGC.Serve(netListener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
