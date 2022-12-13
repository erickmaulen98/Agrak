package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func LoggingServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}
