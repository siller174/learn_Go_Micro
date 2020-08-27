package pubsub

import (
	"context"
	"grpcProject/proto/msg"

	"github.com/sirupsen/logrus"
)

type Subscriber struct {
}

func NewSubscriber() *Subscriber {

	return &Subscriber{}
}

// Sub processes messages
func (subscriber *Subscriber) Subscribe(ctx context.Context, msg *msg.Msg) error {
	logrus.Infof("Received a message %v \n", msg)
	return nil
}
