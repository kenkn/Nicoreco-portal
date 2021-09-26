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
	route.Get("/question/detail/:id", controllers.QuestionDetail)
	route.Post("/question/post", controllers.PostQuestion)
	route.Get("/lgtm/question/:question_id/:user_id", controllers.IsQuestionLgtmed)
	route.Post("/lgtm/question", controllers.LgtmQuestion)
	route.Get("/lgtm/answer/:answer_id/:user_id", controllers.IsAnswerLgtmed)
	route.Post("/lgtm/answer", controllers.LgtmAnswer)
	route.Get("/answer/:parent_id", controllers.Answer)
	route.Post("/answer/post", controllers.PostAnswer)
	route.Get("/reply/:question_id", controllers.Reply)
	route.Post("/reply/post", controllers.PostReply)
}
