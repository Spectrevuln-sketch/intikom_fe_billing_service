package kafkaConsumer

import (
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	generalTransactionConsumer *kafka.Reader
}

func NewConsumer() *Consumer {
	return &Consumer{
		generalTransactionConsumer: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "general-transactions",
			GroupID: "trx-worker",
		}),
	}
}
