package service

import (
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/database/schema"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UploadService struct {
	db *gorm.DB
}

func NewUploadService(db *gorm.DB) *UploadService {
	return &UploadService{db}
}

// Get All
func (service *UploadService) FindAll(c *fiber.Ctx) ([]entity.Upload, int64, error) {
	var data []entity.Upload
	var count int64

	var queryPage string
	var queryPageSize string

	queryPage = c.Query("page")
	queryPageSize = c.Query("pageSize")

	if queryPage == "" {
		queryPage = "1"
	}

	if queryPageSize == "" {
		queryPageSize = "10"
	}

	page, _ := strconv.Atoi(queryPage)
	pageSize, _ := strconv.Atoi(queryPageSize)

	err := service.db.Model(entity.Upload{}).
		// Where("name ILIKE ?", "%admin%").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&data).
		Count(&count).Error

	if err != nil {
		return data, count, err
	}

	return data, count, nil
}

// Find By Id
func (service *UploadService) FindById(id uuid.UUID) (entity.Upload, error) {
	var data entity.Upload

	err := service.db.Model(entity.Upload{}).Where("id = ?", id).First(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// create
func (service *UploadService) Create(input schema.UploadSchema) (entity.Upload, error) {
	data := entity.Upload{}

	data.ID = uuid.New()
	data.KeyFile = input.KeyFile
	data.Filename = input.Filename
	data.Mimetype = input.Mimetype
	data.Size = input.Size
	data.SignedUrl = input.SignedUrl
	data.ExpiryDateUrl = input.ExpiryDateUrl

	err := service.db.Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// Update
func (service *UploadService) Update(id uuid.UUID, input schema.UploadSchema) (entity.Upload, error) {
	data, err := service.FindById(id)
	if err != nil {
		return data, err
	}

	data.KeyFile = input.KeyFile
	data.Filename = input.Filename
	data.Mimetype = input.Mimetype
	data.Size = input.Size
	data.SignedUrl = input.SignedUrl
	data.ExpiryDateUrl = input.ExpiryDateUrl

	err = service.db.Save(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// restore
func (service *UploadService) Restore(id uuid.UUID) error {
	err := service.db.Model(entity.Upload{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error

	if err != nil {
		return err
	}

	return nil
}

// soft delete
func (service *UploadService) SoftDelete(id uuid.UUID) error {
	err := service.db.Delete(&entity.Upload{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// force delete
func (service *UploadService) ForceDelete(id uuid.UUID) error {
	err := service.db.Unscoped().Delete(&entity.Upload{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
