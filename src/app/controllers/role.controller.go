package controllers

import (
	"gofiber-gorm/src/app/service"
	"gofiber-gorm/src/database/schema"
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/pkg/helpers"
	"gofiber-gorm/src/pkg/modules/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Find All Role
func FindAllRole(c *fiber.Ctx) error {
	db := config.GetDB()

	roleService := service.NewRoleService(db)
	data, total, err := roleService.FindAll(c)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "error to get roles",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    data,
		"total":   total,
	})
}

// Find Role By Id
func FindRoleById(c *fiber.Ctx) error {
	db := config.GetDB()
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	roleService := service.NewRoleService(db)
	data, err := roleService.FindById(id)

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

// Create Role
func CreateRole(c *fiber.Ctx) error {
	db := config.GetDB()
	roleSchema := new(schema.RoleSchema)

	if err := helpers.ParseFormDataAndValidate(c, roleSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HttpErrorResponse(err))
	}

	// role service
	roleService := service.NewRoleService(db)
	data, err := roleService.Create(*roleSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to create role",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been added",
		"data":    data,
	})
}

// Update Role
func UpdateRole(c *fiber.Ctx) error {
	db := config.GetDB()
	roleSchema := new(schema.RoleSchema)

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	if err := helpers.ParseFormDataAndValidate(c, roleSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HttpErrorResponse(err))
	}

	// role service
	roleService := service.NewRoleService(db)
	data, err := roleService.Update(id, *roleSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to update role",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been updated",
		"data":    data,
	})
}

// Restore
func RestoreRoleById(c *fiber.Ctx) error {
	db := config.GetDB()
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	roleService := service.NewRoleService(db)
	err = roleService.Restore(id)

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
func SoftDeleteRoleById(c *fiber.Ctx) error {
	db := config.GetDB()
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	roleService := service.NewRoleService(db)
	err = roleService.SoftDelete(id)

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
func ForceDeleteRoleById(c *fiber.Ctx) error {
	db := config.GetDB()
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	roleService := service.NewRoleService(db)
	err = roleService.ForceDelete(id)

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
