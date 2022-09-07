package entity

type Role struct {
	Base        // Base Entity
	Name string `json:"name" validate:"required"`
}
