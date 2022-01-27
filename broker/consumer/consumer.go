package main

import (
	"fmt"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cmd"
	"log"
	"os"

	// To enable rabbitmq plugin uncomment
	//_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/asim/go-micro/plugins/broker/kafka/v4"
)

var (
	topic = "goTopic"
)

// Example of a shared subscription which receives a subset of messages
func sharedSub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("consumer"))
	if err != nil {
		fmt.Println(err)
	}
}

// Example of a subscription which receives all the messages
func sub() {

	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func sub2() {

	//config := &sarama.Config{}
	//config.Consumer.Offsets.Initial=sarama.OffsetOldest

	/*ctx := context.Background()
	context.WithValue(ctx,config.Consumer.Offsets.Initial,sarama.OffsetOldest)*/
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub2] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("myGroup2"))

	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	//	&cli.StringFlag{
	//			Name:    "broker",
	//			EnvVars: []string{"MICRO_BROKER"},
	//			Usage:   "Broker for pub/sub. http, nats, rabbitmq",
	//		},
	//		&cli.StringFlag{
	//			Name:    "broker_address",
	//			EnvVars: []string{"MICRO_BROKER_ADDRESS"},
	//			Usage:   "Comma-separated list of broker addresses",
	//		},
	os.Setenv("MICRO_BROKER", "kafka")
	os.Setenv("MICRO_BROKER_ADDRESS", "192.168.144.17:9092")
	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	//sub()
	sub2()
	select {}
}
