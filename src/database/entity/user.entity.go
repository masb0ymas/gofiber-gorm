package entity

import (
	"gofiber-gorm/src/pkg/helpers"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type User struct {
	Base                    // Base Entity
	Fullname    string      `json:"fullname" gorm:"not null" validate:"required"`
	Email       *string     `json:"email" gorm:"unique; size:255; not null" validate:"required"`
	Password    string      `json:"password" gorm:"not null" validate:"required"`
	Phone       null.String `json:"phone" gorm:"size:20"`
	TokenVerify null.String `json:"token_verify" gorm:"type:text; column:token_verify;"`
	IsActive    bool        `json:"is_active" gorm:"type:boolean; default:false; not_null"`
	IsBlocked   bool        `json:"is_blocked" gorm:"type:boolean; default:false; not_null"`
	RoleId      string      `json:"role_id" gorm:"type:uuid; not null" validate:"required"`
	Role        Role        `json:"Role"`
	UploadId    null.String `json:"upload_id" gorm:"type:uuid"`
}

type UserResponse struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Fullname  string         `json:"fullname"`
	Email     *string        `json:"email"`
	Phone     null.String    `json:"phone"`
	IsActive  bool           `json:"is_active"`
	IsBlocked bool           `json:"is_blocked"`
	RoleId    string         `json:"role_id"`
	Role      Role           `json:"Role"`
	UploadId  null.String    `json:"upload_id"`
}

// GORM Hooks
func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := helpers.HashPassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}
