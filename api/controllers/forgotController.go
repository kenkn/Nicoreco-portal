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

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := RandStringRunes(12)
	passwordReset := models.PasswordReset {
		Email : data["email"],
		Token : token,
	}

	// DBに保存
		database.DB.Create(&passwordReset)
 
	// SMTPメール送信
	from := "skt7tp@gmail.com"
	to := []string{
		data["email"],
	}
	sendFrom := fmt.Sprintf("From: %s\n", from)
	subject := fmt.Sprintf("Subject; %s\n", "Password Reset")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// Vue.jsのアドレス
	url := "http://localhost:8080/reset/" + token
	message := fmt.Sprintf("Click <a href=\"%s\">here</a> to reset password!", url)
	err := smtp.SendMail(
		"smtp:1025", // コンテナサービス名+port
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

// ランダム文字列を返す関数
func RandStringRunes(n int) string {
	var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// Unicode文字列を16進数から10進数に変換してbに格納
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRunes[rand.Intn(len(lettersRunes))]
	}
	return string(b)
}

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
	// JWT Tokenからデータを取得
	err := database.DB.Where("token = ?", data["token"]).Last(&passwordReset)
	if err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message" : "違法なトークンです",
		})
	}

	// パスワードをエンコード
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message" : "成功.",
	})
}