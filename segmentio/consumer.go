package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func consumer() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "test-topic", 0)
	_ = conn.SetDeadline(time.Now().Add(time.Second * 8))

	// m, _ := conn.ReadMessage(1e6)
	batch := conn.ReadBatch(1e3, 1e9) // 1000
	
	bytes := make([]byte, 1e9)
	for {
		_, err := batch.Read(bytes)
		// if there is an error, stop reading messages
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}
