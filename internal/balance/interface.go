package balance

import (
	"github.com/gin-gonic/gin"
)

type CtrlInter interface {
	PostBalance(ctx *gin.Context)
	GetBalance(ctx *gin.Context)
	GetHistory(ctx *gin.Context)
}