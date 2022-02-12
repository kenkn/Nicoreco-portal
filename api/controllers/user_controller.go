//
// userController.go
// ユーザ認証
//

package controllers

import (
	"auth-api/database"
	"auth-api/models"
	"auth-api/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// /user (GET)
// 機能 : ユーザ情報の取得
// 受信するJSON :
//  * jwt : JWTトークン
// 戻り値 :
// 	* 成功時 : 該当ユーザのユーザ情報(JSON)
// 	* 失敗時 : エラー文(JSON)
func User(c *fiber.Ctx) error {

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
	id := claims.Issuer

	var user models.User
	database.DB.Where("user_id = ?", id).First(&user)

	return c.JSON(user)

}

// /user (PUT)
// 機能 : ユーザ情報の更新
// 受信するJSON :
//  * display_name : 更新後の表示名
//  * grade        : 更新後の学年
// 戻り値 : 更新後のユーザ情報のJSON
// 例外発行 :
// 	* リクエストデータのパースに失敗した場合
//  * 認証されていないユーザからのリクエストだった場合
//  * データが欠けている場合
func UpdateUserInfo(c *fiber.Ctx) error {
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
	claims := token.Claims.(*utils.Claims)
	userID := claims.Issuer

	newName, isDisplayNameThere := data["display_name"]
	newGrade, isGradeThere := data["grade"]
	if !isDisplayNameThere || !isGradeThere {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "display_nameまたはgradeが欠けています",
		})
	}

	var user models.User
	database.DB.Model(&user).Where("user_id", userID).Update("display_name", newName).Update("grade", newGrade)

	return c.JSON(user)
}

// /register (POST)
// 機能 : ユーザの登録
// 受信するJSON :
//  * display_name     : ユーザの表示名
//  * user_id		   : ユーザID
//  * email            : 登録するメールアドレス
//  * grade 		   : 学年
//  * password         : ユーザが入力したパスワード
//  * password_confirm : ユーザが入力したパスワード(確認用)
// 戻り値 : ユーザ情報のJSON
// 例外発行 :
// 	* リクエストデータのパースに失敗した場合
func Register(c *fiber.Ctx) error {

	data, err := utils.ParseData(c)
	if err != nil {
		return err
	}

	// IDとパスワードの存在チェック
	_, isDisplayNameThere := data["display_name"]
	_, isUserIDThere := data["user_id"]
	_, isEmailThere := data["email"]
	_, isGradeThere := data["grade"]
	_, isPasswordThere := data["password"]
	_, isPasswordConfirmThere := data["password_confirm"]

	// ID,パスワードの存在チェック
	if !isDisplayNameThere || !isUserIDThere || !isEmailThere || !isGradeThere || !isPasswordThere || !isPasswordConfirmThere {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "IDまたはパスワードが欠けています",
		})
	}

	// ASCIIであるかどうかのチェック
	if !(utils.IsJustifiableUserinfo(data["user_id"], data["password"])) {
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

	data, err := utils.ParseData(c)
	if err != nil {
		return err
	}

	var user models.User
	// emailに紐づくユーザーを取得
	// &userを指定することでDBから取得したデータを直接格納できる
	res := database.DB.Where("email = ?", data["email"]).First(&user)

	// ユーザが見つからなかったとき
	if res.Error != nil {
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
		Issuer:    user.UserID,            // ユーザIDをstringに変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // JWTトークンの有効期限
	}

	// JWT tokenの発行
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Cookieを設定
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"jwt": token,
	})

}

// /logout (GET)
// 機能 : ログアウト機能
// 戻り値 : 成功メッセージ(JSON)
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",                         // tokenを空にする
		Expires:  time.Now().Add(-time.Hour), // マイナス値を入れて期限切れ
		HTTPOnly: true,
	}
 
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "ログアウトに成功しました",
	})
}
