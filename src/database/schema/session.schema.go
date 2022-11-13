package schema

type SessionSchema struct {
	Token     string `json:"token" validate:"required"`
	IpAddress string `json:"ip_address" validate:"required"`
	UserAgent string `json:"user_agent" validate:"required"`
	UserId    string `json:"user_id" validate:"required"`
}
