package infra

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	Channel *amqp.Channel
	Queue   string
}

func NewRabbitMQConsumer() *RabbitMQConsumer {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	log.Println("RabbitMQ consumer initialized and connected to queue:", "email-queue")

	return &RabbitMQConsumer{
		Channel: ch,
		Queue:   "email-queue",
	}
}

func (c *RabbitMQConsumer) StartConsuming(handleMessage func(body []byte)) {
	msgs, err := c.Channel.Consume(
		c.Queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("Consumer started. Waiting for messages...")
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", msg.Body)
			handleMessage(msg.Body)
		}
	}()
	<-forever
}
