package io.conduktor.kafka;

import org.apache.kafka.clients.producer.*;
import org.apache.kafka.common.serialization.StringSerializer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Properties;

// kafka-topics --bootstrap-server localhost:9092 --topic demo_java --create --partitions 3

public class ProducerWithCallback {

    private static final Logger log = LoggerFactory.getLogger(ProducerWithCallback.class.getSimpleName());

    public static void main(String[] args) {
        log.info("Starting Producer Callback");

        // Create Producer properties
        Properties properties = new Properties();

        // Connect to Localhost
        properties.setProperty("bootstrap.servers", "localhost:9092");
        properties.setProperty("security.protocol", "PLAINTEXT");

        // Set serializers, serialize key and value into bytes before sending message to Kafka
        properties.setProperty("key.serializer", StringSerializer.class.getName());
        properties.setProperty("value.serializer", StringSerializer.class.getName());

        properties.setProperty("batch.size", "400");

        // NOT RECOMMENDED IN PRODUCTION, USE BATCHING
//        properties.setProperty("partitioner.class", RoundRobinPartitioner.class.getName());

        // Create Producer
        // <String, String> means key is type string and value is type string, will be serialized by StringSerializer
        KafkaProducer<String, String> producer = new KafkaProducer<>(properties);

        for (int j = 0; j < 10; j++) {
            // Batch of 30 messages
            for (int i = 0; i < 30; i++) {
                // Create a Producer record
                ProducerRecord<String, String> record =
                        new ProducerRecord<>("demo_java", "Hello Java " + i);

                // Send data to producer with callback function
                producer.send(record, new Callback() {
                    @Override
                    public void onCompletion(RecordMetadata metadata, Exception e) {
                        // Executes every time a record is successfully sent or an exception is thrown
                        if (e == null) {
                            // Record was successfully sent
                            log.info("Topic: {}, Partition: {}, Offset: {}, Timestamp: {}",
                                    metadata.topic(), metadata.partition(), metadata.offset(), metadata.timestamp());
                        } else {
                            log.error("Error while producing", e);
                        }
                    }
                });
            }

            try {
                Thread.sleep(500);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        // Tell the producer to send all data and block until done -- synchronous
        producer.flush();

        // Flush and close the producer
        producer.close();
    }
}
