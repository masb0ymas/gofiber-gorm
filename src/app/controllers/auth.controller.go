package controllers

import (
	"gofiber-gorm/src/app/service"
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/database/schema"
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/pkg/constants"
	"gofiber-gorm/src/pkg/helpers"
	"gofiber-gorm/src/pkg/modules/response"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Register 		func for register account.
// @Description Create a new account.
// @Summary 		create a new account
// @Tags 				Auth
// @Accept 			x-www-form-urlencoded
// @Produce 		json
// @Param 			fullname formData string true "Fullname"
// @Param 			email formData string true "Email"
// @Param 			password formData string true "Password"
// @Param 			phone formData string false "Phone"
// @Success 		200 {string} status "Ok"
// @Router 			/v1/auth/sign-up [post]
func Register(c *fiber.Ctx) error {
	db := config.GetDB()
	userSchema := new(schema.RegisterSchema)

	token, _ := helpers.GenerateToken(uuid.New())

	userSchema.RoleId = constants.ROLE_USER
	userSchema.TokenVerify = token

	if err := helpers.ParseFormDataAndValidate(c, userSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HttpErrorResponse(err))
	}

	// user service
	userService := service.NewUserService(db)
	data, err := userService.Register(*userSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "registration success",
		"data":    data,
	})
}

// Login 				func for login account.
// @Description Login account.
// @Summary 		Login account
// @Tags 				Auth
// @Accept 			x-www-form-urlencoded
// @Produce 		json
// @Param 			email formData string true "Email"
// @Param 			password formData string true "Password"
// @Success 		200 {string} status "Ok"
// @Router 			/v1/auth/sign-in [post]
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
			"message": err.Error(),
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

	expiresToken := os.Getenv("JWT_ACCESS_TOKEN_EXPIRED")
	expiresIn, _ := strconv.Atoi(expiresToken) // expires in days
	expiresCookie := time.Now().Add(time.Hour * 24 * time.Duration(expiresIn))

	// set login by cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiresCookie,
		HTTPOnly: true,
		SameSite: "lax",
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "login successfully",
		"data":    resData,
	})
}

// VerifySession	func for verify session login account.
// @Description 	Verify Session Login account.
// @Summary 			Verify Session Login account
// @Tags 					Auth
// @Accept 				json
// @Produce 			json
// @Success 			200 {string} status "Ok"
// @Security 			ApiKeyAuth
// @Router 				/v1/auth/verify-session [get]
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
