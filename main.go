package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/harish1907/go-udemy/database"
	"github.com/harish1907/go-udemy/routes"
)

func init() {
	database.EnviormentVariable()
	database.DBConnection()
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.SetupRoute(app)
	app.Listen(os.Getenv("PORT"))
}
