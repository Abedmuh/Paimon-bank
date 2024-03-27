package transaction

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlInter interface {
	PostTransaction(ctx *gin.Context)
}

type CtrlImpl struct {
	service SvcInter
	DB      *sql.DB
	validate *validator.Validate
}

func NewTransactionController( service SvcInter,DB *sql.DB, validate *validator.Validate) CtrlInter{
	return &CtrlImpl{
		service: service,
    DB:      DB,
    validate: validate,
	}
}

func (uc *CtrlImpl) PostTransaction(ctx *gin.Context) {
	var req ReqTransaction
  if err := ctx.ShouldBindJSON(&req); err!= nil {
    ctx.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := uc.validate.Struct(req); err!= nil {
    ctx.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := uc.service.AddTransaction(req, uc.DB, ctx); err!= nil {
    ctx.JSON(400, gin.H{"error": err.Error()})
    return
  }
  ctx.JSON(200, gin.H{"message": "Transaction created successfully"})
}