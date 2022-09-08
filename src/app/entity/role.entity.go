package entity

// Role struct to describe role object.
type Role struct {
	Base        // Base Entity
	Name string `json:"name" validate:"required"`
}
