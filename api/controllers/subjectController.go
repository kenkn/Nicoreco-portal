//
// subjectController.go
// 科目の質問
//

package controllers

import (
	"auth-api/database"
	"auth-api/models"
	"strconv"
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

func Answer(c *fiber.Ctx) error {

	parent := c.Params("parent_id")
	answers := []models.Answer{}
	database.DB.Where("parent_id = ?", parent).Find(&answers)
	return c.JSON(answers)

}

func PostAnswer(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	parent_id, _ := strconv.Atoi(data["parent_id"])
	parent_id_uint := uint(parent_id)
	var questionInfo models.Question
	database.DB.Where("id = ?", parent_id).First(&questionInfo)

	answer := models.Answer {
		ParentID : parent_id_uint,
		UserID	 : data["user_id"],
		Body		 : data["body"],
		Lgtm		 : 0,
	}

	database.DB.Create(&answer)

	return c.JSON(answer)

}

func Reply(c *fiber.Ctx) error {

	parent := c.Params("parent_id")
	replys := []models.Reply{}
	database.DB.Where("parent_id = ?", parent).Find(&replys)
	return c.JSON(replys)

}

func PostReply(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	parent_id, _ := strconv.Atoi(data["parent_id"])
	parent_id_uint := uint(parent_id)
	var replyInfo models.Reply
	database.DB.Where("id = ?", parent_id_uint).First(&replyInfo)

	reply := models.Reply {
		ParentID : parent_id_uint,
		UserID	 : data["user_id"],
		Body		 : data["body"],
		Lgtm		 : 0,
	}

	database.DB.Create(&reply)

	return c.JSON(reply)

}