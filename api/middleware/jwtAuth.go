package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

// Claimsの型
type Claims struct {
	jwt.StandardClaims
}

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

func GetAuthToken(tokenString string) (*jwt.Token, bool) {
	return tryAuth(tokenString)
}

func CanAuth(tokenString string) bool {
	_, canAuth := tryAuth(tokenString)
	return canAuth
}
