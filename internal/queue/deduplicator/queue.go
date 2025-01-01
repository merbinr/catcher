package deduplicator_queue

import (
	"fmt"
	"os"

	"github.com/merbinr/catcher/internal/config"
	"github.com/streadway/amqp"
)

type RabbitMQConnForDedupe struct {
	Queue_Client  *amqp.Connection
	Queue_Channel *amqp.Channel
	Queue         amqp.Queue
}

var DedupeRabbitmqConn RabbitMQConnForDedupe

func CreateDedupeQueueConn() error {
	var err error

	password := os.Getenv("CATCHER_OUTGOING_QUEUE_PASSWORD")
	if password == "" {
		return fmt.Errorf("queue password is not set, please set the CATCHER_OUTGOING_QUEUE_PASSWORD env")
	}
	host := os.Getenv("CATCHER_OUTGOING_QUEUE_HOST")
	if host == "" {
		return fmt.Errorf("queue host is not set, please set the CATCHER_OUTGOING_QUEUE_HOST env")
	}

	conn_string := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.ConfigData.RabbitMq.User,
		password,
		host,
		config.ConfigData.RabbitMq.Port)

	DedupeRabbitmqConn.Queue_Client, err = amqp.Dial(conn_string)
	if err != nil {
		return fmt.Errorf("unable to open the dedupe queue connection, error: %s", err)
	}

	DedupeRabbitmqConn.Queue_Channel, err = DedupeRabbitmqConn.Queue_Client.Channel()
	if err != nil {
		return fmt.Errorf("unable to open the dedupe queue channel, error: %s", err)
	}

	queueName := config.ConfigData.RabbitMq.Name
	DedupeRabbitmqConn.Queue, err = DedupeRabbitmqConn.Queue_Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("unable to create dedupe queue from queue channel")
	}
	return nil
}

func (r RabbitMQConnForDedupe) SendLogMessageToDedupeQueue(msg *[]byte) error {
	err := r.Queue_Channel.Publish(
		"",           // exchange
		r.Queue.Name, // routing key (queue name)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        *msg,
		},
	)
	if err != nil {
		return fmt.Errorf("unable to send message to dedupe queue")
	}
	return nil
}
