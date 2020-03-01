package mokapos

import "github.com/gin-gonic/gin"

type TransactionsHandlerInterface interface {
	GetLatestTransactions(ctx *gin.Context)
}

type TransactionsHandler struct {
}
