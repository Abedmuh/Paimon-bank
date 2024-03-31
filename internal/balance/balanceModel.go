package balance

// main db
type Balance struct {
	Id        string `json:"id"`
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
	RecipientBankAccountNumber string `json:"recipientBankAccountNumber" validate:"required,min=5,max=30"`
	RecipientBankName          string `json:"recipientBankName" validate:"required,min=5,max=30"`
	FromCurrency               string `json:"fromCurrency" validate:"required,iso4217,min=1,max=3"`
	Balance                    int64 `json:"balances" validate:"required"`
}

type Reqbalance struct {
	SenderBankAccountNumber string `json:"senderBankAccountNumber" validate:"required,min=5,max=30"`
	SenderBankName          string `json:"senderBankName" validate:"required,min=5,max=30"`
	AddedBalance            int64 `json:"addedBalance" validate:"required,min=1"`
	Currency                string `json:"currency" validate:"required,iso4217,min=1,max=3"`
	TransferProofImg        string `json:"transferProofImg" validate:"required,myUrl"`
}

// main
type Transaction struct {
	TransactionId    string `json:"transactionId"`
	Balance          int64  `json:"balance"`
	Currency         string `json:"currency"`
	TransferProofImg *string `json:"transferProofImg"`
	CreatedAt        string `json:"createdAt"`
	Source           source `json:"source"`
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
