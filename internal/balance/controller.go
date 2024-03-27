package balance

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlImpl struct {
	service  SvcInter
	DB       *sql.DB
	validate *validator.Validate
}

func NewBalanceController(service SvcInter, DB *sql.DB, validate *validator.Validate) CtrlInter {
	return &CtrlImpl{
		service:  service,
		DB:       DB,
		validate: validate,
	}
}

func (c *CtrlImpl) PostBalance(ctx *gin.Context)  {
	var req Reqbalance
	if err := ctx.ShouldBindJSON(&req); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

	if err := c.validate.Struct(req); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

	status, err := c.service.AuthoBalance(req, c.DB, ctx)
	if err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

	if status{
		if err := c.service.AddBalance(req, c.DB, ctx); err!= nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.service.UpdateBalance(req, c.DB, ctx); err!= nil {
      ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
      return
    }
	}
  ctx.JSON(200, gin.H{"status": status})
}

func (c *CtrlImpl) GetBalance(ctx *gin.Context)  {
	
  if err := c.service.CheckBalance(c.DB, ctx); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

	res, err := c.service.GetBalance(c.DB, ctx)
	if err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
	
	ctx.JSON(200, gin.H{
		"message": "OK",
	  "data": res,
  })
}

func (c *CtrlImpl) GetHistory(ctx *gin.Context)  {
  ctx.JSON(200, gin.H{
		"message": "OK",

  })
}

