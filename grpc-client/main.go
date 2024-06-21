package main

import (
	"log"

	"github.com/sn0326/grpc-sample/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := service.NewSampleServiceClient(conn)

	response, err := c.GetData(context.Background(), &service.Message{Body: "送信データ"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Print(response.Body)

	defer conn.Close()
}
