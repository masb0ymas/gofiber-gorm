package service

import (
	"gofiber-gorm/src/app/entity"
	"gofiber-gorm/src/app/schema"
	"gofiber-gorm/src/pkg/config"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAll(queryFiltered config.QueryFiltered) ([]entity.Role, int64, error)
	Create(data schema.RoleSchema) (entity.Role, error)
}

type RoleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleRepository {
	return &RoleService{db}
}

// Get All
func (service *RoleService) GetAll(queryFiltered config.QueryFiltered) ([]entity.Role, int64, error) {
	var data []entity.Role
	var count int64

	queryPage, _ := strconv.Atoi(queryFiltered.Page)
	queryPageSize, _ := strconv.Atoi(queryFiltered.PageSize)

	page := queryPage | 1
	pageSize := queryPageSize | 10

	err := service.db.Model(entity.Role{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&data).Error

	// total
	service.db.Model(entity.Role{}).Count(&count)

	if err != nil {
		return data, count, err
	}

	return data, count, nil
}

// create
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
