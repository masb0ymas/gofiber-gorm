package controllers

import (
	"gofiber-gorm/src/app/schema"
	"gofiber-gorm/src/app/service"
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/pkg/helpers"
	"gofiber-gorm/src/pkg/modules/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Get All
func GetAll(c *fiber.Ctx) error {
	db := config.GetDB()

	var queryFiltered config.QueryFiltered

	queryFiltered.Page = c.Query("page")
	queryFiltered.PageSize = c.Query("pageSize")

	roleService := service.NewRoleService(db)
	data, total, err := roleService.GetAll(queryFiltered)

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

	return c.Status(http.StatusOK).JSON(response.HttpResponse(http.StatusOK, "data has been added", data))
}
