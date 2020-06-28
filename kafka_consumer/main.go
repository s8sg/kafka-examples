package main

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"os"
	"strings"
	"sync"
)

// Consumer
func kafkaConsumer(consumer int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting consumer %d\n", consumer)

	kafkaBrokers := []string{"localhost:9092"}
	brokersEnv := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if brokersEnv != "" {
		kafkaBrokers = strings.Split(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"), ",")
	}

	groupId := "load_balance"
	if os.Getenv("BROADCAST") == "true" {
		groupId = fmt.Sprintf("load_balance_%d", consumer)
	}

	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaBrokers,
		GroupID:  groupId,
		Topic:    "my_topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("message received by %d at topic/partition/offset %v/%v/%v: %s = %s\n", consumer, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	fmt.Printf("Stopping consumer %d\n", consumer)

	r.Close()
}

// Main function
func main() {

	total := 3

	var wg sync.WaitGroup
	wg.Add(total)

	for i := 1; i <= total; i++ {
		go kafkaConsumer(i, &wg)
	}

	wg.Wait()
}
