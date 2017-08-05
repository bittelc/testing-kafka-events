package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func mainConsumer(partition int32) {
	kafka := newKafkaConsumer()
	defer kafka.Close()

	consumer, err := kafka.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("Kafka errors: %s\n", err)
		os.Exit(-1)
	}

	go consumeEvents(consumer)

	fmt.Println("Press [Enter] to exit consumer\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Terminating")
}
