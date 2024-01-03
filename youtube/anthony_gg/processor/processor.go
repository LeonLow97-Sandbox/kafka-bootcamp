package main

import (
	"fmt"
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

	// Continuous polling for new messages
	for {
		// Polling Kafka for events (messages or errors) every 100ms
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			// Handling consumed Kafka messages
			fmt.Printf("Processing order: %s\n", string(e.Value))
		case *kafka.Error:
			// Handling Kafka errors, if any
			fmt.Printf("%v\n", e)
		}
	}
}
