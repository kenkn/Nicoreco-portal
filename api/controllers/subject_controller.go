//
// subjectController.go
// 科目の質問
//

package controllers

import (
	"auth-api/database"
	"auth-api/models"
	"auth-api/utils"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// /questions/:subject (GET)
// 機能 : 科目の質問の全取得
// 戻り値 : 質問のJSON
func GetQuestions(c *fiber.Ctx) error {

	// GETの内容を取得
	sbj := c.Params("subject")

	// 質問を全検索してリストで取得
	questions := []models.Question{}
	database.DB.Where("subject = ?", sbj).Find(&questions)
	return c.JSON(questions)

}

// /question/:id (GET)
// 機能 : 質問の詳細情報取得
// 戻り値 : 質問の詳細情報のJSON
func GetQuestion(c *fiber.Ctx) error {

	// GETの内容を取得
	id := c.Params("id")

	var question models.Question
	database.DB.Where("id = ?", id).First(&question)

	return c.JSON(question)

}

// /question/post (POST)
// 機能 : 質問をDBに追加する
// 受信するJSON :
//  * jwt			: JWTトークン
//  * questioner_id : 質問者のユーザID
//  * subject       : 質問の科目
//  * title         : 質問のタイトル
//  * body          : 質問の本文
// 戻り値 : 投稿した質問のJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func PostQuestion(c *fiber.Ctx) error {

	data, err := utils.ParseData(c)
	if err != nil {
		return err
	}

	// CookieからJWTを取得(Loginにて保存したユーザ情報)
	cookie := c.Cookies("jwt")
	// JWTtoken取得
	token, err := jwt.ParseWithClaims(cookie, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "認証されていません．",
		})
	}

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

// /lgtm/question/:question_id:user_id (GET)
// 機能 : LGTMされたどうかの判定(question用)
// 戻り値 : LGTM情報のJSON
func IsQuestionLgtmed(c *fiber.Ctx) error {

	// GETの内容を取得
	question_id := c.Params("question_id")
	user_id := c.Params("user_id")

	// 質問を全検索してリストで取得
	lgtm := []models.LgtmQuestion{}
	database.DB.Where("user_id = ?", user_id).Where("question_id = ?", question_id).Find(&lgtm)

	return c.JSON(lgtm)

}

// /lgtm/answer/:answer_id/:user_id (GET)
// 機能 : LGTMされたどうかの判定(answer用)
// 戻り値 : LGTM情報のJSON
func IsAnswerLgtmed(c *fiber.Ctx) error {

	// GETの内容を取得
	answer_id := c.Params("answer_id")
	user_id := c.Params("user_id")

	// 質問を全検索してリストで取得
	lgtm := []models.LgtmAnswer{}
	database.DB.Where("user_id = ?", user_id).Where("answer_id = ?", answer_id).Find(&lgtm)
	return c.JSON(lgtm)

}

// /lgtm/question (POST)
// /lgtm/answer (POST)
// 機能 : 質問のLGTM数を加算する
// 受信するJSON :
//  * jwt 	  : JWTトークン
//  * id 	  : LGTMする質問のID
//  * user_id : LGTMしたユーザーID
// 戻り値 : LGTMした質問のJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func LgtmQuestion(c *fiber.Ctx) error {

	data, err := utils.ParseData(c)
	if err != nil {
		return err
	}

	// CookieからJWTを取得(Loginにて保存したユーザ情報)
	cookie := c.Cookies("jwt")
	// JWTtoken取得
	token, err := jwt.ParseWithClaims(cookie, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "認証されていません．",
		})
	}

	// 既にLGTMされているならDBから削除して、LGTMされてないなら新たにDBに加える
	lgtmData := []models.LgtmQuestion{}
	database.DB.Where("user_id = ?", data["user_id"]).Where("question_id = ?", data["question_id"]).Find(&lgtmData)

	if len(lgtmData) > 0 {
		database.DB.Where("user_id = ?", data["user_id"]).Where("question_id = ?", data["question_id"]).Delete(&lgtmData[0])
	} else {
		parent_id, _ := strconv.Atoi(data["question_id"])
		question_id_uint := uint(parent_id)
		lgtm := models.LgtmQuestion{
			QuestionID: question_id_uint,
			UserID:     data["user_id"],
		}
		database.DB.Create(&lgtm)
	}

	// LGTMの更新
	lgtms := []models.LgtmQuestion{}
	database.DB.Where("question_id = ?", data["question_id"]).Find(&lgtms)
	var question models.Question
	database.DB.Model(&question).Where("id = ?", data["question_id"]).Update("lgtm", len(lgtms))

	return c.JSON(question)

}

func LgtmAnswer(c *fiber.Ctx) error {

	data, err := utils.ParseData(c)
	if err != nil {
		return err
	}

	// CookieからJWTを取得(Loginにて保存したユーザ情報)
	cookie := c.Cookies("jwt")
	// JWTtoken取得
	token, err := jwt.ParseWithClaims(cookie, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "認証されていません．",
		})
	}

	// 既にLGTMされているならDBから削除して、LGTMされてないなら新たにDBに加える
	lgtmData := []models.LgtmAnswer{}
	database.DB.Where("user_id = ?", data["user_id"]).Where("answer_id = ?", data["answer_id"]).Find(&lgtmData)

	if len(lgtmData) > 0 {
		database.DB.Where("user_id = ?", data["user_id"]).Where("answer_id = ?", data["answer_id"]).Delete(&lgtmData[0])
	} else {
		parent_id, _ := strconv.Atoi(data["answer_id"])
		answer_id_uint := uint(parent_id)
		lgtm := models.LgtmAnswer{
			AnswerID: answer_id_uint,
			UserID:   data["user_id"],
		}
		database.DB.Create(&lgtm)
	}

	// LGTMの更新
	lgtms := []models.LgtmAnswer{}
	database.DB.Where("answer_id = ?", data["answer_id"]).Find(&lgtms)
	var answer models.Answer
	database.DB.Model(&answer).Where("id = ?", data["answer_id"]).Update("lgtm", len(lgtms))

	return c.JSON(answer)

}

// /answer/:parent_id (GET)
// 機能 : 該当のquestionに対するanswer一覧を取得する
// 戻り値 : Answer一覧のJSON
func GetAnswer(c *fiber.Ctx) error {

	parent := c.Params("parent_id")
	answers := []models.Answer{}
	database.DB.Where("parent_id = ?", parent).Find(&answers)
	return c.JSON(answers)

}

// /answer/post (POST)
// 機能 : 質問に対する回答をDBに追加する
// 受信するJSON :
//  * parent_id : 回答のquestionのID
//  * user_id 	: 回答したユーザーID
//  * body 		: 回答の本文
// 戻り値 : 回答のJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func PostAnswer(c *fiber.Ctx) error {

	data, err := utils.ParseData(c)
	if err != nil {
		return err
	}

	// CookieからJWTを取得(Loginにて保存したユーザ情報)
	cookie := c.Cookies("jwt")
	// JWTtoken取得
	token, err := jwt.ParseWithClaims(cookie, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "認証されていません．",
		})
	}

	parent_id, _ := strconv.Atoi(data["parent_id"])
	parent_id_uint := uint(parent_id)

	answer := models.Answer{
		ParentID: parent_id_uint,
		UserID:   data["user_id"],
		Body:     data["body"],
		Lgtm:     0,
	}

	database.DB.Create(&answer)

	// questionの回答数の加算
	answers := []models.Answer{}
	database.DB.Where("parent_id = ?", data["parent_id"]).Find(&answers)
	var question models.Question
	database.DB.Model(&question).Where("id = ?", data["parent_id"]).Update("answer_count", len(answers))

	return c.JSON(answer)

}

// /reply/:question_id (GET)
// 機能 : 該当のquestionに対するreply一覧を取得する
// 戻り値 : reply一覧のJSON
// TODO labControllerと同じようにanswerに対してreplyJSONを返すようにする(ページング実装のため)
func GetReply(c *fiber.Ctx) error {

	parent := c.Params("question_id")
	replys := []models.Reply{}
	database.DB.Where("question_id = ?", parent).Find(&replys)
	return c.JSON(replys)

}

// /reply/post (POST)
// 機能 : 回答に対するリプライをDBに追加する
// 受信するJSON :
//  * parent_id : questionのID
//  * user_id 	: リプライを送信したユーザーID
//  * body 		: リプライの本文
// 戻り値 : リプライのJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func PostReply(c *fiber.Ctx) error {

	data, err := utils.ParseData(c)
	if err != nil {
		return err
	}

	// CookieからJWTを取得(Loginにて保存したユーザ情報)
	cookie := c.Cookies("jwt")
	// JWTtoken取得
	token, err := jwt.ParseWithClaims(cookie, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "認証されていません．",
		})
	}

	question_id, _ := strconv.Atoi(data["question_id"])
	question_id_uint := uint(question_id)
	parent_id, _ := strconv.Atoi(data["parent_id"])
	parent_id_uint := uint(parent_id)

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
