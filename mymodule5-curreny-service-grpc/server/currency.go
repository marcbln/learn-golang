package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	protos "mymodule5/protos/currency"
)

type Currency struct {
	log hclog.Logger
	// Embed the unimplemented server (to prevent breakages when new methods are added, see https://github.com/grpc/grpc-go/issues/3794)
	protos.UnimplementedCurrencyServer
}

// constructor
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

// implement the interface of CurrencyServer
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
