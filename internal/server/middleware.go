package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func LogRequestInfoMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("request: %v", req)
		log.Printf("info: %v", info)

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			log.Printf("metadata: %v", md)
		}
		return handler(ctx, req)
	}
}
