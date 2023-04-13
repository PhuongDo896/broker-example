package main

import (
	"log"
	"os"

	"github.com/PhuongDo896/rabbitmq-example/rabbitmq"
)

const (
	PRODUCER = "producer"
	CONSUMER = "consumer"
)

func main() {
	connection, err := rabbitmq.OpenConnection()
	if err != nil {
		log.Fatalf("failed connection: %s", err)
	}

	defer func() {
		if err := connection.Close(); err != nil {
			log.Fatalf("failed close connection: %s", err)
		}
	}()

	channel, err := rabbitmq.NewChannel(connection).Create()
	if err != nil {
		log.Fatalf("failed create channel: %s", err)
	}

	queue, err := rabbitmq.NewQueue(channel, "message-broker").Create()
	if err != nil {
		log.Fatalf("failed queue declare: %s", err)
	}

	var message = "Random message!"

	switch os.Args[1] {
	case PRODUCER:
		if err := rabbitmq.NewProducer(channel, queue.Name).Publish(message); err != nil {
			log.Fatalf("failed publish message: %s", err)
		}

	case CONSUMER:
		if err := rabbitmq.NewConsumer(channel, queue.Name).Consume(); err != nil {
			log.Fatalf("failed consume: %s", err)
		}
	}
}
