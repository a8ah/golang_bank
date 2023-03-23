package dto

// TransactionDTO structure
type TransactionDTO struct {
	Origin       string  `json:"origin_accounts"`
	Destination  string  `json:"destination_accounts"`
	Amount       float32 `json:"amount"`
	CurrencyUUID string  `json:"currency"`
}
