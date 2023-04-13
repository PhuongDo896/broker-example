package main

import (
	"log"
	"os"

	"github.com/PhuongDo896/rabbitmq-example/activemq"
)

const (
	PRODUCER = "producer"
	CONSUMER = "consumer"
)

func main() {
	connection, err := activemq.Connect()
	if err != nil {
		log.Fatalf("failed connection: %s", err)
	}

	defer func() {
		if err := connection.Disconnect(); err != nil {
			log.Fatalf("failed close connection: %s", err)
		}
	}()

	var (
		message = "Random message to activemq!"
		queue   = "message-broker"
	)

	switch os.Args[1] {
	case PRODUCER:
		if err := activemq.NewProducer(connection, queue).Publish(message); err != nil {
			log.Fatalf("failed publish message: %s", err)
		}

	case CONSUMER:
		activemq.NewConsumer(connection, queue).Consume()
	}
}
