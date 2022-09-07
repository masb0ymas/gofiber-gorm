package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	// gorm.Model
	ID        uuid.UUID      `json:"id" gorm:"unique; type:uuid; column:id; default:uuid_generate_v4(); not_null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
