package models

type PasswordReset struct {
	ID 	  uint
	Email string
	Token string
}