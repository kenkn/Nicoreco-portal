//
// routes.go
// ルータ
//

package routes

import (
	"auth-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/register", controllers.Register)	
	route.Post("/login", controllers.Login)
	route.Get("/user", controllers.User)
	route.Get("/logout", controllers.Logout)
	route.Post("/forgot", controllers.Forgot)
	route.Post("/reset", controllers.Reset)
	route.Get("/question/:subject", controllers.Question)
	route.Post("/qpost", controllers.PostQuestion)
}
