package main

import (
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	protos "mymodule5/protos/currency"
	"mymodule5/server"
	"net"
	"os"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gs, cs)

	// enable reflection api so that grpcurl can use it: grpcurl --plaintext localhost:9092 list
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
