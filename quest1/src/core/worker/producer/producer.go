package kafkaProducer

import "github.com/segmentio/kafka-go"

type Producer struct {
	generalTransactionOut *kafka.Writer
}

func NewProducer() *Producer {
	return &Producer{
		generalTransactionOut: kafka.NewWriter(kafka.WriterConfig{
			Brokers:      []string{"localhost:9092"},
			Topic:        "general-transactions",
			WriteTimeout: 100000,
		}),
	}
}
