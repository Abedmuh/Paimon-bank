package users

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCtrlImpl struct {
	service UserSvcInter
	DB      *sql.DB
	validate *validator.Validate
}

func NewUserController( service UserSvcInter,DB *sql.DB, validate *validator.Validate) UserCtrlInter{
	return &UserCtrlImpl{
		service: service,
    DB:      DB,
    validate: validate,
	}
}

func (uc *UserCtrlImpl) PostUser(ctx *gin.Context) {
	var req ReqUserReg
  if err := ctx.ShouldBindJSON(&req); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := uc.validate.Struct(req); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
			"message": "Fail to validate",
		})
    return
  }

	if err := uc.service.CheckUserReg(req.Email,uc.DB, ctx); err!= nil {
    ctx.AbortWithStatusJSON(409, gin.H{
			"error": err.Error(),
			"message": "Fail to check",
		})
    return
  }

	res,err := uc.service.CreateUser(req, uc.DB, ctx)
  if  err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
			"message": "Fail to create",
		})
    return
  }
  ctx.JSON(201, gin.H{
		"message": "User created successfully",
		"data": res,
	})
}

func (uc *UserCtrlImpl) LoginUser(ctx *gin.Context) {
	var req ReqUserLog
  if err := ctx.ShouldBindJSON(&req); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := uc.validate.Struct(req); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

	user,err := uc.service.CheckUserLog(req.Email, uc.DB, ctx)
	if  err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

	res,err := uc.service.LoginUser(user, req, uc.DB, ctx)
  if err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  ctx.JSON(200, gin.H{
		"message": "User logged in successfully",
	  "data": res,
  })
	
}