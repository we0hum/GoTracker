package queue

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	go func() {
		broker := os.Getenv("KAFKA_BROKER")
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{broker},
			Topic:   "orders",
			GroupID: "gotracker-group",
		})
		defer r.Close()

		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				fmt.Println("Kafka error:", err)
				continue
			}
			fmt.Printf("[Kafka] Новый заказ: %s\n", string(m.Value))
		}
	}()
}
