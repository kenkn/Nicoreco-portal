//
// subjectController.go
// 科目の質問
//

package controllers

import (
	"auth-api/database"
	"auth-api/models"

	"github.com/gofiber/fiber/v2"
)

func Question(c *fiber.Ctx) error {

	sbj := c.Params("subject")
	questions := []models.Question{}
	database.DB.Where("subject = ?", sbj).Find(&questions)
	return c.JSON(questions)

}

func PostQuestion(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var userInfo models.User
	database.DB.Where("user_id = ?", data["questioner_id"]).First(&userInfo)

	question := models.Question {
		QuestionerID : data["questioner_id"],
		Subject		 : data["subject"],
		Title		 : data["title"],
		Body		 : data["body"],
		Lgtm		 : 0,
	}

	database.DB.Create(&question)

	return c.JSON(question)

}