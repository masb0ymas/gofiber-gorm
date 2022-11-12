package seeds

import (
	"gorm.io/gorm"
)

func InitialSeed(db *gorm.DB) {
	// run role seed
	RoleSeed(db)
}
