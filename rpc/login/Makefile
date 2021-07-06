gen:
	protoc -I=. --go_out=pb --grpc-gateway_out=pb --go-grpc_out=pb --swagger_out=swagger auth.proto
clean:
	rm pb/*
server:
	go run server.go
client:
	go run client.go