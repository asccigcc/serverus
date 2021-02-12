package serverus

import (
	"context"
	"log"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type Serverus struct {
	port         string
	lis          net.Listener
	server       *grpc.Server
	interceptors []grpc.UnaryServerInterceptor
}

// Move to interceptors
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	h, err := handler(ctx, req)

	log.Printf(
		"request - Method:%s\tDuration:%s\tError:%v\t",
		info.FullMethod,
		time.Since(start),
		err,
	)

	return h, err
}

// Serverus receivers

// Chain add interceptor to the interceptors
// array so we can add more and pass into the
// grpc server when this is instantiated.
func (s *Serverus) ChainInterceptors(inter interface{}) {}

// Initialize gRPC server passing the interceptors
func (s *Serverus) InitGRPC() {
	log.Println("Initializing gRPC")
	s.server = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(s.interceptors...)), // register the interceptors injected and the base
	)
}

type RegisterHandler interface {
	RegisterHandlerServer(s grpc.ServiceRegistrar) bool
}

// Register the Protobuf server handler
func (s *Serverus) RegisterHandler(handler RegisterHandler) {
	handler.RegisterHandlerServer(s.server)
}

// Start the gRPC server
func (s *Serverus) StartServerus() {
	if err := s.server.Serve(s.lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

// Initialize new gRPC serverus server
func NewServerus(port string, interceptors ...grpc.UnaryServerInterceptor) *Serverus {
	port = getPort(port)

	log.Println("Creating new serverus")
	return &Serverus{
		port: port,
		lis:  getListener(port),
	}
}

func getListener(port string) net.Listener {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Fialed to listen: %v", err)
	}

	return lis
}

// @TODO move to helpers
func getPort(port string) string {
	if len(port) > 0 {
		return port
	}

	return ":3000"
}
