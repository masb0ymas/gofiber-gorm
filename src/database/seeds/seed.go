package seeds

import (
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) {
	// list seeder
	RoleSeed(db)
	UserSeed(db)
}
