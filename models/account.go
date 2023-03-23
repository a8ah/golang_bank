package models

import "github.com/google/uuid"

// Client structure
type Account struct {
	BaseModel    BaseModel `gorm:"embedded"`
	Number       int       `gorm:"not null" json:"number"`
	Balance      float32   `gorm:"not null" json:"balance"`
	Limit        int       `gorm:"not null; default: 10000" json:"limit"`
	ClientUUID   uuid.UUID `gorm:"type:uuid REFERENCES clients(uuid)" json:"client_uuid" `
	CurrencyUUID uuid.UUID `gorm:"type:uuid REFERENCES currencies(uuid)" json:"currency_uuid" `
}
