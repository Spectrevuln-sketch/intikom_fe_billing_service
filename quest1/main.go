package main

import (
	"billing-service/src/core/config"
	kafkaConsumer "billing-service/src/core/worker/consumer"
	"billing-service/src/routes"
)

func main() {
	config.InitDB()

	go startKafkaConsumer()

	router := routes.SetupRouter()
	router.Run(":8080")
}

func startKafkaConsumer() {
	kafkaConsumer.NewConsumer().StartGeneralTransactionConsumer()
}
