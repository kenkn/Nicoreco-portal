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

// /question/:subject (GET)
// 機能 : 科目の質問の全取得
// 戻り値 : 質問のJSON
func Question(c *fiber.Ctx) error {

	// GETの内容を取得
	sbj := c.Params("subject")

	// 質問を全検索してリストで取得
	questions := []models.Question{}
	database.DB.Where("subject = ?", sbj).Find(&questions)
	return c.JSON(questions)

}

// /question/detail/:id (GET)
// 機能 : 質問の詳細情報取得
// 戻り値 : 質問の詳細情報のJSON
func QuestionDetail(c *fiber.Ctx) error {

	// GETの内容を取得
	id := c.Params("id")

	var detail models.Question
	database.DB.Where("id = ?", id).First(&detail)

	return c.JSON(detail)

}

// /question/post (POST)
// 機能 : 質問の投稿
// 受信するJSON :
//  * questioner_id : 質問者のユーザID
//  * subject       : 質問の科目
//  * title         : 質問のタイトル
//  * body          : 質問の本文
// 戻り値 : 投稿した質問のJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func PostQuestion(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var userInfo models.User
	database.DB.Where("user_id = ?", data["questioner_id"]).First(&userInfo)

	question := models.Question{
		QuestionerID: data["questioner_id"],
		Subject:      data["subject"],
		Title:        data["title"],
		Body:         data["body"],
		Lgtm:         0,
	}

	// データ登録(CreateはGORMメソッド)
	database.DB.Create(&question)

	return c.JSON(question)

}

// /lgtm/question (POST)
// /lgtm/answer (POST)
// 機能 : 質問のLGTM数を加算する
// 受信するJSON :
//  * id : LGTMする質問のID
//  * user_id : LGTMしたユーザーID
// 戻り値 : LGTMした質問のJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func LgtmQuestion(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// 既にLGTMされているならDBから削除して、LGTMされてないなら新たにDBに加える
	pressed := []models.LgtmQuestion{}
	database.DB.Where("user_id = ?", data["user_id"]).Where("parent_id = ?", data["id"]).Find(&pressed)
	if len(pressed) >= 1 {
		database.DB.Where("user_id = ?", data["user_id"]).Where("parent_id = ?", data["id"]).Delete(&pressed[0])
	} else {
		parent_id, _ := strconv.Atoi(data["id"])
		parent_id_uint := uint(parent_id)
		lgtm := models.LgtmQuestion{
			ParentID: parent_id_uint,
			UserID:   data["user_id"],
		}
		database.DB.Create(&lgtm)
	}

	// LGTMの更新
	lgtms := []models.LgtmQuestion{}
	database.DB.Where("parent_id = ?", data["id"]).Find(&lgtms)
	var question models.Question
	database.DB.Model(&question).Where("id = ?", data["id"]).Update("lgtm", len(lgtms))

	return c.JSON(question)

}

func LgtmAnswer(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// 既にLGTMされているならDBから削除して、LGTMされてないなら新たにDBに加える
	pressed := []models.LgtmAnswer{}
	database.DB.Where("user_id = ?", data["user_id"]).Where("parent_id = ?", data["id"]).Find(&pressed)
	if len(pressed) >= 1 {
		database.DB.Where("user_id = ?", data["user_id"]).Where("parent_id = ?", data["id"]).Delete(&pressed[0])
	} else {
		parent_id, _ := strconv.Atoi(data["id"])
		parent_id_uint := uint(parent_id)
		lgtm := models.LgtmAnswer{
			ParentID: parent_id_uint,
			UserID:   data["user_id"],
		}
		database.DB.Create(&lgtm)
	}

	// LGTMの更新
	lgtms := []models.LgtmAnswer{}
	database.DB.Where("parent_id = ?", data["id"]).Find(&lgtms)
	var answer models.Answer
	database.DB.Model(&answer).Where("id = ?", data["id"]).Update("lgtm", len(lgtms))

	return c.JSON(answer)

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

	answer := models.Answer{
		ParentID: parent_id_uint,
		UserID:   data["user_id"],
		Body:     data["body"],
		Lgtm:     0,
	}

	database.DB.Create(&answer)

	return c.JSON(answer)

}

func Reply(c *fiber.Ctx) error {

	parent := c.Params("question_id")
	replys := []models.Reply{}
	database.DB.Where("question_id = ?", parent).Find(&replys)
	return c.JSON(replys)

}

func PostReply(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	question_id, _ := strconv.Atoi(data["question_id"])
	question_id_uint := uint(question_id)
	parent_id, _ := strconv.Atoi(data["parent_id"])
	parent_id_uint := uint(parent_id)

	var replyInfo models.Reply
	database.DB.Where("id = ?", parent_id_uint).First(&replyInfo)

	reply := models.Reply{
		QuestionID: question_id_uint,
		ParentID:   parent_id_uint,
		UserID:     data["user_id"],
		Body:       data["body"],
		Lgtm:       0,
	}

	database.DB.Create(&reply)

	return c.JSON(reply)

}
