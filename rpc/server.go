package main

import (
	"awesomeProject/rpc/order"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type OrderSerivceImpl struct {
}

func (os *OrderSerivceImpl) GetOrderInfo(ctx context.Context, request *order.OrderRequest) (*order.OrderInfo, error) {
	fmt.Println(request)
	return &order.OrderInfo{
		Id:          1,
		Name:        "abc",
		OrderStatus: "aaaaa",
	}, nil
}

func main() {
	server := grpc.NewServer()
	order.RegisterOrderServiceServer(server, &OrderSerivceImpl{})

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)
}
