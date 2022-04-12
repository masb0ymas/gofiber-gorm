package main

import (
	"gofiber-gorm/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// default middleware
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{Max: 100}))

	// static file
	app.Static("/", "./public")

	// initial app route
	routes.InitialRoutes(app)

	// listening app
	log.Fatal(app.Listen(":8000"))
}
