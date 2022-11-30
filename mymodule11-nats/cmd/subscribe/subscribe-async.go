package main

import (
	natsClient "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"mymodule11-nats/nats"
	"sync"
)

func main() {
	nc, err := natsClient.Connect(natsClient.DefaultURL)
	if err != nil {
		log.WithError(err).Fatal("Could not connect to server")
	}

	// ---- json encoded connection
	ec, err := natsClient.NewEncodedConn(nc, natsClient.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	// ---- async subscribe
	sub, err := ec.Subscribe(nats.SUBJECT, func(msg *natsClient.Msg) {
		log.WithField("subject", msg.Subject).Println("received message")
		log.WithField("data", string(msg.Data)).Print("received data")
	})
	if err != nil {
		log.WithError(err).WithField("subject", nats.SUBJECT).Fatal("could not subscribe")
	}

	// ---- wait
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Wait for a message to come in
	wg.Wait()

	// ---- clean up
	sub.Unsubscribe()
	sub.Drain()

}
