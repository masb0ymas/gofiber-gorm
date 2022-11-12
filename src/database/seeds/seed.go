package seeds

import (
	"gorm.io/gorm"
)

func InitialSeed(db *gorm.DB) {
	// list seeder
	RoleSeed(db)
	UserSeed(db)
}
