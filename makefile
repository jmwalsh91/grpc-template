.PHONY: proto run docker-build docker-run

GRPC_PORT ?= 50051

# Compile the protobuf files
proto:
	protoc --go_out=plugins=grpc:pkg/myservice api/v1/*.proto

# Run the server using AIR for hot reloading
run:
	@echo "Reading environment variables from .env file..."
	@export $$(grep -v '^#' .env | xargs) && air -c .air.toml

# Build the Docker container
docker-build:
	docker build -t my-grpc-service .

# Run the Docker container
# make docker-run GRPC_PORT=50052
docker-run:
	docker run -p $(GRPC_PORT):$(GRPC_PORT) -e GRPC_PORT=$(GRPC_PORT) my-grpc-service
