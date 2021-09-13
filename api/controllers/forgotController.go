//
// forgotController.go
// パスワードを忘れたときのリセット
//

package controllers

import (
	"auth-api/database"
	"auth-api/models"
	"math/rand"
	"fmt"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// /forgot (POST)
// 機能 : パスワードリセット用URLを送信
// 受信するJSON :
//  * email : 入力されたメールアドレス
// 戻り値 : 成功メッセージ(JSON)
// 例外処理 : 
// 	* リクエストデータのパースに失敗した場合に例外を発行
//	* メール送信失敗時に例外を発行
// TODO :
// 	Tokenの被りをはじく
func Forgot(c *fiber.Ctx) error {
	
	RandStringWordCount := 12
	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// 12文字のランダム文字列を生成
	token := RandStringRunes(RandStringWordCount)
	passwordReset := models.PasswordReset {
		Email : data["email"],
		Token : token,
	}

	// DBに保存
	database.DB.Create(&passwordReset)
 
	// メール受信側の情報
	// TODO 送信元アドレスの変更
	from := "skt7tp@gmail.com"
	to := []string{
		data["email"],
	}
	sendFrom := fmt.Sprintf("From: %s\n", from)
	subject := fmt.Sprintf("Subject; %s\n", "Password Reset")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// パスワードリセット用のURL
	// TODO サーバにデプロイする際に変更
	url := "http://localhost:8080/reset/" + token
	message := fmt.Sprintf("Click <a href=\"%s\">here</a> to reset password!", url)
	err := smtp.SendMail(
		"smtp:1025",
		nil,
		from,
		to,
		[]byte(sendFrom + subject + mime + message),
	)
 
	if err != nil {
		return err
	}
 
	return c.JSON(fiber.Map{
		"message": "成功",
	})

}

// RandStringRunes
// 機能 : n文字の英字ランダム文字列を返す
// 戻り値 : n文字のランダム文字列
func RandStringRunes(n int) string {

	var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// Unicode文字列を16進数から10進数に変換してbに格納
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRunes[rand.Intn(len(lettersRunes))]
	}
	return string(b)

}

// /reset (POST)
// 機能 : パスワードのリセット
// 受信するJSON :
//	* token            : リセット用トークン(/reset/<token>)
//  * password         : ユーザが入力したパスワード
//  * password_confirm : ユーザが入力したパスワード(確認用)
// 戻り値 :
// 	* 成功時 : 成功メッセージ(JSON)
// 	* 失敗時 : ステータスコード400とエラー文(JSON)
// 例外処理 : 
// 	* リクエストデータのパースに失敗した場合に例外を発行
func Reset(c *fiber.Ctx) error {

	var data map[string]string

	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// パスワードのチェック
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message" : "パスワードが一致しません．",
		})
	}

	var passwordReset = models.PasswordReset{}
	// Tokenからデータを取得
	err := database.DB.Where("token = ?", data["token"]).Last(&passwordReset)
	// 未発行トークンだった場合
	if err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message" : "違法なトークンです",
		})
	}

	// パスワードをエンコード
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	// DBを更新
	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message" : "成功.",
	})

}