package models

import (
	"time"

	"github.com/google/uuid"
)

// BaseModel structure
type BaseModel struct {
	UUID      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Enabled   bool `gorm:"default:true"`
}
