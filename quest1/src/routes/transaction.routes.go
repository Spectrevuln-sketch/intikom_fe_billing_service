package routes

import (
	"billing-service/src/handler/transactions"
	"billing-service/src/handler/transactions/service"

	"github.com/gin-gonic/gin"
)

type TransactionRoutes struct {
	router *gin.RouterGroup
}

func NewTransactionRoutes(router *gin.RouterGroup) *TransactionRoutes {
	return &TransactionRoutes{
		router: router,
	}
}

func (r *TransactionRoutes) RegisterRoutes() {
	transactionRouteGroup := r.router.Group("/transactions")

	service := service.NewTransactionService()
	handler := transactions.NewHandler(service)

	transactionRouteGroup.POST("", handler.CreateTransaction)
	transactionRouteGroup.GET("accounts/:senderId", handler.GetTransactionSender)
}
