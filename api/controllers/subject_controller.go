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
	database.DB.Where("subject = ?", sbj).Order("lgtm desc").Find(&questions)
	return c.JSON(questions)

}

// /question/:id (GET)
// 機能 : 質問の詳細情報取得
// <user>が"unauthorized"の時は非ログイン時であるためLGTMedはFalseとする
// 戻り値 : 質問の詳細情報のJSON
func GetQuestionInfo(c *fiber.Ctx) error {

	// GETの内容を取得
	id := c.Params("id")
	cookie := c.Cookies("jwt")
	isGetLgtmled := true
	// JWTtoken取得
	token, err := jwt.ParseWithClaims(cookie, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		isGetLgtmled = false
	}

	claims := token.Claims.(*utils.Claims)
	// User IDを取得
	userID := claims.Issuer

	// Question情報及びuserがLGTMedか否かの取得
	var question models.Question
	database.DB.Where("id = ?", id).First(&question)
	isQuestionLgtmed := false
	if isGetLgtmled {
		var lgtmQuestion models.LgtmQuestion
		res := database.DB.Where("question_id = ?", id).Where("lgtmer_id", userID).First(&lgtmQuestion)
		if res.Error == nil && lgtmQuestion.IsLgtmed {
			isQuestionLgtmed = true
		}
	}

	// Answer情報の取得
	answers := []models.Answer{}
	database.DB.Where("question_id = ?", question.ID).Order("lgtm desc").Find(&answers)

	// Reply情報及びAnswerがLGTMedか否かの取得
	answer := []fiber.Map{}
	for _, ans := range answers {
		reply := []models.Reply{}
		database.DB.Where("answer_id = ?", ans.ID).Order("created_at").Find(&reply)
		isAnswerLgtmed := false
		if isGetLgtmled {
			var lgtmAnswer models.LgtmAnswer
			res := database.DB.Where("answer_id = ?", ans.ID).Where("lgtmer_id = ?", userID).First(&lgtmAnswer)
			if res.Error == nil {
				isAnswerLgtmed = true
			}
		} 
		answer = append(answer, fiber.Map{
			"answer":   ans,
			"islgtmed": isAnswerLgtmed,
			"reply":    reply,
		})
	}

	return c.JSON(fiber.Map{
		"question": question,
		"islgtmed": isQuestionLgtmed,
		"answers":  answer,
	})

}

// /question/post (POST)
// 機能 : 質問をDBに追加する
// 受信するJSON :
//  * jwt			: JWTトークン
//  * user_id       : 質問者のユーザID
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
		QuestionerID: data["user_id"],
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
	database.DB.Where("lgtmer_id = ?", user_id).Where("question_id = ?", question_id).Find(&lgtm)

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
	database.DB.Where("lgtmer_id = ?", user_id).Where("answer_id = ?", answer_id).Find(&lgtm)
	return c.JSON(lgtm)

}

// /lgtm/question/:question_id (PUT)
// /lgtm/answer/:answer_id (PUT)
// 機能 : 質問のLGTM数を加算する
// 受信するJSON :
//  * id 	  : LGTMする質問のID
// 戻り値 : LGTMした質問のJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
// HACK: LGTMクエリをDELETEしない，増やさない
func LgtmQuestion(c *fiber.Ctx) error {

	questionID := c.Params("question_id")

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
	claims := token.Claims.(*utils.Claims)
	// User IDを取得
	userID := claims.Issuer

	lgtmData := models.LgtmQuestion{}
	res := database.DB.Where("lgtmer_id = ?", userID).Where("question_id = ?", questionID).First(&lgtmData)

	if res.Error == nil {
		// LGTM情報が既にある場合
		lq := models.LgtmQuestion{}
		database.DB.Model(&lq).Where("lgtmer_id = ?", userID).Where("question_id = ?", questionID).Update("is_lgtmed", !(lgtmData.IsLgtmed))

		var question models.Question
		database.DB.Where("id = ?", questionID).First(&question)
		if lgtmData.IsLgtmed {
			database.DB.Model(&question).Where("id = ?", questionID).Update("lgtm", question.Lgtm-1)
		} else {
			database.DB.Model(&question).Where("id = ?", questionID).Update("lgtm", question.Lgtm+1)
		}

		return c.JSON(question)
	} else {
		// LGTM情報がない場合
		parent_id, _ := strconv.Atoi(questionID)
		question_id_uint := uint(parent_id)
		lgtm := models.LgtmQuestion{
			QuestionID: question_id_uint,
			LgtmerID:   userID,
			IsLgtmed:   true,
		}
		database.DB.Create(&lgtm)
		var question models.Question
		database.DB.Where("id = ?", questionID).First(&question)
		database.DB.Model(&question).Where("id = ?", questionID).Update("lgtm", question.Lgtm+1)

		return c.JSON(question)
	}
}

func LgtmAnswer(c *fiber.Ctx) error {

	answerID := c.Params("answer_id")

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
	claims := token.Claims.(*utils.Claims)
	// User IDを取得
	userID := claims.Issuer

	// 既にLGTMされているならDBから削除して、LGTMされてないなら新たにDBに加える
	lgtmData := []models.LgtmAnswer{}
	database.DB.Where("lgtmer_id = ?", userID).Where("answer_id = ?", answerID).Find(&lgtmData)

	if len(lgtmData) > 0 {
		database.DB.Where("lgtmer_id = ?", userID).Where("answer_id = ?", answerID).Delete(&lgtmData[0])
	} else {
		parent_id, _ := strconv.Atoi(answerID)
		answer_id_uint := uint(parent_id)
		lgtm := models.LgtmAnswer{
			AnswerID: answer_id_uint,
			LgtmerID: userID,
		}
		database.DB.Create(&lgtm)
	}

	// LGTMの更新
	lgtms := []models.LgtmAnswer{}
	database.DB.Where("answer_id = ?", answerID).Find(&lgtms)
	var answer models.Answer
	database.DB.Model(&answer).Where("id = ?", answerID).Update("lgtm", len(lgtms))

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

	question_id, _ := strconv.Atoi(data["parent_id"])
	question_id_uint := uint(question_id)

	answer := models.Answer{
		QuestionID: question_id_uint,
		AnswererID: data["user_id"],
		Body:       data["body"],
		Lgtm:       0,
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
	answer_id, _ := strconv.Atoi(data["parent_id"])
	answer_id_uint := uint(answer_id)

	reply := models.Reply{
		QuestionID: question_id_uint,
		AnswerID:   answer_id_uint,
		ReplyerID:  data["user_id"],
		Body:       data["body"],
		Lgtm:       0,
	}

	database.DB.Create(&reply)

	return c.JSON(reply)

}
