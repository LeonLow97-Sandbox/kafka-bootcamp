package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func producer() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "test-topic", 0)
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(
		kafka.Message{
			Value: []byte("Lebron James"),
		},
		kafka.Message{
			Value: []byte("Michael Jordan"),
		},
		kafka.Message{
			Value: []byte("Kobe Bryant"),
		},
		kafka.Message{
			Value: []byte("Stephen Curry"),
		},
	)
}
