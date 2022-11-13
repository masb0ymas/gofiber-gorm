package service

import (
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/database/schema"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionService struct {
	db *gorm.DB
}

func NewSessionService(db *gorm.DB) *SessionService {
	return &SessionService{db}
}

// Get All
func (service *SessionService) FindAll(c *fiber.Ctx) ([]entity.Session, int64, error) {
	var data []entity.Session
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

	err := service.db.Model(entity.Session{}).
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
func (service *SessionService) FindById(id uuid.UUID) (entity.Session, error) {
	var data entity.Session

	err := service.db.Model(entity.Session{}).Where("id = ?", id).First(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// create
func (service *SessionService) Create(input schema.SessionSchema) (entity.Session, error) {
	data := entity.Session{}

	data.ID = uuid.New()
	data.Token = input.Token
	data.IpAddress = input.IpAddress
	data.UserAgent = input.UserAgent
	data.UserId = input.UserId

	err := service.db.Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// Update
func (service *SessionService) Update(id uuid.UUID, input schema.SessionSchema) (entity.Session, error) {
	data, err := service.FindById(id)
	if err != nil {
		return data, err
	}

	data.Token = input.Token
	data.IpAddress = input.IpAddress
	data.UserAgent = input.UserAgent
	data.UserId = input.UserId

	err = service.db.Save(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// restore
func (service *SessionService) Restore(id uuid.UUID) error {
	err := service.db.Model(entity.Session{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error

	if err != nil {
		return err
	}

	return nil
}

// soft delete
func (service *SessionService) SoftDelete(id uuid.UUID) error {
	err := service.db.Delete(&entity.Session{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// force delete
func (service *SessionService) ForceDelete(id uuid.UUID) error {
	err := service.db.Unscoped().Delete(&entity.Session{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
