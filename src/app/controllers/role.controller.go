package controllers

import (
	"gofiber-gorm/src/app/schema"
	"gofiber-gorm/src/app/service"
	"gofiber-gorm/src/config"
	"gofiber-gorm/src/helpers"
	"gofiber-gorm/src/modules/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

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
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to create role",
		})
	}

	return c.Status(http.StatusOK).JSON(response.HttpResponse(http.StatusOK, "data has been added", data))
}
