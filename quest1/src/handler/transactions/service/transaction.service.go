package service

import (
	"billing-service/src/core/config"
	kafkaProducer "billing-service/src/core/worker/producer"
	"billing-service/src/handler/transactions/model"
	"billing-service/src/handler/transactions/repository"
	"errors"
)

type transactionService struct {
	repo     *repository.TransactionRepository
	producer *kafkaProducer.Producer
}

func NewTransactionService() TransactionService {
	return &transactionService{
		repo:     repository.NewTransactionRepository(config.DB),
		producer: kafkaProducer.NewProducer(),
	}
}

func (s *transactionService) CreateTransaction(req model.TransactionRequest) (string, error) {
	if req.Amount <= 0 {
		return "", errors.New("amount must be > 0")
	}
	err := s.producer.PublishTransaction(req)
	if err != nil {
		return "", errors.New("failed publish")
	}

	return "accepted", nil
}

func (s *transactionService) ProcessTransaction(trx model.TransactionRequest) error {
	return s.repo.ProcessTransaction(trx)
}

func (s *transactionService) GetTransactionSender(userId string, page int, limit int) (map[string]interface{}, error) {
	total, balance, data, err := s.repo.GetTransactionSender(userId, page, limit)
	if err != nil {
		return nil, err
	}
	totalPage := (total + limit - 1) / limit

	return map[string]interface{}{
		"total":   totalPage,
		"page":    page,
		"limit":   limit,
		"balance": balance,
		"data":    data,
	}, nil

}
