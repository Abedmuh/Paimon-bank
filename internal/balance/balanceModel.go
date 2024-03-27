package balance

type Reqbalance struct {
	SenderBankAccountNumber string `json:"senderBankAccountNumber"`
	SenderBankName          string `json:"senderBankName"`
	AddedBalance            uint64 `json:"addedBalance"`
	Currency                string `json:"currency"`
	TransferProofImg        string `json:"transferProofImg"`
}

// main
type Balance struct {
	Id        string `json:"id"`
	BankOwner string `json:"bank_owner"`
	Balance   uint64 `json:"balance"`
	Currency  string `json:"currency"`
}

// response
type Resbalance struct {
	Balance  uint64 `json:"balance"`
	Currency string `json:"currency"`
}