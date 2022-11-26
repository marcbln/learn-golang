package nats

import (
	natsClient "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"mymodule11-nats/data"
)

type NatsNotifier struct {
	url string
}

// factory method
func NewNatsNotifier(url string) *NatsNotifier {
	return &NatsNotifier{url}
}

func (natsNotifier *NatsNotifier) InformAboutNewUserRegistration(userRegistration *data.UserRegistration) error {
	notifierLogger := log.WithField("UserRegistration", userRegistration.RequestId)
	notifierLogger.Info("Inform about new user registration")
	notifierLogger.Infof("connecting to NATS server at %s", natsNotifier.url)
	nc, err := natsClient.Connect(natsNotifier.url)
	if err != nil {
		notifierLogger.WithError(err).Error("Could not connect to server: ")
		return err
	}
	defer nc.Close()
	c, _ := natsClient.NewEncodedConn(nc, natsClient.JSON_ENCODER)
	defer c.Close()

	// ---- publish notification on nats ... check on cli with: nats subscribe "theproject.userregistration.new"
	const SUBJECT = "theproject.userregistration.new"
	notifierLogger.Infof("publishing message with subject %q", SUBJECT)
	return c.Publish(SUBJECT, userRegistration)

}
