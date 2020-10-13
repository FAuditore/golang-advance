package main

import (
	"awesomeProject/rpc/order"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn,err:=grpc.Dial(":8081",grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	orderServiceClient := order.NewOrderServiceClient(conn)

	orderInfo,err:=orderServiceClient.GetOrderInfo(context.Background(),&order.OrderRequest{
		OrderId:   1,
		TimeStamp: time.Now().Unix(),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(orderInfo.Id)
	fmt.Println(orderInfo.Name)
	fmt.Println(orderInfo.OrderStatus)
}
