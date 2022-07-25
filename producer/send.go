package main

import (
	"log"

	"github.com/turgut-nergin/rabbit_mq/broker"
	"github.com/turgut-nergin/rabbit_mq/config"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	config := config.EnvConfig["local"]
	conn := broker.CreateConnection(config)
	defer conn.Close()

	ch := broker.CreateChannel(conn)
	defer ch.Close()

	q := broker.CreateQueue(ch)
	body := "WOW"

	err := broker.PublishConsumer(ch, q, body)
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s\n", body)
}
