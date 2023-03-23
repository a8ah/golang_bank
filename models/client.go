package models

import (
	"gorm.io/gorm"
)

// Client structure
type Client struct {
	gorm.Model
	Name     string `gorm:"text; not null" json:"name"`
	Lastname string `gorm:"text; not null" json:"last_name"`
	Dni      string `gorm:"text; not null" json:"dni"`
	Account  []Account
}

// Updating data in same transaction
// func (u *Client) AfterUpdate(tx *gorm.DB) (err error) {
// 	u.UpdatedAt = time.Now()
// 	return
// }
