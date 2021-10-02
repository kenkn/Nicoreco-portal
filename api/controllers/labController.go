//
// labController.go
// 研究室レビュー
//

package controllers

import (
	"auth-api/database"
	"auth-api/models"

	"github.com/gofiber/fiber/v2"
)

// /lab/reviews/:lab (GET)
// 機能 : 研究室レビューの情報取得
// 戻り値 : 研究室レビューの情報のJSON
func LabReviews(c *fiber.Ctx) error {

	// GETの内容を取得
	lab := c.Params("lab")

	reviews := []models.LabReview{}
	database.DB.Where("lab = ?", lab).Find(&reviews)

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
//  * lab_reviewer_id : レビュー者のユーザID
//  * lab             : レビューする研究室名
//  * body            : レビューの本文
// 戻り値 : 投稿したレビューのJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func PostLabReview(c *fiber.Ctx) error {
	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	review := models.LabReview{
		Lab:           data["lab"],
		LabReviewerID: data["lab_reviewer_id"],
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
func LabReply(c *fiber.Ctx) error {

	// GETの内容を取得
	lab_review_id := c.Params("lab_review_id")

	replys := []models.LabReply{}
	database.DB.Where("lab_review_id = ?", lab_review_id).Find(&replys)

	return c.JSON(replys)
}

// /lab/reply/post (POST)
// 機能 : 研究室レビューのリプライ投稿
// 受信するJSON :
//  * lab_review_id : リプライするレビューID
//  * user_id       : リプライしたユーザーID
//  * body          : レビューの本文
// 戻り値 : 投稿したリプライのJSON
// 例外発行 :
//  * リクエストデータのパースに失敗した場合に例外を発行
func PostLabReply(c *fiber.Ctx) error {
	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	reply := models.LabReply{
		LabReviewID: data["lab_review_id"],
		Body:        data["body"],
		UserID:      data["user_id"],
	}
	// データ登録(CreateはGORMメソッド)
	database.DB.Create(&reply)

	return c.JSON(reply)
}
