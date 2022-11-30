package nats

import (
	natsClient "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"mymodule11-nats/data"
)

type NatsPublisher struct {
	url string
}

// factory method
func NewNatsPublisher(url string) *NatsPublisher {
	return &NatsPublisher{url}
}

func (natsPublisher *NatsPublisher) PublishNewUserRegistration(userRegistration *data.UserRegistration) error {
	// ---- logging
	notifierLogger := log.WithField("UserRegistration", userRegistration.RequestId)
	notifierLogger.Info("Inform about new user registration")
	notifierLogger.Infof("connecting to NATS server at %s", natsPublisher.url)

	// ---- nats
	nc, err := natsClient.Connect(natsPublisher.url)
	if err != nil {
		notifierLogger.WithError(err).Error("Could not connect to server: " + err.Error())
		return err
	}
	defer nc.Close()

	// ---- json encoded connection
	ec, err := natsClient.NewEncodedConn(nc, natsClient.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	// ---- publish notification on nats ... check on cli with: nats subscribe "theproject.userregistration.new"
	notifierLogger.Infof("publishing message with subject %q", SUBJECT)
	return ec.Publish(SUBJECT, userRegistration)

}
