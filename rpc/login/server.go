package main

import (
	"advance/rpc/login/pb"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
)

var _ pb.AuthServiceServer = &AuthService{}

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func (srv *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for k, v := range md {
			fmt.Println(k, v)
		}
	}
	log.Printf("%#v\n", req)
	if req.Username == "liubo" && req.Password == "123" {
		return &pb.LoginResponse{AccessToken: "abcdefg"}, nil
	}
	return nil, status.Error(codes.NotFound, "帐号或密码错误")
}

func main() {
	go func() {
		listener, err := net.Listen("tcp", ":8081")
		if err != nil {
			panic(err)
		}
		mux := runtime.NewServeMux()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var opt []grpc.DialOption
		opt = append(opt, grpc.WithInsecure())

		//直接注册rest服务
		//pb.RegisterAuthServiceHandlerServer(ctx, mux, &AuthService{})

		//grpc-gateway rest代理到grpcServer
		err = pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, ":8080", opt)
		if err != nil {
			panic(err)
		}

		log.Printf("Start REST server at %s\n", listener.Addr())

		http.Serve(listener, mux)

	}()

	listener2, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, &AuthService{})
	reflection.Register(server)
	log.Printf("Start gRPC server at %s\n", listener2.Addr())
	server.Serve(listener2)
}
