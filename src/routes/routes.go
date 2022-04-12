package routes

import (
	"gofiber-gorm/src/helpers"
	"net/http"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func InitialRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"code":      http.StatusOK,
			"message":   "gofiber-gorm",
			"maintaner": "masb0ymas, <n.fajri@outlook.com>",
			"source":    "https://github.com/masb0ymas/gofiber-gorm",
		},
		)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"code":       http.StatusOK,
			"status":     "Ok",
			"go-version": runtime.Version(),
			"date":       helpers.TimeIn("ID").Format(time.RFC850),
		})
	})

	app.Get("/monitor", monitor.New())

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.NewError(http.StatusNotFound))
	})
}
