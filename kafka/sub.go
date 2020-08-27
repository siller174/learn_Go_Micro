package kafka

import (
	"github.com/micro/go-micro/broker"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	sub broker.Subscriber
}

func (subscriber *Subscriber) Subscribe(topic string) error {
	s, err := broker.Subscribe(topic, printEvent)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	subscriber.sub = s
	return nil
}

func (subscriber *Subscriber) Unsubscribe() error {
	return subscriber.sub.Unsubscribe()
}

func (subscriber *Subscriber) Topic() string {
	return subscriber.sub.Topic()
}

func printEvent(p broker.Event) error {
	logrus.Info("Received message:", string(p.Message().Body), "header", p.Message().Header)
	return nil
}
