package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harish1907/go-udemy/controllers"
	"github.com/harish1907/go-udemy/middleware"
)

func SetupRoute(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.LoginAPI)

	
	app.Use(middleware.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/users", controllers.AllUser)
	app.Post("/api/createuser", controllers.CreateUsers)
	app.Get("/api/users/:id", controllers.GetSingleUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
}