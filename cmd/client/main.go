package main

// client

import (
	"fmt"
	"grpcProject/constants"
	"time"

	"github.com/micro/go-micro/v2/broker"
	_ "github.com/micro/go-plugins/broker/kafka/v2"

	"github.com/micro/micro/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	// Read go micro cmd line args
	cmd.Init()

	// Kafka
	if err := broker.Init(); err != nil {
		logrus.Errorf("Broker Init error: %v", err)
		return
	}
	if err := broker.Connect(); err != nil {
		logrus.Errorf("Broker Connect error: %v", err)
		return
	}
	go sub()

	// sub := new(kafka.Subscriber)
	// sub.Subscribe(constants.KafkaTopic)
	<-time.After(time.Second * 100)

	// wrapper := new(middleware.Wrapper)

	// create and initialise a new service
	// service := service.New(
	// 	service.Name(constants.ClientName),
	// 	service.WrapSubscriber(wrapper.NewWrapSubscriber()),
	// )
	// service.Init()

	// subscribe to the topic
	// subscriber := pubsub.NewSubscriber()
	// service.Subscribe(constants.PubsubName, subscriber.Subscribe)

	// if err := service.Run(); err != nil {
	// 	logrus.Error(err)
	// }
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

// package main

// import (
// 	"fmt"
// 	"grpcProject/constants"
// 	"time"

// 	"github.com/micro/go-micro/v2/broker"
// 	_ "github.com/micro/go-plugins/broker/kafka/v2"
// 	"github.com/micro/micro/cmd"
// )

var (
	topic = constants.KafkaTopic
)

func sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

// func main() {
// 	cmd.Init()

// 	// if err := broker.Init(); err != nil {
// 	// 	log.Fatalf("Broker Init error: %v", err)
// 	// }
// 	// if err := broker.Connect(); err != nil {
// 	// 	log.Fatalf("Broker Connect error: %v", err)
// 	// }

// 	// go pub()

// 	<-time.After(time.Second * 10)
// }
