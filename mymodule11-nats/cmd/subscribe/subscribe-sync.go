package main

import (
	natsClient "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"mymodule11-nats/nats"
	"time"
)

func main() {
	nc, err := natsClient.Connect(natsClient.DefaultURL)
	if err != nil {
		log.WithError(err).Fatal("Could not connect to server")
	}

	// ---- sync subscribe
	sub, err := nc.SubscribeSync(nats.SUBJECT)
	if err != nil {
		log.WithError(err).WithField("subject", nats.SUBJECT).Fatal("could not subscribe")
	}
	m, err := sub.NextMsg(20 * time.Second)
	if m == nil {
		log.Println("No messages available")
		return
	}

	log.WithField("subject", m.Subject).Println("received message")
	log.WithField("data", string(m.Data)).Print("received data")
	sub.Unsubscribe()
	sub.Drain()
}
