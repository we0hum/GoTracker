package queue

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(broker string) {
	if broker == "" {
		fmt.Println("KAFKA_BROKER не настроен, Kafka consumer отключён")
		return
	}

	go func() {
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
