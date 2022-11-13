package controllers

import (
	"gofiber-gorm/src/app/service"
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/database/schema"
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
	var sessionData entity.Session
	var data entity.User

	db := config.GetDB()
	loginSchema := new(schema.LoginSchema)

	if err := helpers.ParseFormDataAndValidate(c, loginSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HttpErrorResponse(err))
	}

	userService := service.NewUserService(db)
	token, data, err := userService.Login(*loginSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "invalid email or password",
		})
	}

	userAgent := c.Get("User-Agent")

	sessionData = entity.Session{
		Token:     token,
		IpAddress: c.IP(),
		UserId:    data.ID.String(),
		UserAgent: userAgent,
	}

	// create session
	err = db.Model(&entity.Session{}).Create(&sessionData).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to create session",
		})
	}

	resData := fiber.Map{
		"access_token": token,
		"token_type":   "Bearer",
		"uid":          data.ID,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "login successfully",
		"data":    resData,
	})
}

func VerifySession(c *fiber.Ctx) error {
	db := config.GetDB()

	uid := c.Locals("uid")
	newUID := helpers.ParseUUID(uid)

	userService := service.NewUserService(db)
	data, err := userService.FindById(newUID)

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
