package services

import (
	productPb "github.com/verlinof/go-grpc/pb/products"
)

type ProductService struct {
	//Interface dari gRPC
	productPb.UnimplementedProductServiceServer
}
