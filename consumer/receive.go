package main

import (
	"log"

	"github.com/turgut-nergin/rabbit_mq/broker"
	"github.com/turgut-nergin/rabbit_mq/config"
)

func main() {
	config := config.EnvConfig["local"]
	conn := broker.CreateConnection(config)
	defer conn.Close()

	ch := broker.CreateChannel(conn)
	defer ch.Close()

	q := broker.CreateQueue(ch)
	msgs := broker.RegisterConsumer(ch, q)

	broker.HandleDelivery(msgs)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	var forever chan struct{}
	<-forever
}
