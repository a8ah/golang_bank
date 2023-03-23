package models

import (
	"gorm.io/gorm"
)

// Client structure
type Account struct {
	gorm.Model
	ClientID uint
	Number   int     `gorm:"not null" json:"number"`
	Balance  float32 `gorm:"not null" json:"balance"`
	Limit    int     `gorm:"not null" json:"limit" default: 10000`
}

// // Updating data in same transaction
// func (u *Account) AfterUpdate(tx *gorm.DB) (err error) {
// 	u.UpdatedAt = time.Now()
// 	return
// }
