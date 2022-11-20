package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	protos "mymodule5/protos/currency"
)

func main() {
	const ADDRESS = "localhost:9092"
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connecting to %v failed: %v", ADDRESS, err)
	}
	defer conn.Close()
	cc := protos.NewCurrencyClient(conn)

	rr := &protos.RateRequest{
		Base:        protos.Currencies_EUR,
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	}

	resp, err := cc.GetRate(context.Background(), rr)

	fmt.Printf("resp: %#v\n", resp)
	fmt.Printf("\n\n===== The rate is: %v\n\n", resp.Rate)
}
