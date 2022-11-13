package entity

import "gopkg.in/guregu/null.v4"

// Session struct to describe role object.
type Session struct {
	Base                  // Base Entity
	Token     string      `json:"token" gorm:"not null" validate:"required"`
	IpAddress string      `json:"ip_address" gorm:"not null" validate:"required"`
	Device    null.String `json:"device"`
	Platform  null.String `json:"platform"`
	UserId    string      `json:"user_id" gorm:"type:uuid; not null" validate:"required"`
	User      User        `json:"User"`
}
