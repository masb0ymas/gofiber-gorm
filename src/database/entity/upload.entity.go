package entity

import "time"

// Upload struct to describe upload object.
type Upload struct {
	Base                    // Base Entity
	KeyFile       string    `json:"key_file" gorm:"not null" validate:"required"`
	Filename      string    `json:"filename" gorm:"not null" validate:"required"`
	Mimetype      string    `json:"mimetype" gorm:"not null" validate:"required"`
	Size          int32     `json:"size" gorm:"not null" validate:"required"`
	SignedUrl     string    `json:"signed_url" gorm:"type:text; not null" validate:"required"`
	ExpiryDateUrl time.Time `json:"expiry_date_url" gorm:"not null" validate:"required"`
}
