package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {

	// 通信方式を指定
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// サーバを起動
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
