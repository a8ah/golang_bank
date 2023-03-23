package models

// Currenci structure
type Currency struct {
	BaseModel BaseModel `gorm:"embedded"`
	Name      string    `gorm:"text; not null; unique" json:"name"`
	Acronym   string    `gorm:"text; not null; unique" json:"acronym"`
}
