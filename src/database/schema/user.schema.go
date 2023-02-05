package schema

import "gopkg.in/guregu/null.v4"

type UserSchema struct {
	Fullname    string      `json:"fullname" validate:"required"`
	Email       *string     `json:"email" validate:"required"`
	Password    string      `json:"password" validate:"required"`
	Phone       null.String `json:"phone"`
	TokenVerify null.String `json:"token_verify"`
	IsActive    bool        `json:"is_active"`
	IsBlocked   bool        `json:"is_blocked"`
	RoleId      string      `json:"role_id"`
	UploadId    null.String `json:"upload_id"`
}

type RegisterSchema struct {
	Fullname    string      `json:"fullname" validate:"required"`
	Email       *string     `json:"email" validate:"required"`
	Password    string      `json:"password" validate:"required"`
	Phone       null.String `json:"phone"`
	TokenVerify string      `json:"token_verify"`
	RoleId      string      `json:"role_id"`
}

type LoginSchema struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
