package service

import (
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/database/schema"
	"gofiber-gorm/src/pkg/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

// Get All
func (service *UserService) FindAll(c *fiber.Ctx) ([]entity.UserResponse, int64, error) {
	var data []entity.UserResponse
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

	err := service.db.Model(&entity.User{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Preload("Role").
		Find(&data).
		Count(&count).Error

	if err != nil {
		return data, count, err
	}

	return data, count, nil
}

// Find By Id
func (service *UserService) FindById(id uuid.UUID) (entity.UserResponse, error) {
	var data entity.UserResponse

	err := service.db.Model(entity.User{}).Where("id = ?", id).Preload("Role").First(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// create
func (service *UserService) Create(input schema.UserSchema) (entity.User, error) {
	data := entity.User{}

	data.ID = uuid.New()
	data.Fullname = input.Fullname
	data.Email = input.Email
	data.Password = input.Password
	data.Phone = input.Phone
	data.TokenVerify = input.TokenVerify
	data.IsActive = input.IsActive
	data.RoleId = input.RoleId

	err := service.db.Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// Update
func (service *UserService) Update(id uuid.UUID, input schema.UserSchema) (entity.UserResponse, error) {
	data, err := service.FindById(id)
	if err != nil {
		return data, err
	}

	data.Fullname = input.Fullname

	err = service.db.Save(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

// restore
func (service *UserService) Restore(id uuid.UUID) error {
	err := service.db.Model(entity.User{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error

	if err != nil {
		return err
	}

	return nil
}

// soft delete
func (service *UserService) SoftDelete(id uuid.UUID) error {
	err := service.db.Delete(&entity.User{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// force delete
func (service *UserService) ForceDelete(id uuid.UUID) error {
	err := service.db.Unscoped().Delete(&entity.User{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// Login
func (service *UserService) Login(input schema.LoginSchema) (string, entity.User, error) {
	var err error
	var data entity.User

	err = service.db.Model(entity.User{}).Where("email = ?", input.Email).First(&data).Error

	if err != nil {
		return "", entity.User{}, err
	}

	err = helpers.ComparePassword(input.Password, data.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", entity.User{}, err
	}

	token, err := helpers.GenerateToken(data.ID)

	if err != nil {
		return "", entity.User{}, err
	}

	return token, data, nil
}
