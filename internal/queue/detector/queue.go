package detector_queue

import (
	"fmt"
	"os"

	"github.com/merbinr/catcher/internal/config"
	"github.com/streadway/amqp"
)

type RabbitMQConnForDetector struct {
	Queue_Client  *amqp.Connection
	Queue_Channel *amqp.Channel
	Queue         amqp.Queue
}

var DetectorRabbitmqConn RabbitMQConnForDetector

func CreateDetectorQueueConn() error {
	var err error

	password := os.Getenv("CATCHER_DETECTOR_RABBITMQ_PASSWORD")
	if password == "" {
		return fmt.Errorf("queue password is not set, please set the CATCHER_DETECTOR_RABBITMQ_PASSWORD env")
	}
	host := os.Getenv("CATCHER_DETECTOR_QUEUE_HOST")
	if host == "" {
		return fmt.Errorf("queue host is not set, please set the CATCHER_DETECTOR_QUEUE_HOST env")
	}
	conn_string := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.ConfigData.RabbitMq.User,
		password,
		host,
		config.ConfigData.RabbitMq.Port)

	DetectorRabbitmqConn.Queue_Client, err = amqp.Dial(conn_string)
	if err != nil {
		return fmt.Errorf("unable to open the detector queue connection, error: %s", err)
	}

	DetectorRabbitmqConn.Queue_Channel, err = DetectorRabbitmqConn.Queue_Client.Channel()
	if err != nil {
		return fmt.Errorf("unable to open the detector queue channel, error: %s", err)
	}
	queueName := config.ConfigData.RabbitMq.Name
	DetectorRabbitmqConn.Queue, err = DetectorRabbitmqConn.Queue_Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("unable to create detector queue from queue channel")
	}
	return nil
}

func (r RabbitMQConnForDetector) SendLogMessageToDetectorQueue(msg *[]byte) error {
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
		return fmt.Errorf("unable to send message to detector queue")
	}
	return nil
}
