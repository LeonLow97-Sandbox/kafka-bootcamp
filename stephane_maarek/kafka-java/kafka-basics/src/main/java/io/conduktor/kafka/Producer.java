package io.conduktor.kafka;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.serialization.StringSerializer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Properties;

public class Producer {

    private static final Logger log = LoggerFactory.getLogger(Producer.class.getSimpleName());

    public static void main(String[] args) {
        log.info("Starting Producer");

        // Create Producer properties
        Properties properties = new Properties();

        // Connect to Localhost
        properties.setProperty("bootstrap.servers", "localhost:9092");
        properties.setProperty("security.protocol", "PLAINTEXT");

        // Set serializers, serialize key and value into bytes before sending message to Kafka
        properties.setProperty("key.serializer", StringSerializer.class.getName());
        properties.setProperty("value.serializer", StringSerializer.class.getName());

        // Create Producer
        // <String, String> means key is type string and value is type string, will be serialized by StringSerializer
        KafkaProducer<String, String> producer = new KafkaProducer<>(properties);

        // Create a Producer record
        ProducerRecord<String, String> record =
                new ProducerRecord<>("demo_java", "Producing 1 Message");

        // Send data to producer
        producer.send(record);

        // Tell the producer to send all data and block until done -- synchronous
        producer.flush();

        // Flush and close the producer
        producer.close();
    }
}
