package balance

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlInter interface {
	PostBalance(ctx *gin.Context)
	GetBalance(ctx *gin.Context)
	GetHistory(ctx *gin.Context)
	PostTransaction(ctx *gin.Context)
}

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
    ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
			"messsage":  "request doesn’t pass validation",
		})
    return
  }

	status, err := c.service.AuthoBalance(req.Currency, c.DB, ctx)
	if err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
			"message": "AuthoBalance",
		})
    return
  }

	if status{
		if err := c.service.UpdateBalance(req, c.DB, ctx); err!= nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
				"message": "UpdateBalance",
			})
			return
		}
	} else {
		if err := c.service.AddBalance(req, c.DB, ctx); err!= nil {
      ctx.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
				"message": "AddBalance",
			})
      return
    }
	}

	if err := c.service.AddLogTransaction(req, c.DB, ctx); err!= nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
			"message": "Log fail",
		})
		return
	}
  ctx.JSON(200, gin.H{"message": "successfully add balance"})
}

func (c *CtrlImpl) GetBalance(ctx *gin.Context)  {
	
  if err := c.service.CheckBalance(c.DB, ctx); err!= nil {
    ctx.AbortWithStatusJSON(404, gin.H{
			"error": err.Error(),
		  "message": "No balance found",
    })
    return
  }

	res, err := c.service.GetBalance(c.DB, ctx)
	if err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		  "message": "GetBalance",
    })
    return
  }
	
	ctx.JSON(200, gin.H{
		"message": "success",
	  "data": res,
  })
}

func (c *CtrlImpl) PostTransaction(ctx *gin.Context) {
	var req ReqTransaction
  if err := ctx.ShouldBindJSON(&req); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := c.validate.Struct(req); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
	status, err := c.service.AuthoBalance(req.FromCurrency, c.DB, ctx)
	if err!= nil {
    ctx.AbortWithStatusJSON(404, gin.H{
			"error": err.Error(),
			"message": "AuthoBalance",
		})
    return
  }
	if !status {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "no balance",
		})
		return
	}

  if err := c.service.AddTransaction(req, c.DB, ctx); err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  ctx.JSON(200, gin.H{"message": "Transaction created successfully"})
}

func (c *CtrlImpl) GetHistory(ctx *gin.Context)  {
	param, err := getParam(ctx)
	if err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		  "message": "Invalid params",
    })
    return
  }
	res,meta, err := c.service.GetLogBalance(param, c.DB, ctx)
	if err!= nil {
    ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		  "message": "Get history",
    })
    return
  }
	
  ctx.JSON(200, gin.H{
		"message": "successfully get balance history",
		"data": res,
		"meta": meta,
  })
}

func getParam(ctx *gin.Context) (Params, error)  {
	limit := ctx.DefaultQuery("limit", "5")
	offset := ctx.DefaultQuery("offset", "0")
	limitInt, err := strconv.ParseUint(limit, 10 , 16)
	if err != nil {
		return Params{}, err
	}
	offsetInt, err := strconv.ParseUint(offset, 10, 16)
	if err != nil {
		return Params{}, err
	}

	param := Params{
		Limit:  uint16(limitInt),
    Offset: uint16(offsetInt),
	}
	return param, nil
}



