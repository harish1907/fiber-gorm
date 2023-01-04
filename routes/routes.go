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

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/createrole", controllers.CreateRoles)
	app.Get("/api/roles/:id", controllers.GetSingleRole)
	app.Put("/api/roles/:id", controllers.UpdateRoles)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/view/permissions", controllers.AllViewPermissions)
}