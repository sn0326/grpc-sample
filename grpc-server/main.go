package main

import (
	"log"
	"net"

	"github.com/sn0326/grpc-sample/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	log.Print("main start")

	// 9000番ポートでクライアントからのリクエストを受け付けるようにする
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Sample構造体のアドレスを渡すことで、クライアントからGetDataリクエストされると
	// GetDataメソッドが呼ばれるようになる
	service.RegisterSampleServiceServer(grpcServer, &Sample{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Print("main end")
}

type Sample struct {
	name string
}

func (s *Sample) GetData(
	ctx context.Context,
	message *service.Message,
) (*service.Message, error) {
	log.Print(message.Body)
	return &service.Message{Body: "レスポンスデータ"}, nil
}
