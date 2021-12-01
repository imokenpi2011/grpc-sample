package main

import (
	"fmt"
	"log"
	"net"

	"grpc-sample/chat"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	// 通信方式を指定
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// chatサービスの定義
	s := chat.Server{}

	// サーバを起動
	grpcServer := grpc.NewServer()

	// サーバにchatサービスを登録
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
