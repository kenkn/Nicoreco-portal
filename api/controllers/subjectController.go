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

// /post/question (POST)
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

    question := models.Question {
        QuestionerID : data["questioner_id"],
        Subject      : data["subject"],
        Title        : data["title"],
        Body         : data["body"],
        Lgtm         : 0,
    }

    // データ登録(CreateはGORMメソッド)
    database.DB.Create(&question)

    return c.JSON(question)

}

// /lgtm (POST)
// 機能 : 質問のLGTM数を加算する
// 受信するJSON :
//  * id : LGTMする質問のID
// 戻り値 : LGTMした質問のJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func Lgtm(c *fiber.Ctx) error {

    var data map[string]string
    // リクエストデータをパース
    if err := c.BodyParser(&data); err != nil {
        return err
    }

    // リクエストデータをint型にキャスト
    lgtm, _ := strconv.Atoi(data["id"])

    // 元のLGTM数に1加算して更新
    var question models.Question
    database.DB.Model(&question).Where("id = ?", data["id"]).Update("lgtm", lgtm+1)

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
        UserID   : data["user_id"],
        Body         : data["body"],
        Lgtm         : 0,
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
        UserID   : data["user_id"],
        Body         : data["body"],
        Lgtm         : 0,
    }

    database.DB.Create(&reply)

    return c.JSON(reply)

}
