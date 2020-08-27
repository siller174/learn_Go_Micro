package pubsub

import (
	"context"
	"grpcProject/proto/msg"

	"github.com/micro/micro/v3/service"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	ev *service.Event
}

func NewPublisher(name string) *Publisher {
	ev := service.NewEvent(name)

	return &Publisher{
		ev: ev,
	}
}

func (publisher *Publisher) Push(id, text string) {
	logrus.Infof("Push message %v to pubsub with id %v", text, id)
	publisher.ev.Publish(context.TODO(), &msg.Msg{ //todo remove ctx
		Id:   id,
		Text: text,
	})
}
