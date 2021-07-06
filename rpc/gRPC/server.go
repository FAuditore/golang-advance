package main

import (
	pb "advance/rpc/gRPC/order"
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"time"
)

var serverOptions []grpc.ServerOption

var orders = []*pb.OrderInfo{
	{
		Id:    0,
		Name:  "aaaa",
		Desc:  "美团",
		State: pb.Status_VALID,
	},
	{
		Id:    1,
		Name:  "bbb",
		Desc:  "饿了么",
		State: pb.Status_VALID,
	},
	{
		Id:    2,
		Name:  "ccc",
		Desc:  "京东",
		State: pb.Status_VALID,
	},
}

type OrderService struct {
}

func (srv *OrderService) Get(ctx context.Context, req *pb.OrderRequest) (*pb.OrderInfo, error) {
	for _, v := range orders {
		if v.Id == req.GetId() {
			return v, nil
		}
	}
	return nil, errors.New("order not found")
}

func (srv *OrderService) GetAll(req *pb.EmptyRequest, stream pb.Order_GetAllServer) error {
	for _, v := range orders {
		select {
		case <-stream.Context().Done():
			return processDone(stream.Context())
		default:
			err := stream.Send(v)
			if err != nil {
				return err
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func (srv *OrderService) Add(stream pb.Order_AddServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		fmt.Printf("metadata :%+v\n", md)
	}
	for {
		select {
		case <-stream.Context().Done():
			return processDone(stream.Context())
		default:
			addOrder, err := stream.Recv()
			if err == io.EOF {
				goto finish
			}
			if err != nil {
				return err
			}

			if addOrder.Id <= 0 {
				return status.Errorf(codes.InvalidArgument, "Invalid argument: %d", addOrder.Id)
			}
			for _, v := range orders {
				if v.Id == addOrder.Id {
					return status.Error(codes.AlreadyExists, "Order alreadyExist")
				}
			}
			orders = append(orders, addOrder)
			fmt.Printf("Rpc add :%+v\n", addOrder)
		}
	}
finish:
	return stream.SendAndClose(&pb.AddResponse{Success: true})
}

func (srv *OrderService) Chat(stream pb.Order_ChatServer) error {
	finish := make(chan error)
	go func() {
		defer fmt.Println("Chat closed...")
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				fmt.Println(msg.GetMessage())
				break
			}
			if err != nil {
				finish <- err
			}
			err = stream.Send(&pb.ServerMessage{Message: "recv : " + msg.GetMessage()})
			if err != nil {
				panic(err)
			}
		}
	}()
	return <-finish
}

func processDone(ctx context.Context) error {
	if ctx.Err() == context.Canceled {
		log.Println("Connection has been closed...")
		return status.Error(codes.Canceled, "Connection has been closed")
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Connection deadline exceeded...")
		return status.Error(codes.DeadlineExceeded, "Deadline Exceeded")
	}
	log.Println("Unknown error...")
	return status.Error(codes.Unknown, "Unknown error")
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("-->unary interceptor: ", info.FullMethod)
	i, err := handler(ctx, req)
	if err != nil {
		log.Printf("\033[31m[Error occured]: - %v\033[0m\n", err)
	}
	return i, err
}

func streamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("-->stream interceptor: ", info.FullMethod)
	err := handler(srv, ss)
	if err != nil {
		log.Printf("\033[31m[Error occured]: - %v\033[0m\n", err)
	}
	return err
}

func main() {
	var port *int
	port = flag.Int("port", 8080, "port")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	//creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	//if err != nil {
	//	panic(err)
	//}
	//serverOptions := []grpc.ServerOption{grpc.Creds(creds)}
	//拦截器
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(unaryInterceptor))
	serverOptions = append(serverOptions, grpc.StreamInterceptor(streamServerInterceptor))

	server := grpc.NewServer(serverOptions...)

	pb.RegisterOrderServer(server, &OrderService{})

	reflection.Register(server)

	//\033[显示方式;字体颜色;背景颜色m 中间是变颜色的内容 \033[0m
	log.Printf("\033[36mgRPC server started at :%d\033[0m\n", *port)
	server.Serve(listener)
}
