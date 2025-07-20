mac_kafka:
	brew install kafka

docker_kafka:
	git clone https://github.com/conduktor/kafka-stack-docker-compose.git
	ls -l
	docker-compose -f ./kafka-stack-docker-compose/zk-single-kafka-single.yml up -d

	@echo "Ensure both services (Kafka and Zookeeper) are running"
	docker-compose -f ./kafka-stack-docker-compose/zk-single-kafka-single.yml ps

	@echo "Running commands against Kafka running on Docker"
	docker exec -it kafka1 /bin/bash

	## Run this command inside the docker container to create the config file
	# echo "bootstrap.servers=localhost:9092
	# security.protocol=PLAINTEXT
	# auto.offset.reset=earliest" > playground.config

docker_kafka_stop:
	docker-compose -f ./kafka-stack-docker-compose/zk-single-kafka-single.yml stop
	docker-compose -f ./kafka-stack-docker-compose/zk-single-kafka-single.yml down
	rm -rf kafka-stack-docker-compose
