package routes

import (
	"database/sql"

	"github.com/Abedmuh/Paimon-bank/internal/balance"
	"github.com/Abedmuh/Paimon-bank/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BalanceRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {
	service := balance.NewBalanceService()
	controler := balance.NewBalanceController(service, db, validate)

	path := route.Group("/balance")
	path.Use(middleware.Authentication())
	{
		path.POST("/", controler.PostBalance)
		path.GET("/", controler.GetBalance)
		path.GET("/history", controler.GetHistory)
	}
}

