package kafkaProducer

import (
	"billing-service/src/handler/transactions/model"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

func (p *Producer) PublishTransaction(req model.TransactionRequest) error {
	msg, _ := json.Marshal(req)

	return p.generalTransactionOut.WriteMessages(context.Background(),
		kafka.Message{Value: msg},
	)
}
