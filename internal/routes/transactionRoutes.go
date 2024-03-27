package routes

import (
	"database/sql"

	"github.com/Abedmuh/Paimon-bank/internal/transaction"
	"github.com/Abedmuh/Paimon-bank/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TransactionRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {
	service := transaction.NewTransactionService()
	controler := transaction.NewTransactionController(service, db, validate)

	path := route.Group("/Transaction")
	path.Use(middleware.Authentication())
	{
		path.POST("/", controler.PostTransaction)
	}
}

