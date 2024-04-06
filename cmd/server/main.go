package main

import (
	"log"
	"os"

	"grpc-template/pkg/myservice"
)

func main() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}
	address := ":" + port
	log.Println("Starting gRPC server...")
	if err := myservice.StartGRPCServer(address); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
