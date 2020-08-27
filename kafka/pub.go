package kafka

import (
	"github.com/micro/go-micro/broker"

	"github.com/sirupsen/logrus"
)

type Publisher struct {
	topic string
}

func NewPublisher(topic string) *Publisher {
	// Kafka
	if err := broker.Init(); err != nil {
		logrus.Errorf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		logrus.Errorf("Broker Connect error: %v", err)
	}

	return &Publisher{
		topic: topic,
	}
}

func (publisher *Publisher) Push(headers map[string]string, message string) {

	msg := &broker.Message{
		Header: headers,
		Body:   []byte(message),
	}
	logrus.Infof("Push message %v to kafka with headers %v", message, headers)

	if err := broker.Publish(publisher.topic, msg); err != nil {
		logrus.Errorf("failed: %v", err)
	} else {
		logrus.Debug("Pubbed message:", string(msg.Body))
	}
}
