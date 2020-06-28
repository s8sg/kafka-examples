package main

import (
	"bufio"
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"os"
	"strings"
)

// Main functiontest
func main() {
	kafkaBrokers := []string{"localhost:9092"}
	brokersEnv := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if brokersEnv != "" {
		kafkaBrokers = strings.Split(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"), ",")
	}

	// make a writer that produces to topic-A, using the least-bytes distribution
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  kafkaBrokers,
		Topic:    "my_topic",
		Balancer: &kafka.LeastBytes{},
	})

	reader := bufio.NewReader(os.Stdin)

	for i := 1; ; i++ {

		fmt.Print("> ")
		msg, _ := reader.ReadString('\n')

		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("%d", i)),
				Value: []byte(msg),
			},
		)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	w.Close()
}
