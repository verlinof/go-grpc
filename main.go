package main

import (
	"fmt"
	"log"

	pb "github.com/verlinof/go-grpc/protobuf/pb"
	"google.golang.org/protobuf/proto"
)

func main() {

	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    2,
				Name:  "Nike Air",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "Shoes",
				},
			},
		},
	}

	// Marshal / Encode data dari array
	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Compact binary wire format
	fmt.Println(data)

	// Decode dari marshal
	testProduct := &pb.Products{}
	if err = proto.Unmarshal(data, testProduct); err != nil {
		log.Fatal(err.Error())
	}

	for _, product := range testProduct.GetData() {
		fmt.Println(product.GetName())
		fmt.Println(product.GetCategory())
	}
}
