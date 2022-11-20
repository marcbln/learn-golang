package main

import (
	"fmt"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	protos "mymodule5/protos/currency"
	"mymodule5/server"
	"net"
	"os"
)

func main() {
	const ADDRESS = ":9092"

	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gs, cs)

	// enable reflection api so that grpcurl can use it: grpcurl --plaintext localhost:9092 list (remove on production)
	reflection.Register(gs)

	fmt.Printf("listening to address %s\n", ADDRESS)
	l, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
