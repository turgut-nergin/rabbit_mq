package broker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/turgut-nergin/rabbit_mq/config"
	"github.com/turgut-nergin/rabbit_mq/errors"
)

func CreateConnection(config config.Config) *amqp.Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.UserName, config.Password, config.Host, config.Port)
	conn, err := amqp.Dial(url)
	errors.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	errors.FailOnError(err, "Failed to open a channel")
	return ch
}

func CreateQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	errors.FailOnError(err, "Failed to declare a queue")
	return q
}

func RegisterConsumer(ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	errors.FailOnError(err, "Failed to register a consumer")
	return msgs
}

func HandleDelivery(msgs <-chan amqp.Delivery) {
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}
