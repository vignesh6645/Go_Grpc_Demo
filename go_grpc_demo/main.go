package main

import (
	"fmt"
	"log"

	"github.com/vignesh/grpc_demo/protogen/golang/orders"
	"github.com/vignesh/grpc_demo/server"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	orderItem := orders.Order{
		OrderId:    10,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  timestamppb.Now(),
		Products: []*orders.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: orders.ProductType_DRINK},
		},
	}
	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatal("deserialization error:", err)
	}

	fmt.Println(string(bytes))
	server.Init()
}
