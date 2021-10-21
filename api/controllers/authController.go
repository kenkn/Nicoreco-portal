//
// authController.go
// ユーザ認証
//

package controllers

import (
	"auth-api/database"
	"auth-api/middleware"
	"auth-api/models"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// checkPattern(private)
// 機能 : ID，パスワードの正当性チェック
// 引数 :
//  * id   : 検証したいID
//  * pass : 検証したいパスワード
// 戻り値 : 正しいID，パスワードならばtrue
func checkPattern(id, pass string) bool {
	// ASCIIかどうかのチェック
	if !(utf8.ValidString(id) && utf8.RuneCountInString(id) == len(id)) ||
		!(utf8.ValidString(pass) && utf8.RuneCountInString(pass) == len(pass)) {
		return false
	}

	// idに空白文字もしくは制御文字があるかどうかのチェック
	for _, c := range id {
		if c == ' ' {
			return false
		}
		if unicode.IsControl(c) {
			return false
		}
	}

	// パスワードに空白文字もしくは制御文字があるかどうかのチェック
	for _, c := range pass {
		if c == ' ' {
			return false
		}
		if unicode.IsControl(c) {
			return false
		}
	}

	return true
}

// /user (POST)
// 機能 : ユーザ情報の取得
// 受信するJSON :
//  * jwt : JWTトークン
// 戻り値 :
// 	* 成功時 : 該当ユーザのユーザ情報(JSON)
// 	* 失敗時 : エラー文(JSON)
func User(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// JWTの認証
	token, canAuth := middleware.GetAuthToken(data["jwt"])
	if !canAuth {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "認証されていません．",
		})
	}

	claims := token.Claims.(*(middleware.Claims))
	// User IDを取得
	id := claims.Issuer

	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)

}

// /register (POST)
// 機能 : ユーザの登録
// 受信するJSON :
//  * display_name     : ユーザの表示名
//  * email            : 登録するメールアドレス
//  * password         : ユーザが入力したパスワード
//  * password_confirm : ユーザが入力したパスワード(確認用)
// 戻り値 : ユーザ情報のJSON(User型)
// 例外発行 :
// 	* リクエストデータのパースに失敗した場合に例外を発行
func Register(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// IDとパスワードの存在チェック
	_, isDisplayNameThere := data["display_name"]
	_, isUserIDThere := data["user_id"]
	_, isPasswordThere := data["password"]
	_, isPasswordConfirmThere := data["password_confirm"]

	// ID,パスワードの存在チェック
	if !isDisplayNameThere || !isUserIDThere || !isPasswordThere || !isPasswordConfirmThere {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "IDまたはパスワードが欠けています",
		})
	}

	// ASCIIであるかどうかのチェック
	if !checkPattern(data["user_id"], data["password"]) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "IDとパスワードはASCIIでスペースを用いてはいけません",
		})
	}

	// IDの長さのチェック
	if !(3 <= len(data["user_id"])) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "ユーザIDは3文字以上で入力してください",
		})
	}

	// パスワードの長さのチェック
	if !(8 <= len(data["password"])) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "パスワードは8文字以上で入力してください",
		})
	}

	// パスワードチェック
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "パスワードが違います",
		})
	}

	// パスワードをエンコード(暗号の強度: 10)
	// 暗号の強度は，高いほどセキュリティ性は高いがパフォーマンス低下に繋がる
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	user := models.User{
		DisplayName: data["display_name"],
		UserID:      data["user_id"],
		Grade:       data["grade"],
		Email:       data["email"],
		Password:    password,
	}

	// データ登録(CreateはGORMメソッド)
	database.DB.Create(&user)

	return c.JSON(user)

}

// /login (POST)
// 機能 : ログイン機能
// 受信するJSON :
// 	* email    : ユーザが入力したメールアドレス
// 	* password : ユーザが入力したパスワード
// 戻り値 :
// 	* 成功時 : 該当ユーザのJWTトークン(JSON)
// 	* 失敗時 : ステータスコード404とエラー文(JSON)
// 例外処理 :
// 	* リクエストデータのパースに失敗した場合に例外を発行
// 	* JSONの内容呼び出しに失敗した場合に例外を発行
func Login(c *fiber.Ctx) error {

	var data map[string]string
	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	// emailに紐づくユーザーを取得
	// &userを指定することでDBから取得したデータを直接格納できる
	database.DB.Where("email = ?", data["email"]).First(&user)

	// ユーザが見つからなかったとき
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "メールアドレスまたはユーザ名が違います",
		})
	}

	// パスワードが間違っていた場合
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			// messageの内容はユーザが見つからなかった時と一緒にする(安全性確保のため)
			"message": "メールアドレスまたはユーザ名が違います",
		})
	}

	// JWT Claimsの発行
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),            // ユーザIDをstringに変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // JWTトークンの有効期限
	}

	// JWT tokenの発行
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"jwt": token,
	})

}
