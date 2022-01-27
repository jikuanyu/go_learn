package main

import (
	"fmt"
	_ "github.com/asim/go-micro/plugins/broker/kafka/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cmd"
	"log"
	"os"
	"time"
)

var (
	topic = "goTopic"
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s 时间", i, time.Now().String())),
		}
		if err := broker.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
		if i == 20 {
			break
		}
	}
	defer tick.Stop()
}

func main() {
	os.Setenv("MICRO_BROKER", "kafka")
	os.Setenv("MICRO_BROKER_ADDRESS", "192.168.144.17:9092")

	cmd.Init()
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	/*newBroker = kafka.NewBroker(broker.Addrs("192.168.144.17:9092"))
	if err := newBroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := newBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}*/
	pub()
}
