package main

import (
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func main() {
	app := fiber.New()

	port := config.Env("APP_PORT", "8000")

	// default middleware
	app.Use(cors.New(config.Cors()))
	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{Max: 100}))

	// static file
	app.Static("/", "./public")

	// Connect to the Database
	config.ConnectDB()

	// initial app route
	routes.InitialRoutes(app)

	// listening app
	log.Fatal(app.Listen(":" + port))
}
