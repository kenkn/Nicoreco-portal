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
	route.Post("/post/question", controllers.PostQuestion)
	route.Post("/lgtm", controllers.Lgtm)
	route.Get("/answer/:parent_id", controllers.Answer)
	route.Post("/answer/post", controllers.PostAnswer)
	route.Get("/reply/:parent_id", controllers.Reply)
	route.Post("/reply/post", controllers.PostReply)
}
