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
	route.Post("/user", controllers.User)
	route.Post("/forgot", controllers.Forgot)
	route.Post("/reset", controllers.Reset)
	route.Get("/questions/:subject", controllers.GetQuestions)
	route.Get("/question/:id", controllers.GetQuestion)
	route.Post("/question/post", controllers.PostQuestion)
	route.Get("/lgtm/question/:question_id/:user_id", controllers.IsQuestionLgtmed)
	route.Post("/lgtm/question", controllers.LgtmQuestion)
	route.Get("/lgtm/answer/:answer_id/:user_id", controllers.IsAnswerLgtmed)
	route.Post("/lgtm/answer", controllers.LgtmAnswer)
	route.Get("/answer/:parent_id", controllers.GetAnswer)
	route.Post("/answer/post", controllers.PostAnswer)
	route.Get("/reply/:question_id", controllers.GetReply)
	route.Post("/reply/post", controllers.PostReply)
	route.Post("/lab/review/post", controllers.PostLabReview)
	route.Get("/lab/reviews/:lab", controllers.LabReviews)
	route.Get("/lab/review/:id", controllers.LabReview)
	route.Post("/lab/reply/post", controllers.PostLabReply)
	route.Get("/lab/reply/:lab_review_id", controllers.GetLabReply)
	route.Get("/lgtm/lab/:lab_review_id/:user_id", controllers.IsLabReviewLgtmed)
	route.Post("/lgtm/lab", controllers.LgtmLabReview)
}
