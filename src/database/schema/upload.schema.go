package schema

import "time"

type UploadSchema struct {
	KeyFile       string    `json:"key_file" validate:"required"`
	Filename      string    `json:"filename" validate:"required"`
	Mimetype      string    `json:"mimetype" validate:"required"`
	Size          int32     `json:"size" validate:"required"`
	SignedUrl     string    `json:"signed_url" validate:"required"`
	ExpiryDateUrl time.Time `json:"expiry_date_url" validate:"required"`
}
