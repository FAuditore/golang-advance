package main

import (
	pb "advance/rpc/gRPC/order"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var dialOptions []grpc.DialOption

func main() {
	//creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	//if err != nil {
	//	panic(err)
	//}
	//dialOptions = append(dialOptions, grpc.WithTransportCredentials(creds))
	dialOptions = append(dialOptions, grpc.WithInsecure())
	conn, err := grpc.Dial(":8080", dialOptions...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewOrderClient(conn)

	//Get(client)
	//time.Sleep(500 * time.Millisecond)

	GetAll(client)
	//time.Sleep(500 * time.Millisecond)

	//Add(client)

	//Chat(client)
}

func Get(client pb.OrderClient) {
	fmt.Println("Get call...")
	res, err := client.Get(context.Background(), &pb.OrderRequest{
		Id: 0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}

func GetAll(client pb.OrderClient) {
	fmt.Println("Get All call...")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stream, err := client.GetAll(ctx, &pb.EmptyRequest{})
	if err != nil {
		panic(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%+v\n", res)
	}
}

func Add(client pb.OrderClient) {
	fmt.Println("Add call...")
	md := metadata.New(map[string]string{"no": "2020"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := client.Add(ctx)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		datas := strings.Split(scanner.Text(), " ")
		id, err := strconv.ParseInt(datas[0], 10, 64)
		if err != nil {
			panic(err)
		}
		order := &pb.OrderInfo{
			Id:    id,
			Name:  datas[1],
			Desc:  datas[2],
			State: pb.Status_VALID,
		}
		err = stream.Send(order)
		if err != nil {
			panic(err)
		}
		recv, err := stream.CloseAndRecv()
		if err != nil {
			panic(err)
		}
		fmt.Printf("recv: %+v\n", recv)
	}
}

func Chat(client pb.OrderClient) {
	fmt.Println("Chat call...")
	stream, err := client.Chat(context.Background())
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(res.GetMessage())
		}
	}()
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		err := stream.Send(&pb.ClientMessage{Message: scanner.Text()})
		if err != nil {
			panic(err)
		}
	}
	stream.CloseSend()
}
