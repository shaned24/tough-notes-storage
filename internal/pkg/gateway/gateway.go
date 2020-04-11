package gateway

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"fmt"
	"github.com/shaned24/tough-notes-storage/api/notespb"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type HttpGateway struct {
	Port          string
	GRPCProxyHost string
	GRPCProxyPort string
}

func NewHttpGateway(port, grpcProxyHost, grpcProxyPort string) *HttpGateway {
	return &HttpGateway{
		Port:          port,
		GRPCProxyHost: grpcProxyHost,
		GRPCProxyPort: grpcProxyPort,
	}
}

func (s *HttpGateway) Start() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := notespb.RegisterNoteServiceHandlerFromEndpoint(
		ctx, mux,
		fmt.Sprintf("%s:%s", s.GRPCProxyHost, s.GRPCProxyPort),
		opts,
	)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint
	log.Println("Serving HttpProxy on", fmt.Sprintf(":%s", s.Port))
	return http.ListenAndServe(fmt.Sprintf(":%s", s.Port), mux)
}
