package schema

import "gopkg.in/guregu/null.v4"

type SessionSchema struct {
	Token     string      `json:"token" validate:"required"`
	IpAddress string      `json:"ip_address" validate:"required"`
	Device    null.String `json:"device"`
	Platform  null.String `json:"platform"`
	UserId    string      `json:"user_id" validate:"required"`
}
