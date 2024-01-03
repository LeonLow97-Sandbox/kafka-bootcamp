## Apache Kafka with Go

- [Course Link](https://www.youtube.com/watch?v=-yVxChp7HoQ&t=987s)
- [Kafka Golang](https://docs.confluent.io/kafka-clients/go/current/overview.html)

## Multiple Consumers reading from the same Partition

- Use different consumer group id
- By doing so, these 2 consumers can read data from the same queue at the same time.

```go
// Consumer 1
consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
    "bootstrap.servers": "localhost:9092",
    "group.id":          "website",
    "auto.offset.reset": "smallest",
})

// Consumer 2
consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
    "bootstrap.servers": "localhost:9092",
    "group.id":          "website_data", // different consumer group id
    "auto.offset.reset": "smallest",
})
```
