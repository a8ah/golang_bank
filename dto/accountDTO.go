package dto

import "github.com/google/uuid"

// AccountcreateDTO structure
type AccountCreateDTO struct {
	ClientUUID   uuid.UUID `json:"client_uuid"`
	CurrencyUUID uuid.UUID `json:"currency_uuid"`
	Number       string    `json:"account_number"`
}
