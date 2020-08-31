package kafka

import (
	"fmt"

	"github.com/micro/go-micro/broker"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	sub broker.Subscriber
}

func (subscriber *Subscriber) Subscribe(topic string) error {

	go func() error {
		_, err := broker.Subscribe(topic, func(p broker.Event) error {
			fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
			return nil
		})
		if err != nil {
			logrus.Error("Could not subscribe to kafka", err)
			return err
		}
		return nil
	}()
	// _, err := broker.Subscribe(topic, printEvent)
	// if err != nil {
	// 	logrus.Debug(err)
	// 	return err
	// }
	// // subscriber.sub = s
	// return nil
	return nil

}

// func sub() {
// 	_, err := broker.Subscribe(topic, func(p broker.Event) error {
// 		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

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
