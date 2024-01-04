## Apache Kafka Broker on docker-compose

1. Run `docker-compose up`
1. Create Topic in kafka broker container
    - Run `docker-compose exec broker bash`
    - create the topic `kafka-topics --create --topic test-topic --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1`
    - list the topic `kafka-topics --list --bootstrap-server localhost:9092`
1. Test Read and Write to the topic
    - write to topic `kafka-console-producer --topic test-topic --bootstrap-server localhost:9092`
    - read from topic `kafka-console-consumer --topic test-topic --from-beginning --bootstrap-server localhost:9092`

```
// Create Topic
kafka-topics --create --topic test-topic --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1

// List Topics
kafka-topics --list --bootstrap-server localhost:9092

// Write to Topic
kafka-console-producer --topic test-topic --bootstrap-server localhost:9092

// Read Topic
kafka-console-consumer --topic test-topic --from-beginning --bootstrap-server localhost:9092
```