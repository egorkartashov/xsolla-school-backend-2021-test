package services

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

const (
	checkUrlQueueName     = "check-url-queue"
	checkAttemptsMaxCount = 2
)

type MessageBrokerInterface interface {
	SendMessage()
}

type MessageBrokerService struct {
	connectionString string
}

func NewMessageBrokerService() *MessageBrokerService {
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")
	address := os.Getenv("RABBITMQ_ADDRESS")

	connectionString := fmt.Sprintf("amqp://%s:%s@%s/", user, pass, address)

	return &MessageBrokerService{
		connectionString: connectionString,
	}
}

func (broker *MessageBrokerService) SendUrlForCheck(url string) {
	conn, err := amqp.Dial(broker.connectionString)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
		return
	}
	defer func(conn *amqp.Connection) {
		_ = conn.Close()
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to create channel with RabbitMQ: %s", err)
		return
	}
	defer func(ch *amqp.Channel) {
		_ = ch.Close()
	}(ch)

	q, err := ch.QueueDeclare(
		checkUrlQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to declare queue: %s", err)
		return
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(url),
		Headers: amqp.Table{
			"ttl": checkAttemptsMaxCount,
		},
	})

	if err != nil {
		log.Printf("Failed to publish message: %s", err)
		return
	}
}
