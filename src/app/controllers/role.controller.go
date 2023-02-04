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

// GetRoles 		func gets all exists roles.
// @Description Get all exists roles.
// @Summary 		get all exists roles
// @Tags 				Role
// @Accept 			json
// @Produce 		json
// @Success 		200 {string} status "Ok"
// @Router 			/v1/role [get]
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

// GetRole 			func gets role by given ID or 404 error.
// @Description Get role by given ID.
// @Summary 		get role by given ID
// @Tags 				Role
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Role ID"
// @Success 		200 {string} status "Ok"
// @Router 			/v1/role/{id} [get]
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

// CreateRole 	func for creates a new role.
// @Description Create a new role.
// @Summary 		create a new role
// @Tags 				Role
// @Accept 			x-www-form-urlencoded
// @Produce 		json
// @Param 			name formData string true "Name"
// @Success 		200 {string} status "Ok"
// @Security 		ApiKeyAuth
// @Router 			/v1/role [post]
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

// UpdateRole 	func for updates role by given ID.
// @Description Update role.
// @Summary 		update role
// @Tags 				Role
// @Accept 			x-www-form-urlencoded
// @Produce 		json
// @Param 			id path string true "Role ID"
// @Param 			name formData string true "Name"
// @Success 		200 {string} status "Ok"
// @Security 		ApiKeyAuth
// @Router 			/v1/role/{id} [put]
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

// RestoreRole 	func for Restores role by given ID.
// @Description Restore role by given ID.
// @Summary 		Restore role by given ID
// @Tags 				Role
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Role ID"
// @Success 		200 {string} status "Ok"
// @Security 		ApiKeyAuth
// @Router 			/v1/role/restore/{id} [put]
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

// SoftDeleteRole 	func for Soft Deletes role by given ID.
// @Description 		Soft Delete role by given ID.
// @Summary 				Soft Delete role by given ID
// @Tags 						Role
// @Accept 					json
// @Produce 				json
// @Param 					id path string true "Role ID"
// @Success 				200 {string} status "Ok"
// @Security 				ApiKeyAuth
// @Router 					/v1/role/soft-delete/{id} [delete]
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

// ForceDeleteRole 	func for Force Deletes role by given ID.
// @Description 		Force Delete role by given ID.
// @Summary 				Force Delete role by given ID
// @Tags 						Role
// @Accept 					json
// @Produce 				json
// @Param 					id path string true "Role ID"
// @Success 				200 {string} status "Ok"
// @Security 				ApiKeyAuth
// @Router 					/v1/role/force-delete/{id} [delete]
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
