package kafkaConsumer

import (
	"billing-service/src/handler/transactions/model"
	"billing-service/src/handler/transactions/service"
	"context"
	"encoding/json"
	"log"
)

func (tc *Consumer) StartGeneralTransactionConsumer() {

	svc := service.NewTransactionService()

	log.Println("worker started...")

	for {
		msg, err := tc.generalTransactionConsumer.ReadMessage(context.Background())
		if err != nil {
			log.Println("error reading message:", err)
			continue
		}

		log.Println("message received:", string(msg.Value))

		var trx model.TransactionRequest
		err = json.Unmarshal(msg.Value, &trx)
		if err != nil {
			log.Println("json error:", err)
			continue
		}

		err = svc.ProcessTransaction(trx)
		if err != nil {
			log.Println("process error:", err)
		} else {
			log.Println("transaction processed:", trx.ReferenceID)
		}
	}
}
