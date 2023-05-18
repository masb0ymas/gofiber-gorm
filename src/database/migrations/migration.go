package migrations

import (
	"fmt"
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/pkg/helpers"
	"strings"

	"gorm.io/gorm"
)

// Run Migration
func Initialize(db *gorm.DB) {
	// List Auto Migrate Table from struct model
	db.AutoMigrate(&entity.Role{}, &entity.User{}, &entity.Session{}, &entity.Upload{})

	collectSchema := []string{
		BaseMigration(),
	}

	schema := strings.Join(collectSchema, ` `)
	result := db.Exec(schema)

	logMigrate := helpers.PrintLog("GORM", "Migration Successfully")
	fmt.Println(logMigrate, result)
}
