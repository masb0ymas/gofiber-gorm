package entity

// Seeder struct to describe role object.
type Seeder struct {
	Base        // Base Entity
	Name string `json:"name" gorm:"not null" validate:"required"`
}
