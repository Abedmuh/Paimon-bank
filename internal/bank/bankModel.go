package bank

type Bank struct {
	Id        string `json:"id"`
	Owner     string `json:"owner"`
	Name      string `json:"name"`
	AccNumber string `json:"acc_number"`
}