package controllers

import (
	"gofiber-gorm/src/app/service"
	"gofiber-gorm/src/database/schema"
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/pkg/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Find All User
func FindAllUser(c *fiber.Ctx) error {
	db := config.GetDB()

	userService := service.NewUserService(db)
	data, total, err := userService.FindAll(c)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "error to get users",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    data,
		"total":   total,
	})
}

// Find User By Id
func FindUserById(c *fiber.Ctx) error {
	db := config.GetDB()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get user",
		})
	}

	userService := service.NewUserService(db)
	data, err := userService.FindById(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    data,
	})
}

// Create User
func CreateUser(c *fiber.Ctx) error {
	db := config.GetDB()

	userSchema := new(schema.UserSchema)

	if err := helpers.ParseFormDataAndValidate(c, userSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.NewError(http.StatusBadRequest))
	}

	// user service
	userService := service.NewUserService(db)
	data, err := userService.Create(*userSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to create user",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been added",
		"data":    data,
	})
}

// Update User
func UpdateUser(c *fiber.Ctx) error {
	db := config.GetDB()

	userSchema := new(schema.UserSchema)

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get user",
		})
	}

	if err := helpers.ParseFormDataAndValidate(c, userSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.NewError(http.StatusBadRequest))
	}

	// user service
	userService := service.NewUserService(db)
	data, err := userService.Update(id, *userSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to update user",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been updated",
		"data":    data,
	})
}

// Restore
func RestoreUserById(c *fiber.Ctx) error {
	db := config.GetDB()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get user",
		})
	}

	userService := service.NewUserService(db)
	err = userService.Restore(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been updated",
	})
}

// Soft Delete
func SoftDeleteUserById(c *fiber.Ctx) error {
	db := config.GetDB()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get user",
		})
	}

	userService := service.NewUserService(db)
	err = userService.SoftDelete(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been deleted",
	})
}

// Force Delete
func ForceDeleteUserById(c *fiber.Ctx) error {
	db := config.GetDB()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get user",
		})
	}

	userService := service.NewUserService(db)
	err = userService.ForceDelete(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been deleted",
	})
}
