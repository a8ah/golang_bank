package models

import "github.com/google/uuid"

// Client structure
type Account struct {
	BaseModel    BaseModel `gorm:"embedded"`
	Number       uint64    `gorm:"not null; unique" json:"number"`
	Balance      float32   `gorm:"not null; default: 0" json:"balance"`
	Limit        uint64    `gorm:"not null; default: 10000" json:"limit"`
	ClientUUID   uuid.UUID `gorm:"type:uuid REFERENCES clients(uuid)" json:"client_uuid" `
	CurrencyUUID uuid.UUID `gorm:"type:uuid REFERENCES currencies(uuid)" json:"currency_uuid" `
	SecureNumber uint16    `gorm:"not null" json:"secure_number" `
}
