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
	for _, value := range roles {
		var newUUID = uuid.UUID{}
		var data = entity.Role{}

		if value.Name == "Super Admin" {
			newUUID = uuid.MustParse(constants.ROLE_SUPER_ADMIN)
		}

		if value.Name == "Admin" {
			newUUID = uuid.MustParse(constants.ROLE_ADMIN)
		}

		if value.Name == "User" {
			newUUID = uuid.MustParse(constants.ROLE_USER)
		}

		// check ID
		result := db.Model(entity.Role{}).Where("id = ?", newUUID).First(&data)
		dataNotFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

		// modif object data
		data = entity.Role{
			Base: entity.Base{
				ID:        newUUID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: value.Name,
		}

		if dataNotFound {
			// create data
			err := db.Model(&entity.Role{}).Create(&data).Error

			// error create data
			if err != nil {
				log.Fatalf("cannot seed role table: %v", err)
			}
		} else {
			// update data
			err := db.Save(&data).Error

			// error update data
			if err != nil {
				log.Fatalf("cannot seed role table: %v", err)
			}
		}
	}
}