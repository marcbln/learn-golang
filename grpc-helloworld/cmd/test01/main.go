package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "mygrpchelloworldmodule/helloworld"
	"time"
)

func main() {
	const NAME = "Hans Dieter"
	const ADDR = "localhost:50051"

	helloRequest := &pb.HelloRequest{
		Name: NAME,
	}
	helloRequest.Name = NAME
	fmt.Printf("%#v\n%v", helloRequest, helloRequest)

	// ---- connect
	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect to %v failed: %v", ADDR, err)
	}
	defer conn.Close()

	// Contact the server and print out its response.
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	greeterClient := pb.NewGreeterClient(conn)
	r, err := greeterClient.SayHello(ctx, helloRequest)
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
