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
	route.Post("/forgot", controllers.Forgot)
	route.Post("/reset", controllers.Reset)
	route.Get("/questions/:subject", controllers.GetQuestions)
	route.Get("/question/:id/:user", controllers.GetQuestionInfo)
	route.Post("/question/post", controllers.PostQuestion)
	route.Put("/lgtm/question/:question_id", controllers.LgtmQuestion)
	route.Put("/lgtm/answer/:answer_id", controllers.LgtmAnswer)
	route.Post("/answer/post", controllers.PostAnswer)
	route.Post("/reply/post", controllers.PostReply)
	route.Post("/lab/review/post", controllers.PostLabReview)
	route.Get("/lab/reviews/:lab", controllers.GetLabReviews)
	route.Get("/lab/:id", controllers.GetLabReviewInfo)
	route.Post("/lab/reply/post", controllers.PostLabReply)
	route.Put("/lgtm/lab/:lab_review_id", controllers.LgtmLabReview)
	route.Get("/search/+", controllers.Search)
}
