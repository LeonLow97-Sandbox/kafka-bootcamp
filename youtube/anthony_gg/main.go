package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const topic = "HVSE"

type OrderPlacer struct {
	producer *kafka.Producer
	topic    string
	// Channel for delivery reports from Kafka regarding the status of the
	// produced message to ensure successful message delivery
	deliverych chan kafka.Event
}

func NewOrderPlacer(p *kafka.Producer, topic string) *OrderPlacer {
	return &OrderPlacer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 10000),
	}
}

func (op *OrderPlacer) placeOrder(orderType string, size int) error {
	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &op.topic, Partition: kafka.PartitionAny},
		Value:          payload,
	},
		op.deliverych,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Waiting for the delivery report from Kafka and printing its details
	<-op.deliverych

	fmt.Printf("placed order on the queue --> %s\n", format)

	return nil
}

func main() {
	// Creating a new Kafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Kafka broker address
		"client.id":         "something",      // Unique client ID for the producer
		"acks":              "all",            // Acknowledgment mode for message durability
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	op := NewOrderPlacer(p, topic)

	for i := 0; i < 1000; i++ {
		if err := op.placeOrder("BUY", i+1); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 3)
	}

}
