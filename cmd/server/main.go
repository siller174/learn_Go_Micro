package main

import (
	"fmt"
	"grpcProject/constants"
	"grpcProject/handler"
	"grpcProject/kafka"
	"grpcProject/middleware"
	"grpcProject/pubsub"
	"grpcProject/repository"
	"time"

	"github.com/micro/go-micro/config/cmd"
	_ "github.com/micro/go-plugins/broker/kafka" // todo move it
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/sirupsen/logrus"
)

// micro registry getService --service=grpcproject

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	wrapper := new(middleware.Wrapper)
	items := repository.NewItems()
	crud := handler.NewCrud(items)

	// Read go micro cmd line args
	cmd.Init()

	// Pubsub
	// startPublisher()
	startPublisherToKafka()

	// Create service
	srv := service.New(
		service.Name(constants.ServerName),
		service.Version("latest"),
		service.WrapHandler(wrapper.NewWrapHandler()),
	)

	// Register handler
	srv.Handle(crud)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func startPublisherToKafka() {
	publisher := kafka.NewPublisher(constants.KafkaTopic)
	go func() {
		for i := 0; ; i++ {
			publisher.Push(map[string]string{
				"h1": "test",
				"h2": "test2",
			}, fmt.Sprintf("[Kafka] it's %v msg", i))
			time.Sleep(2 * time.Second)
		}
	}()
}

func startPublisher() {
	publisher := pubsub.NewPublisher(constants.PubsubName)

	go func() {
		for i := 0; ; i++ {
			publisher.Push("ID HERE", fmt.Sprintf("[pubsub] it's %v msg", i))
			time.Sleep(2 * time.Second)
		}
	}()
}

// continue here https://itnext.io/micro-in-action-part-2-71230f01d6fb
