package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	routerGroup := r.Group("/api/v1")

	transactionRoute := NewTransactionRoutes(routerGroup)
	transactionRoute.RegisterRoutes()

	return r
}
