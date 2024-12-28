package server

import (
	"log"
	"net"

	"github.com/vignesh/grpc_demo/protogen/golang/orders"
	"github.com/vignesh/grpc_demo/service"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:50051"

func Init() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}

	server := grpc.NewServer()
	orderService := service.NewOrderService()
	orders.RegisterOrdersServer(server, &orderService)
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
