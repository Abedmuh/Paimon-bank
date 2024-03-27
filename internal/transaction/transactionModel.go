package transaction

type ReqTransaction struct {
	RecipientBankAccountNumber string `json:"recipientBankAccountNumber"`
	RecipientBankName          string `json:"recipientBankName"`
	FromCurrency               string `json:"fromCurrency"`
	Balance                    uint64 `json:"balance"`
}
