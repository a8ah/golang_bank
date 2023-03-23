package models

// Client structure
type Client struct {
	BaseModel BaseModel `gorm:"embedded"`
	Name      string    `gorm:"text; not null" json:"name"`
	Surname   string    `gorm:"text; not null" json:"surname"`
	Dni       string    `gorm:"text; not null; unique" json:"dni"`
	Account   []Account
}
