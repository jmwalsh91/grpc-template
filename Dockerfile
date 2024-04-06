# Use the official Go image to create a build artifact.
FROM golang:1.16 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
COPY go.* ./
RUN go mod download

# Build the binary.
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server ./cmd/server

# Use the official Alpine image for a lean production container.
FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /server

# Service must listen to $PORT environment variable.
# This default value facilitates local development.
ENV GRPC_PORT 50051
EXPOSE $GRPC_PORT

# Run the web service on container startup.
CMD ["/server"]
