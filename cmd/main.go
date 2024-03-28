package main

import (
	"github.com/Abedmuh/Paimon-bank/internal/routes"
	"github.com/Abedmuh/Paimon-bank/pkg/dbconfig"
	"github.com/Abedmuh/Paimon-bank/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	db, err := dbconfig.GetDBConnection()
	if err!= nil {
    panic(err)
  }
	defer db.Close()


	app := gin.Default()
	app.Use(middleware.RecoveryMiddleware())
	app.MaxMultipartMemory = 2 << 20 

	validate:= validator.New()


	v1 := app.Group("v1")
	{
		routes.UserRoutes(v1, db, validate)
		routes.BalanceRoutes(v1, db, validate)
	}

	app.Run(":8080")
}