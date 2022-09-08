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
	FindAll(queryFiltered config.QueryFiltered) ([]entity.Role, int64, error)
	FindById(id uuid.UUID) (entity.Role, error)
	Create(data schema.RoleSchema) (entity.Role, error)
	Update(id uuid.UUID, data schema.RoleSchema) (entity.Role, error)
	Restore(id uuid.UUID) error
	SoftDelete(id uuid.UUID) error
	ForceDelete(id uuid.UUID) error
}

type RoleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleRepository {
	return &RoleService{db}
}

// Get All
func (service *RoleService) FindAll(queryFiltered config.QueryFiltered) ([]entity.Role, int64, error) {
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

// Find By Id
func (service *RoleService) FindById(id uuid.UUID) (entity.Role, error) {
	var data entity.Role

	err := service.db.Model(entity.Role{}).Where("id = ?", id).First(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
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

// Update
func (service *RoleService) Update(id uuid.UUID, input schema.RoleSchema) (entity.Role, error) {
	data, err := service.FindById(id)
	if err != nil {
		return data, err
	}

	data.Name = input.Name

	err = service.db.Save(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// restore
func (service *RoleService) Restore(id uuid.UUID) error {
	err := service.db.Model(entity.Role{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error

	if err != nil {
		return err
	}

	return nil
}

// soft delete
func (service *RoleService) SoftDelete(id uuid.UUID) error {
	err := service.db.Delete(&entity.Role{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// force delete
func (service *RoleService) ForceDelete(id uuid.UUID) error {
	err := service.db.Unscoped().Delete(&entity.Role{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
