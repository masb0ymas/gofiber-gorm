package seeds

import (
	"errors"
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/pkg/constants"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var roles = []entity.Role{
	{
		Name: "Super Admin",
	},
	{
		Name: "Admin",
	},
	{
		Name: "User",
	},
}

func RoleSeed(db *gorm.DB) {
	// role seeder
	for k, v := range roles {
		var newUUID uuid.UUID

		var data entity.Role
		var err error

		if v.Name == "Super Admin" {
			newUUID = uuid.MustParse(constants.ROLE_SUPER_ADMIN)
		}

		if v.Name == "Admin" {
			newUUID = uuid.MustParse(constants.ROLE_ADMIN)
		}

		if v.Name == "User" {
			newUUID = uuid.MustParse(constants.ROLE_USER)
		}

		// execute when data not found
		if k < len(roles) {
			result := db.Model(entity.Role{}).Where("id = ?", newUUID).First(&data)
			recordNotFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

			// modif object data
			data = entity.Role{
				Base: entity.Base{
					ID:        newUUID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Name: v.Name,
			}

			if recordNotFound {
				// create data
				err = db.Model(&entity.Role{}).Create(&data).Error

				// error create data
				if err != nil {
					log.Fatalf("cannot seed role table: %v", err)
				}
			}
		}
	}
}
