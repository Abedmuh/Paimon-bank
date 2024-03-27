package transaction

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type SvcInter interface {
	AddTransaction(req ReqTransaction, tx *sql.DB, ctx *gin.Context) error
}

type SvcImpl struct {
}

func NewTransactionService() SvcInter {
	return &SvcImpl{}
}

func (s *SvcImpl) AddTransaction(req ReqTransaction, tx *sql.DB, ctx *gin.Context) error {
	
  return nil
}