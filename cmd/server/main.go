package main

import (
	"log"

	"grpc-template/pkg/myservice"
)

func main() {
	if err := myservice.StartGRPCServer(":50051"); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
