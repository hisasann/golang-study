package main

import (
	"github.com/hello-grpc/grpc_hello_grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := grpc_hello_grpc.NewHelloGrpcServiceClient(conn)

	response, err := c.GetData(context.Background(), &grpc_hello_grpc.Message{Body: "hello world from client"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Print(response.Body)

	defer conn.Close()
}
