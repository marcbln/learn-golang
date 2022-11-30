package main

import (
	"github.com/google/uuid"
	natsClient "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"mymodule11-nats/data"
	"mymodule11-nats/nats"
	"time"
)

func main() {
	userRegistration := &data.UserRegistration{
		RequestId:     uuid.NewString(),
		Firstname:     "Peter",
		Lastname:      "Lustig",
		TermsAccepted: true,
		Datum:         time.Now(),
		Schulungscode: "GO.EINF",
		Email:         "user1@example.org",
	}
	natsPublisher := nats.NewNatsPublisher(natsClient.DefaultURL)
	err := natsPublisher.PublishNewUserRegistration(userRegistration)
	if err != nil {
		log.Fatal(err)
	}

}
