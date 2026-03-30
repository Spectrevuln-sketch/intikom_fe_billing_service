package transactions

import (
	"billing-service/src/handler/transactions/model"
	"billing-service/src/handler/transactions/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service service.TransactionService
}

func NewHandler(service service.TransactionService) *transactionHandler {
	return &transactionHandler{service: service}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var req model.TransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	result, err := h.service.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}

func (h *transactionHandler) GetTransactionSender(c *gin.Context) {
	senderId := c.Param("senderId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if senderId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "sender id is required",
		})
	}
	result, err := h.service.GetTransactionSender(senderId, page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"senderId": senderId,
		"data":     result,
	})
}
