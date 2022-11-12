package service

import (
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/database/schema"
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/pkg/helpers"
	"strconv"

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
func (service *UserService) FindAll(queryFiltered config.QueryFiltered) ([]entity.User, int64, error) {
	var data []entity.User
	var count int64

	queryPage, _ := strconv.Atoi(queryFiltered.Page)
	queryPageSize, _ := strconv.Atoi(queryFiltered.PageSize)

	page := queryPage | 1
	pageSize := queryPageSize | 10

	err := service.db.Model(&entity.User{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Preload("Role").
		Find(&data).Error

	// total
	service.db.Model(entity.User{}).Count(&count)

	if err != nil {
		return data, count, err
	}

	return data, count, nil
}

// Find By Id
func (service *UserService) FindById(id uuid.UUID) (entity.User, error) {
	var data entity.User

	err := service.db.Model(entity.User{}).Where("id = ?", id).First(&data).Error

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
func (service *UserService) Update(id uuid.UUID, input schema.UserSchema) (entity.User, error) {
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
func (service *UserService) Login(input schema.LoginSchema) (string, error) {
	var err error
	var data entity.User

	err = service.db.Model(entity.User{}).Where("email = ?", input.Email).First(&data).Error

	if err != nil {
		return "", err
	}

	err = helpers.ComparePassword(input.Password, data.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := helpers.GenerateToken(data.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
