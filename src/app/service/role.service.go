package service

import (
	"gofiber-gorm/src/app/entity"
	"gofiber-gorm/src/app/schema"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(data schema.RoleSchema) (entity.Role, error)
}

type RoleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleRepository {
	return &RoleService{db}
}

func (service *RoleService) Create(input schema.RoleSchema) (entity.Role, error) {
	data := entity.Role{}

	data.ID = uuid.New()
	data.Name = input.Name

	err := service.db.Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}
