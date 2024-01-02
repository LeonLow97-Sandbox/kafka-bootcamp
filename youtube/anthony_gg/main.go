package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	topic := "HVSE" // Kafka topic to produce messages to

	// Creating a new Kafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Kafka broker address
		"client.id":         "something",      // Unique client ID for the producer
		"acks":              "all",            // Acknowledgment mode for message durability
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	// Kafka Consumer
	go func() {
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
				fmt.Printf("Consumed message from the queue: %s\n", string(e.Value))
			case *kafka.Error:
				// Handling Kafka errors, if any
				fmt.Printf("%v\n", e)
			}
		}
	}()

	// Channel for delivery reports from Kafka regarding the status of the
	// produced message to ensure successful message delivery
	deliverch := make(chan kafka.Event, 10000)

	for {
		// Producing a Kafka message with payload "FOO" to the specified topic
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("FOO"), // Payload of the message
		},
			deliverch,
		)
		if err != nil {
			log.Fatal(err)
		}

		// Waiting for the delivery report from Kafka and printing its details
		<-deliverch
		time.Sleep(time.Second * 3)
	}
}
