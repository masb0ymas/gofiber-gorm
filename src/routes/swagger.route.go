package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute(app *fiber.App) {
	route := app.Group("/v1/api/docs")

	route.Get("*", swagger.HandlerDefault)
}
