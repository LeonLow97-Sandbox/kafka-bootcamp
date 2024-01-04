package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const topic = "HVSE"

func main() {
	// Creating a new Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Kafka broker address
		"group.id":          "website",        // Consumer group ID
		"auto.offset.reset": "smallest",       // Reset offset to earliest available message
	})
	if err != nil {
		log.Fatal(err)
	}

	// Subscribing to the Kafka topic to consume messages
	if err = consumer.Subscribe(topic, nil); err != nil {
		log.Fatal(err)
	}
}
