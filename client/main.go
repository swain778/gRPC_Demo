package main

import (
	"log"

	pb "gRPC-Demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Shubham", "Rahul", "Mohit", "Akash"},
	}

	//Unary
	callSayHello(client)

	//Server stream
	callSayHelloServerStreaming(client, names)

	//Client stream
	callSayHelloClientStreaming(client, names)

	//Bidirectional stream
	callHelloBidirectionalStream(client, names)
}
