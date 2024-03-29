package balance

import (
	"database/sql"
	"errors"

	"github.com/Abedmuh/Paimon-bank/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SvcInter interface {
	AuthoBalance(Req string, tx *sql.DB, ctx *gin.Context) (bool, error)

	AddBalance(Req Reqbalance, tx *sql.DB, ctx *gin.Context) error
	UpdateBalance(Req Reqbalance, tx *sql.DB, ctx *gin.Context) error

	CheckBalance(tx *sql.DB, ctx *gin.Context) error
	GetBalance(tx *sql.DB, ctx *gin.Context) ([]Resbalance,error)

	AddTransaction(req ReqTransaction, tx *sql.DB, ctx *gin.Context) error

	AddLogBalance(req Reqbalance, tx *sql.DB, ctx *gin.Context) error
	AddLogTransaction(req ReqTransaction, tx *sql.DB, ctx *gin.Context) error


	GetLogBalance(params Params,tx *sql.DB, ctx *gin.Context) ([]Transaction, Params, error)
}

type SvcImpl struct {
}

func NewBalanceService() SvcInter {
	return &SvcImpl{}
}

func (s *SvcImpl) AuthoBalance(req string, tx *sql.DB, ctx *gin.Context) (bool,error) {
	user,_ := ctx.Get("user")
	reqUser, ok := user.(string)
	if !ok {
		return false, errors.New("Unathorized")
	}

	var id string
	query := `SELECT id FROM balances 
		WHERE owner = $1 AND currency = $2
	`
	err := tx.QueryRow(query, 
		reqUser, 
		req).Scan(&id)
	
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

	queryBank := `INSERT INTO balances (id, owner, currency, balance)
	  VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(queryBank, 
		idBank,
		reqUser,
		req.Currency,
		req.AddedBalance)
	if err!= nil {
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
		WHERE owner = $2 AND currency = $3
	`
	_, err = tx.Exec(query,
		req.AddedBalance,
		reqUser,
		req.Currency)
	if err != nil {
		return err
	}

	return nil
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

func (s *SvcImpl) AddTransaction(req ReqTransaction, tx *sql.DB, ctx *gin.Context) error {
	reqUser, err := utils.GetUserID(ctx)
	if err != nil {
    return err
  }

	querycheck := `
		SELECT balance 
		FROM balances
		WHERE owner = $1 AND currency = $2
	`
	var balance uint64
	err = tx.QueryRowContext(ctx, querycheck, 
		reqUser,
		req.FromCurrency).Scan(&balance)
	if err!= nil {
    return err
  }
	if balance < req.Balance {
		return errors.New("out of balance")
	}

	query := `
		UPDATE balances
		SET balance = balance - $1
		WHERE owner = $2 AND currency = $3
	`
	_, err = tx.Exec(query,
		req.Balance,
		reqUser,
		req.FromCurrency)
	if err != nil {
		return err
	}

	return nil
}


func (s *SvcImpl) AddLogBalance(req Reqbalance, tx *sql.DB, ctx *gin.Context) error {
	reqUser, err := utils.GetUserID(ctx)
	if err != nil {
    return err
  }
	id := uuid.New().String()
	queryTransaction := `
		INSERT INTO log_transaction (id, owner, balance, currency, transfer_proof, bank_account, bank_name)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = tx.Exec(queryTransaction,
		id,
		reqUser,
		req.AddedBalance,
		req.Currency,
		req.TransferProofImg,
	  req.SenderBankAccountNumber,
    req.SenderBankName)
	if err != nil {
		return err
	}
	return nil
}
func (s *SvcImpl) AddLogTransaction(req ReqTransaction, tx *sql.DB, ctx *gin.Context) error {
	reqUser, err := utils.GetUserID(ctx)
	if err != nil {
    return err
  }
	id := uuid.New().String()
	queryTransaction := `
		INSERT INTO log_transaction (id, owner, balance, currency, bank_account, bank_name)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = tx.Exec(queryTransaction,
		id,
		reqUser,
		-req.Balance,
		req.FromCurrency,
	  req.RecipientBankAccountNumber,
    req.RecipientBankName)
	if err != nil {
		return err
	}
	return nil
}

func (s *SvcImpl) GetLogBalance(param Params, tx *sql.DB, ctx *gin.Context) ([]Transaction, Params, error) {
	var transactions []Transaction
  reqUser, err := utils.GetUserID(ctx)
  if err != nil {
		return nil, Params{}, err
  }

	var totalRows uint16
	countQuery := `SELECT COUNT(*) FROM log_transaction WHERE owner = $1`
	err = tx.QueryRowContext(ctx, countQuery, reqUser).Scan(&totalRows)
	if err != nil {
			return nil, Params{}, err
	}

	params := Params{
		Limit: param.Limit,
    Offset:  param.Offset,
		Total: totalRows,
	}

  query := `SELECT id, balance, currency, transfer_proof, created_at, bank_account, bank_name
    FROM log_transaction
    WHERE owner = $1
		LIMIT $2 OFFSET $3
  `
  rows, err := tx.QueryContext(ctx, query, 
		reqUser, 
		param.Limit, 
		param.Offset)
  if err != nil {
    return nil, Params{}, err
  }
  defer rows.Close()
	for rows.Next() {
		var transaction Transaction
    err := rows.Scan(
			&transaction.TransactionId, 
			&transaction.Balance, 
			&transaction.Currency, 
			&transaction.TransferProofImg,
			&transaction.CreatedAt, 
			&transaction.Source.BankAccountNumber, 
			&transaction.Source.BankName)
    if err != nil {
      return nil, Params{}, err
    }
    transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
    return nil, Params{}, err
  }
	return transactions, params, nil
}