package seeds

import (
	"errors"
	"fmt"
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
	for i, value := range roles {
		var newUUID uuid.UUID

		var data entity.Role
		var seedData entity.Seeder
		var newSeedData []entity.Seeder
		var err error

		index := i

		if value.Name == "Super Admin" {
			newUUID = uuid.MustParse(constants.ROLE_SUPER_ADMIN)
		}

		if value.Name == "Admin" {
			newUUID = uuid.MustParse(constants.ROLE_ADMIN)
		}

		if value.Name == "User" {
			newUUID = uuid.MustParse(constants.ROLE_USER)
		}

		// get seeder
		db.Model(entity.Seeder{}).Where("name ILIKE ?", "%role-seeder%").Find(&newSeedData)

		// execute when data not found
		if index < len(roles) {
			result := db.Model(entity.Role{}).Where("id = ?", newUUID).First(&data)
			roleNotFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

			// modif object data
			data = entity.Role{
				Base: entity.Base{
					ID:        newUUID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Name: value.Name,
			}

			if roleNotFound {
				// create data
				err = db.Model(&entity.Role{}).Create(&data).Error

				// error create data
				if err != nil {
					log.Fatalf("cannot seed role table: %v", err)
				}
			}
		}

		// object seeder
		newUnixTime := fmt.Sprintf("role-seeder-%v", time.Now().Unix())
		seedData = entity.Seeder{
			Base: entity.Base{
				ID: uuid.New(),
			},
			Name: newUnixTime,
		}

		// create data when not found
		if len(newSeedData) == 0 {
			err = db.Model(&entity.Seeder{}).Create(&seedData).Error

			// error create data
			if err != nil {
				log.Fatalf("cannot seed role table: %v", err)
			}
		}
	}
}
