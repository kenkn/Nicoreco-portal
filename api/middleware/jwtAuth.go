package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

// Claimsの型
type Claims struct {
	jwt.StandardClaims
}

// tryAuth(private)
// 機能 : JWTトークンを用いた認証
// 引数 :
//  * tokenString : JWTトークン
// 戻り値 :
// 	* *jwt.Token : JWTトークンで得られたclaim
//  * bool       : 認証成功ならばtrue
func tryAuth(tokenString string) (*jwt.Token, bool) {
	// JWTtoken取得
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return nil, false
	}

	return token, true
}

// GetAuthToken(public)
// 機能 : JWTトークンで得られたclaimの取得
// 引数 :
//  * tokenString : JWTトークン
// 戻り値 :
// 	* *jwt.Token : JWTトークンで得られたclaim
//  * bool       : 認証成功ならばtrue
func GetAuthToken(tokenString string) (*jwt.Token, bool) {
	return tryAuth(tokenString)
}

// CanAuth(public)
// 機能 : JWT認証ができたかどうかの判定
// 引数 :
//  * tokenString : JWTトークン
// 戻り値 : 認証成功ならばtrue
func CanAuth(tokenString string) bool {
	_, canAuth := tryAuth(tokenString)
	return canAuth
}
