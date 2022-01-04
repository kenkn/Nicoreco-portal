//
// labController.go
// 研究室レビュー
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

// /lab/reviews/:lab (GET)
// 機能 : 研究室レビューの情報取得
// 戻り値 : 研究室レビューの情報のJSON
func GetLabReviews(c *fiber.Ctx) error {

	// GETの内容を取得
	lab := c.Params("lab")

	reviews := []models.LabReview{}
	database.DB.Where("lab = ?", lab).Find(&reviews)
	if len(reviews) == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "質問が見つかりませんでした",
		})
	}

	return c.JSON(reviews)

}

// /lab/review/:id (GET)
// 機能 : 研究室レビューの詳細情報取得
// 戻り値 : 研究室レビューの詳細情報のJSON
func LabReview(c *fiber.Ctx) error {

	// GETの内容を取得
	id := c.Params("id")

	var review models.LabReview
	database.DB.Where("id = ?", id).First(&review)

	return c.JSON(review)

}

// /lab/review/post (POST)
// 機能 : 研究室レビューの投稿
// 受信するJSON :
//  * jwt 	  		  : JWTトークン
//  * user_id : レビュー者のユーザID
//  * lab             : レビューする研究室名
//  * body            : レビューの本文
// 戻り値 : 投稿したレビューのJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func PostLabReview(c *fiber.Ctx) error {

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

	review := models.LabReview{
		Lab:           data["lab"],
		LabReviewerID: data["user_id"],
		Body:          data["body"],
		Lgtm:          0,
	}
	// データ登録(CreateはGORMメソッド)
	database.DB.Create(&review)

	return c.JSON(review)

}

// /lab/reply/:id (GET)
// 機能 : 研究室レビューのリプライ情報取得
// 戻り値 : 研究室レビューのリプライ情報のJSON
func GetLabReply(c *fiber.Ctx) error {

	// GETの内容を取得
	lab_review_id := c.Params("lab_review_id")

	replys := []models.LabReply{}
	database.DB.Where("lab_review_id = ?", lab_review_id).Find(&replys)

	return c.JSON(replys)

}

// /lab/reply/post (POST)
// 機能 : 研究室レビューのリプライ投稿
// 受信するJSON :
//  * jwt 	        : JWTトークン
//  * lab_review_id : リプライするレビューID
//  * user_id       : リプライしたユーザーID
//  * body          : レビューの本文
// 戻り値 : 投稿したリプライのJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
// TODO: CookieのissuerからUser情報を取得
func PostLabReply(c *fiber.Ctx) error {

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

	reply := models.LabReply{
		LabReviewID: data["lab_review_id"],
		Body:        data["body"],
		ReplyerID:   data["user_id"],
	}
	// データ登録(CreateはGORMメソッド)
	database.DB.Create(&reply)

	return c.JSON(reply)

}

// /lgtm/lab/:lab_review_id/:user_id (GET)
// 機能 : LGTMされたどうかの判定(labReview用)
// 戻り値 : LGTM情報のJSON
func IsLabReviewLgtmed(c *fiber.Ctx) error {

	// GETの内容を取得
	lab_review_id := c.Params("lab_review_id")
	user_id := c.Params("user_id")

	// 質問を全検索してリストで取得
	lgtm := []models.LgtmLabReview{}
	database.DB.Where("lgtmer_id = ?", user_id).Where("lab_review_id = ?", lab_review_id).Find(&lgtm)

	return c.JSON(lgtm)

}

// /lgtm/lab (POST)
// 機能 : 研究室レビューのLGTM数を加算する
// 受信するJSON :
//  * jwt 	  		: JWTトークン
//  * lab_review_id : LGTMする研究室レビューのID
//  * user_id 		: LGTMしたユーザーID
// 戻り値 : LGTMした質問のJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func LgtmLabReview(c *fiber.Ctx) error {

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
	lgtmData := []models.LgtmLabReview{}
	database.DB.Where("lgtmer_id = ?", data["user_id"]).Where("lab_review_id = ?", data["lab_review_id"]).Find(&lgtmData)

	if len(lgtmData) > 0 {
		database.DB.Where("lgtmer_id = ?", data["user_id"]).Where("lab_review_id = ?", data["lab_review_id"]).Delete(&lgtmData[0])
	} else {
		parent_id, _ := strconv.Atoi(data["lab_review_id"])
		lab_review_id_uint := uint(parent_id)
		lgtm := models.LgtmLabReview{
			LabReviewID: lab_review_id_uint,
			LgtmerID:    data["user_id"],
		}
		database.DB.Create(&lgtm)
	}

	// LGTMの更新
	lgtms := []models.LgtmLabReview{}
	database.DB.Where("lab_review_id = ?", data["lab_review_id"]).Find(&lgtms)
	var labReview models.LabReview
	database.DB.Model(&labReview).Where("id = ?", data["lab_review_id"]).Update("lgtm", len(lgtms))

	return c.JSON(labReview)

}
