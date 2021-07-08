// routes.go
package routes

import (
	"auth-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)
	app.Post("/api/forgot", controllers.Forgot)
	app.Post("/api/reset", controllers.Reset)
}
