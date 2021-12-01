package controllers

import "github.com/dgrijalva/jwt-go"

// Claimsの型
type Claims struct {
	jwt.StandardClaims
}
