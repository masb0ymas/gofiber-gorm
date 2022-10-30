package middlewares

import (
	"gofiber-gorm/src/pkg/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := helpers.TokenValid(c)

		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}

		return c.Next()
	}
}
