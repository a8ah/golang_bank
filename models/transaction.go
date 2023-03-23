package models

import "github.com/google/uuid"

// Client structure
type Transaction struct {
	BaseModel          BaseModel `gorm:"embedded"`
	OriginAccount      uuid.UUID `gorm:"type:uuid REFERENCES accounts(uuid)" json:"origin_accounts_uuid" `
	DestinationAccount uuid.UUID `gorm:"type:uuid REFERENCES accounts(uuid)" json:"destianation_accounts_uuid" `
	Amount             float32   `gorm:"not null" json:"amount"`
	CurrencyUUID       uuid.UUID `gorm:"type:uuid REFERENCES currencies(uuid)" json:"currency_uuid" `
}

func (t *Transaction) New_origin_account(newOriginAccount uuid.UUID) {
	t.OriginAccount = newOriginAccount
}

func (t *Transaction) New_destination_account(newDestinationAccount uuid.UUID) {
	t.DestinationAccount = newDestinationAccount
}

func (t *Transaction) New_amount(newAmount float32) {
	t.Amount = newAmount
}

func (t *Transaction) New_currency(newCurrency uuid.UUID) {
	t.CurrencyUUID = newCurrency
}
