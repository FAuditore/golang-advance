package main

import (
	"advance/rpc/login/pb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	authServiceClient := pb.NewAuthServiceClient(conn)

	loginResponse, err := authServiceClient.Login(context.Background(), &pb.LoginRequest{
		Username: "liubo",
		Password: "1243",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%#v\n", loginResponse)
}
