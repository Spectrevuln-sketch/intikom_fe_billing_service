package config

import "github.com/segmentio/kafka-go"

func NewKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "event-general-transactions",
		WriteTimeout: 10000,
	}
}
