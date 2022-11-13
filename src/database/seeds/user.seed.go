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

var users = []entity.User{
	{
		Fullname: "Super Admin",
	},
	{
		Fullname: "Admin",
	},
	{
		Fullname: "User",
	},
}

func UserSeed(db *gorm.DB) {
	defaultPassword := "Padang123"

	// user seeder
	for i, value := range users {
		var newRoleUUID uuid.UUID
		var newEmail string

		var data entity.User
		var err error

		index := i

		if value.Fullname == "Super Admin" {
			newRoleUUID = uuid.MustParse(constants.ROLE_SUPER_ADMIN)
			newEmail = "super.admin@mail.com"
		}

		if value.Fullname == "Admin" {
			newRoleUUID = uuid.MustParse(constants.ROLE_ADMIN)
			newEmail = "admin@mail.com"
		}

		if value.Fullname == "User" {
			newRoleUUID = uuid.MustParse(constants.ROLE_USER)
			newEmail = "user@mail.com"
		}

		// execute when data not found
		if index < len(users) {
			result := db.Model(entity.User{}).Where("email = ?", newEmail).First(&data)
			userNotFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

			// modif object data
			data = entity.User{
				Base: entity.Base{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Fullname:  value.Fullname,
				Email:     &newEmail,
				Password:  defaultPassword,
				IsActive:  true,
				IsBlocked: false,
				RoleId:    newRoleUUID.String(),
			}

			if userNotFound {
				// create data
				err = db.Model(&entity.User{}).Create(&data).Error

				// error create data
				if err != nil {
					log.Fatalf("cannot seed user table: %v", err)
				}
			}
		}

	}
}
