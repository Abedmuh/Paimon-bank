package users

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserCtrlInter interface {
	PostUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type UserSvcInter interface {
	CreateUser(req ReqUserReg, tx *sql.DB, ctx *gin.Context) (ResUser, error)
	LoginUser(useDb User,req ReqUserLog , tx *sql.DB, ctx *gin.Context) (ResUser, error)

	// checking
	CheckUserReg(req string, tx *sql.DB, ctx *gin.Context) error
	CheckUserLog(req string, tx *sql.DB, ctx *gin.Context) (User, error)

}

