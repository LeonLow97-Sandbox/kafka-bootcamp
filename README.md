# Course Links

|                                                         Course Name                                                          |   Instructor    |                        Course Link                         |
| :--------------------------------------------------------------------------------------------------------------------------: | :-------------: | :--------------------------------------------------------: |
|                                         Udemy - Learn Apache Kafka for Beginners v3                                          | Stephane Maarek |     [Link](https://www.udemy.com/course/apache-kafka/)     |
|                                             YouTube - Apache Kafka in 6 Minutes                                              |  James Cutajar  |    [Link](https://www.youtube.com/watch?v=Ch5VhJzaoaI)     |
| YouTube - High Available Microservices with Apache Kafka in Golang (uses "github.com/confluentinc/confluent-kafka-go/kafka") |   Anthony GG    | [Link](https://www.youtube.com/watch?v=-yVxChp7HoQ&t=987s) |

## Apache Kafka with Go

- Use `https://github.com/segmentio/kafka-go` for better stability.
- "github.com/confluentinc/confluent-kafka-go/kafka" introduces cross-compilation issues with C Library (Apache Kafka), setting CGO_ENABLED=1 allows Go to compile properly with C code but it still introduced complications when working with different Operating Systems or architectures.