package main

// client

import (
	"grpcProject/constants"
	"grpcProject/kafka"
	"grpcProject/middleware"
	"grpcProject/pubsub"

	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/v3/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	// Read go micro cmd line args
	cmd.Init()

	// Kafka
	if err := broker.Init(); err != nil {
		logrus.Errorf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		logrus.Errorf("Broker Connect error: %v", err)
	}
	sub := new(kafka.Subscriber)
	sub.Subscribe(constants.KafkaTopic)

	wrapper := new(middleware.Wrapper)

	// create and initialise a new service
	service := service.New(
		service.Name(constants.ClientName),
		service.WrapSubscriber(wrapper.NewWrapSubscriber()),
	)
	// service.Init()

	// subscribe to the topic
	subscriber := pubsub.NewSubscriber()
	service.Subscribe(constants.PubsubName, subscriber.Subscribe)

	if err := service.Run(); err != nil {
		logrus.Error(err)
	}
}

// ##
// ## Try implement client to server
// ##
// client := proto.NewCrudService("grpcproject", nil)
// time.Sleep(time.Second * 10)

// for i := 0; i < 10; i++ {
// 	rsp, err := client.GetItem(context.Background(), &proto.ID{
// 		Id: "id",
// 	})
// 	if err != nil {
// 		logrus.Errorf("Error: %v", err)
// 		// return
// 	}

// 	// print the response
// 	logrus.Infof("Response: %+v", rsp)

// 	// let's delay the process for exiting for reasons you'll see below
// 	time.Sleep(time.Second * 10)
// }
