package cmd

import (
	"github.com/Abedmuh/Paimon-bank/internal/routes"
	"github.com/Abedmuh/Paimon-bank/pkg/db"
	"github.com/Abedmuh/Paimon-bank/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	db, err := db.GetDBConnection()
	if err!= nil {
    panic(err)
  }
	defer db.Close()


	app := gin.Default()
	app.Use(middleware.RecoveryMiddleware())
	validate:= validator.New()


	v1 := app.Group("v1")
	{
		routes.UserRoutes(v1, db, validate)
		routes.BalanceRoutes(v1, db, validate)
		routes.TransactionRoutes(v1, db, validate)
	}
}