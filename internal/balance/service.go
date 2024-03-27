package balance

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/Abedmuh/Paimon-bank/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SvcInter interface {
	AuthoBalance(Req Reqbalance, tx *sql.DB, ctx *gin.Context) (bool, error)

	AddBalance(Req Reqbalance, tx *sql.DB, ctx *gin.Context) error
	UpdateBalance(Req Reqbalance, tx *sql.DB, ctx *gin.Context) error

	CheckBalance(tx *sql.DB, ctx *gin.Context) error
	GetBalance(tx *sql.DB, ctx *gin.Context) ([]Resbalance,error)
}

type SvcImpl struct {
}

func NewBalanceService() SvcInter {
	return &SvcImpl{}
}

func (s *SvcImpl) AuthoBalance(req Reqbalance, tx *sql.DB, ctx *gin.Context) (bool,error) {
	user,_ := ctx.Get("user")
	reqUser, ok := user.(string)
	if !ok {
		return false, errors.New("Unathorized")
	}

	var id string
	query := `SELECT id FROM banks 
		WHERE owner = $1 AND acc_number = $2 AND bank_name = $3
	`
	err := tx.QueryRow(query, 
		reqUser, 
		req.SenderBankAccountNumber,
		req.SenderBankName).Scan(&id)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *SvcImpl) AddBalance(req Reqbalance, tx *sql.DB, ctx *gin.Context) error {
	reqUser, err := utils.GetUserID(ctx)
	if err != nil {
		return err
	}
	idBank := uuid.New().String()
	idBalance := uuid.New().String()

	queryBank := `INSERT INTO banks (id, owner, name, acc_number)
	  VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(queryBank, 
		idBank,
		reqUser,
		req.SenderBankName,
	  req.SenderBankAccountNumber)
	if err!= nil {
    return err
  }

	queryBalance := `INSERT INTO balances (id, bank_owner, currency, balance
		) VALUES ($1, $2, $3, $4)`
  _, err = tx.Exec(queryBalance,
		idBalance,
    idBank,
    req.Currency,
    req.AddedBalance)
	if err!= nil {
    return err
  }

	// Tambahkan transaksi baru ke tabel transactions
	err = s.addTransaction(req, tx, reqUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *SvcImpl) UpdateBalance(req Reqbalance, tx *sql.DB, ctx *gin.Context) error {
	reqUser, err := utils.GetUserID(ctx)
	if err != nil {
		return err
	}

	query := `
		UPDATE balances
		SET balance = balance + $1
		WHERE bank_owner = (
			SELECT id
			FROM banks
			WHERE owner = $2 AND acc_number = $3 AND name = $4
		)
		AND currency = $5
	`
	_, err = tx.Exec(query,
		req.AddedBalance,
		reqUser,
		req.SenderBankAccountNumber,
		req.SenderBankName,
		req.Currency)
	if err != nil {
		return err
	}

	// Tambahkan transaksi baru ke tabel transactions
	err = s.addTransaction(req, tx, reqUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *SvcImpl) addTransaction(req Reqbalance, tx *sql.DB, reqUser string) error {
	id := uuid.New().String()
	queryTransaction := `
		INSERT INTO transactions (id, owner, currency, balance, transferProofImg)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := tx.Exec(queryTransaction,
		id,
		reqUser,
		req.Currency,
		strconv.Itoa(int(req.AddedBalance)),
		req.TransferProofImg)
	if err != nil {
		return err
	}
	return nil
}

func (s *SvcImpl) GetBalance(tx *sql.DB, ctx *gin.Context) ([]Resbalance, error) {
	var balances []Resbalance
	reqUser, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	query := `SELECT b.currency, b.balance 
		FROM balances b 
		WHERE owner = $1
	`
	rows, err := tx.QueryContext(ctx, query, reqUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var balance Resbalance
		err := rows.Scan(&balance.Currency, &balance.Balance)
		if err != nil {
			return nil, err
		}
		balances = append(balances, balance)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return balances, nil
}

func (s *SvcImpl) CheckBalance(tx *sql.DB, ctx *gin.Context) error {
	reqUser, err := utils.GetUserID(ctx)
	if err != nil {
		return err
	}

	query := `SELECT EXISTS (SELECT 1 FROM balances WHERE owner = $1)`

	var exists bool
	err = tx.QueryRowContext(ctx, query, reqUser).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("no balance found")
	}

	return nil
}
