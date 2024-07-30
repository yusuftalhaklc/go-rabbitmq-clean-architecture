package infra

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewRabbitMQProducer() *RabbitMQProducer {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	q, err := ch.QueueDeclare(
		"email-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	return &RabbitMQProducer{
		Channel: ch,
		Queue:   q,
	}
}

func (p *RabbitMQProducer) Publish(body []byte) error {
	err := p.Channel.Publish(
		"",
		p.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return err
	}
	log.Printf("Published message: %s", body)
	return nil
}
