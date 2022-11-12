package schema

type RoleSchema struct {
	Name string `json:"name" validate:"required"`
}
