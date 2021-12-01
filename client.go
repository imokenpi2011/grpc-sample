package main

import (
	"context"
	"grpc-sample/chat"
	"log"

	"google.golang.org/grpc"
)

func main() {

	// クライアントを定義しコネクションを形成
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	// 最後にコネクションを閉じる様記載
	defer conn.Close()

	// chatサービスの定義
	c := chat.NewChatServiceClient(conn)

	// サーバー側からメッセージを受け取り表示する
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
