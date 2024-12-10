package rabbitmq

import (
	"fmt"
	"os"

	"github.com/merbinr/catcher/internal/config"
	"github.com/streadway/amqp"
)

type RabbitMQConn struct {
	Queue_Client  *amqp.Connection
	Queue_Channel *amqp.Channel
	Queue         amqp.Queue
}

var Rabbitmq_conn RabbitMQConn

func CreateQueueConn() error {
	var err error

	password := os.Getenv("CATCHER_RABBITMQ_PASSWORD")
	if password == "" {
		return fmt.Errorf("queue password is not set, please set the CATCHER_RABBITMQ_PASSWORD env")
	}
	host := os.Getenv("CATCHER_OUTGOING_QUEUE_HOST")
	if host == "" {
		return fmt.Errorf("queue host is not set, please set the CATCHER_OUTGOING_QUEUE_HOST env")
	}

	conn_string := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.ConfigData.RabbitMQ.User,
		password,
		host,
		config.ConfigData.RabbitMQ.Port)

	Rabbitmq_conn.Queue_Client, err = amqp.Dial(conn_string)
	if err != nil {
		return fmt.Errorf("unable to open the queue connection, error: %s", err)
	}

	Rabbitmq_conn.Queue_Channel, err = Rabbitmq_conn.Queue_Client.Channel()
	if err != nil {
		return fmt.Errorf("unable to open the queue channel, error: %s", err)
	}

	queueName := config.ConfigData.RabbitMQ.Name
	Rabbitmq_conn.Queue, err = Rabbitmq_conn.Queue_Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("unable to create queue from queue channel")
	}
	return nil
}

func (r RabbitMQConn) SendLogMessage(msg []byte) error {
	err := r.Queue_Channel.Publish(
		"",           // exchange
		r.Queue.Name, // routing key (queue name)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
	if err != nil {
		return fmt.Errorf("unable to send message to queue")
	}
	return nil
}
