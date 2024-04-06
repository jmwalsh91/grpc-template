package myservice

import (
	"context"
	"fmt"
	"log"
	"net"

	myservicepb "grpc-template/api/v1"

	"google.golang.org/grpc"
)

type server struct {
	myservicepb.UnimplementedGreetingServiceServer
}

func (s *server) SayHello(ctx context.Context, in *myservicepb.HelloRequest) (*myservicepb.HelloResponse, error) {
	return &myservicepb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func StartGRPCServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myservicepb.RegisterGreetingServiceServer(s, &server{})

	log.Printf("gRPC server started on %s", address)

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
