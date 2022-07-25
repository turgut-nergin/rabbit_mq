package broker

import (
	"fmt"

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
		"signal", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	errors.FailOnError(err, "Failed to declare a queue")
	return q
}

func PublishConsumer(ch *amqp.Channel, q amqp.Queue, body string) error {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	return err
}
