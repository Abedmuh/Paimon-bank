package balance

// main
type Balance struct {
	Id        string `json:"id"`
	BankOwner string `json:"bank_owner"`
	AccNumber string `json:"acc_number"`
	Name      string `json:"name"`
	Balance   int64  `json:"balance"`
	Currency  string `json:"currency"`
}

// response
type Resbalance struct {
	Balance  uint64 `json:"balance"`
	Currency string `json:"currency"`
}

// req
type ReqTransaction struct {
	RecipientBankAccountNumber string `json:"recipientBankAccountNumber" validate:"required"`
	RecipientBankName          string `json:"recipientBankName" validate:"required"`
	FromCurrency               string `json:"fromCurrency" validate:"required"`
	Balance                    uint64 `json:"balances" validate:"required"`
}

type Reqbalance struct {
	SenderBankAccountNumber string `json:"senderBankAccountNumber"`
	SenderBankName          string `json:"senderBankName"`
	AddedBalance            uint64 `json:"addedBalance"`
	Currency                string `json:"currency"`
	TransferProofImg        string `json:"transferProofImg"`
}

// main
type Transaction struct {
	TransactionId    string `json:"transactionId"`
	Balance          int64  `json:"balance"`
	Currency         string `json:"currency"`
	TransferProofImg string `json:"transferProofImg"`
	CreatedAt        string `json:"createdAt"`
	Source           source
}

type source struct {
	BankAccountNumber string `json:"bankAccountNumber"`
	BankName          string `json:"bankName"`
}

type Params struct {
	Limit uint16 `json:"limit"`
	Offset uint16 `json:"offset"`
	Total uint16 `json:"total"`
}
