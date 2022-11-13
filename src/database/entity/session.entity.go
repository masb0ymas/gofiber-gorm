package entity

// Session struct to describe role object.
type Session struct {
	Base             // Base Entity
	Token     string `json:"token" gorm:"not null" validate:"required"`
	IpAddress string `json:"ip_address" gorm:"not null" validate:"required"`
	UserAgent string `json:"user_agent" gorm:"type:text; not null"`
	UserId    string `json:"user_id" gorm:"type:uuid; not null" validate:"required"`
	User      User   `json:"User"`
}
