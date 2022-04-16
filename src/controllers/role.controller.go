package controllers

import (
	"gofiber-gorm/src/helpers"
	"gofiber-gorm/src/models/entity"
	"gofiber-gorm/src/models/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	roleEntity := new(entity.Role)

	if err := helpers.ParseFormDataAndValidate(c, roleEntity); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HttpErrorResponse(err))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"formdata": roleEntity,
	})
}
