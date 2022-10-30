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

func Register(c *fiber.Ctx) error {
	db := config.GetDB()
	userSchema := new(schema.UserSchema)

	if err := helpers.ParseFormDataAndValidate(c, userSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HttpErrorResponse(err))
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
		"message": "registration success",
		"data":    data,
	})
}

func Login(c *fiber.Ctx) error {
	db := config.GetDB()
	loginSchema := new(schema.LoginSchema)

	if err := helpers.ParseFormDataAndValidate(c, loginSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HttpErrorResponse(err))
	}

	userService := service.NewUserService(db)
	token, err := userService.Login(*loginSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "invalid email or password",
		})
	}

	data := fiber.Map{
		"access_token": token,
		"token_type":   "Bearer",
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "login successfully",
		"data":    data,
	})
}
