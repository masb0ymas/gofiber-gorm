package routes

import (
	"gofiber-gorm/src/pkg/helpers"
	"net/http"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Initial routes
func Initialize(app *fiber.App) {

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
			"code":    http.StatusOK,
			"cpu":     runtime.NumCPU(),
			"date":    helpers.TimeIn("ID").Format(time.RFC850),
			"golang":  runtime.Version(),
			"gofiber": fiber.Version,
			"status":  "Ok",
		})
	})

	app.Get("/monitor", monitor.New())

	// initial swagger docs
	SwaggerRoute(app)

	// forbidden route version
	app.Get("/v1", func(c *fiber.Ctx) error {
		return c.Status(http.StatusForbidden).JSON(fiber.NewError(http.StatusForbidden))
	})

	// initial route v1
	RouteV1(app)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.NewError(http.StatusNotFound, "Sorry, HTTP resource you are looking for was not found."))
	})
}
