package users

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Abedmuh/Paimon-bank/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserSvcImpl struct {
}

func NewUserService() UserSvcInter {
	return &UserSvcImpl{}
}

func (us *UserSvcImpl) CreateUser(req ReqUserReg, tx *sql.DB, ctx *gin.Context) (ResUser,error) {
	id :=  uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err!= nil {
    return ResUser{}, err
  }

	var user User
	query := `INSERT INTO users (id, name, email, password )
		VALUES ($1, $2, $3, $4) 
		RETURNING id, name, email
	`
	err = tx.QueryRow(query, 
		id, 
		req.Name, 
		req.Email, 
		hashedPassword).Scan(
			&user.Id,
			&user.Email, 
			&user.Name)
	if err!= nil {
    return ResUser{}, err
  }
	token,_ := utils.GenerateToken(user.Id)
	res := ResUser{
    Name: user.Name,
    Email: user.Email,
    AccessToken: token,
	}

	return res, nil
}

func (us *UserSvcImpl) LoginUser(userDb User, req ReqUserLog,tx *sql.DB, ctx *gin.Context) (ResUser,error) {
	fmt.Println(req.Password)
	fmt.Println(userDb.Password)
	err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(req.Password))
	if err != nil {
		return ResUser{},errors.New("password salah")
	}

	token, err := utils.GenerateToken(userDb.Id)
	if err!= nil {
    return ResUser{}, err
  }

	resLog := ResUser{
		Email: userDb.Email,
    Name: userDb.Name,
    AccessToken: token,
	}
  return resLog, nil
}

func (uc *UserSvcImpl) CheckUserReg(req string, tx *sql.DB, ctx *gin.Context) error {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := tx.QueryRow(query, req).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	return errors.New("user already exists")
}

func (uc *UserSvcImpl) CheckUserLog(req string, tx *sql.DB, ctx *gin.Context) (User, error) {
	var userDb User
  query := `SELECT * FROM users WHERE email = $1`
  err := tx.QueryRow(query, req).Scan(
    &userDb.Id, 
    &userDb.Name, 
    &userDb.Email,
    &userDb.Password, 
		&userDb.CreatedAt)
  if err!= nil {
    return userDb, err
  }
  return userDb, nil
}