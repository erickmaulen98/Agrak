package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"product_service_composite/config"

	"product_service_composite/api/handler"
	handlerOut "product_service_composite/api/handler/grpcCalls/product"
	handlerIn "product_service_composite/api/handler/httpCalls/product"
	pb "product_service_composite/service"
	productUsecase "product_service_composite/usecase/product"

	ctx "github.com/gorilla/context"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
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

	// Instance GrpcHandlerService
	grpcProduct := handlerOut.NewServiceProduct(productClient)

	// Instance Use Case
	productUseCase := productUsecase.NewService(grpcProduct)

	// Instance Handler IN
	productHandlerOut := handlerIn.NewHandlerProductImplementation(*productUseCase)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandler(r, *n, productHandlerOut)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	port := fmt.Sprintf(":%s", config.API_PORT)
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         port,
		Handler:      ctx.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	log.Printf("Server listening at %v", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
