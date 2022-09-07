package entity

import "database/sql"

type User struct {
	Base                       // Base Entity
	Fullname    string         `json:"fullname" validate:"required"`
	Email       *string        `json:"email" gorm:"unique; size:255" validate:"required"`
	Password    string         `json:"password" validate:"required"`
	Phone       sql.NullString `json:"phone" gorm:"size:20"`
	TokenVerify sql.NullString `json:"token_verify" gorm:"type:text; column:token_verify;"`
	IsActive    bool           `json:"is_active" gorm:"type:boolean; default:false; not_null"`
	IsBlocked   bool           `json:"is_blocked" gorm:"type:boolean; default:false; not_null"`
	RoleId      string         `json:"role_id" gorm:"type:uuid" validate:"required"`
	Role        Role           `json:"Role"`
	UploadId    string         `json:"upload_id" gorm:"type:uuid"`
}
