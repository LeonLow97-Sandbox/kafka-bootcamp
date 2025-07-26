package io.conduktor.kafka;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.common.errors.WakeupException;
import org.apache.kafka.common.serialization.StringDeserializer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Arrays;
import java.util.Properties;

public class ConsumerWithShutdown {

    private static final Logger log = LoggerFactory.getLogger(ConsumerWithShutdown.class.getSimpleName());

    public static void main(String[] args) {
        log.info("Starting Consumer");

        String topic = "demo_java";
        String groupId = "my-java-consumer-group";

        // Create Producer properties
        Properties properties = new Properties();

        // Connect to Localhost
        properties.setProperty("bootstrap.servers", "localhost:9092");
        properties.setProperty("security.protocol", "PLAINTEXT");

        // Consumer configuration, consumer deserializes while producer serializers
        properties.setProperty("key.deserializer", StringDeserializer.class.getName());
        properties.setProperty("value.deserializer", StringDeserializer.class.getName());

        // Consumer Group ID
        properties.setProperty("group.id", groupId);

        // Reset offset
        // 3 values - none, earliest, latest
        // none: no existing consumer group, fail
        // earliest: read from beginning of topic, like `--from-beginning`
        // latest: only read from new messages sent from now
        properties.setProperty("auto.offset.reset", "earliest");

        // Create Consumer
        KafkaConsumer<String, String> consumer = new KafkaConsumer<>(properties);

        // Get a reference to the main thread
        final Thread mainThread = Thread.currentThread();

        // Add the shutdown hook
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            log.info("Detected a shutdown, let's exit by calling consumer.wakeup()");
            consumer.wakeup(); // throws Wakeup Exception at consumer.poll() in while loop

            // Join the main thread to allow the execution of the code in the main thread
            try {
                mainThread.join();
            } catch (InterruptedException e) {}
        }));

        try {
            // Subscribe to a topic
            consumer.subscribe(Arrays.asList(topic));

            // Poll for data
            while (true) {
                log.info("Polling...");

                // consumer.poll throws wakeup exception
                ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(1000)); // poll every 1 second

                for (ConsumerRecord<String, String> record : records) {
                    log.info("Key: " + record.key() + ", Value: " + record.value());
                    log.info("Partition: " + record.partition() + ", Offset: " + record.offset());
                }
            }
        } catch (WakeupException e) {
            log.info("Consumer is starting to shut down");
        } catch (Exception e) {
            log.error("Unexpected exception in the consumer", e);
        } finally {
            consumer.close(); // close the consumer, this will also commit offsets
            log.info("The consumer is now gracefully shut down");
        }
    }
}
