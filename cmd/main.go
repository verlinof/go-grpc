package main

import (
	"log"
	"net"

	"github.com/verlinof/go-grpc/cmd/services"
	productPb "github.com/verlinof/go-grpc/pb/products"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {

	// Create a listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Create a gRPC server
	serv := grpc.NewServer()

	// Register Service
	productService := services.ProductService{}
	productPb.RegisterProductServiceServer(serv, &productService)

	log.Printf("Server started on port %v", lis.Addr())
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err.Error())
	}
}
