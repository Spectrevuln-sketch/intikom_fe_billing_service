package service

import (
	"billing-service/src/handler/transactions/model"
)

type TransactionService interface {
	CreateTransaction(req model.TransactionRequest) (string, error)
	ProcessTransaction(trx model.TransactionRequest) error
	GetTransactionSender(userId string, page int, limit int) (map[string]interface{}, error)
}
