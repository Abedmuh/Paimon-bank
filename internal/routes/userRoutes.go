package routes

import (
	"database/sql"

	"github.com/Abedmuh/Paimon-bank/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {
	service := users.NewUserService()
	controler := users.NewUserController(service, db, validate)

	path := route.Group("/user")
	{
		path.POST("/register", controler.PostUser)
		path.POST("/login", controler.LoginUser)
	}
}

