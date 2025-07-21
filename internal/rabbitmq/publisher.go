package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func Publish(message string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err)
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"orders",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("Failed to declare queue:", err)
		return err
	}

	err = ch.Publish(
		"",      
		"orders", 
		false,    
		false,   
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Println("Failed to publish message:", err)
		return err
	}

	log.Println("Published message:", message)
	return nil
}
