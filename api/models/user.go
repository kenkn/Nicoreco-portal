//
// user.go
// ユーザstruct
//

package models

import "gorm.io/gorm"

// ユーザ情報
type User struct {
	gorm.Model
	UserID      string `json:"user_id" gorm:"unique"`
	DisplayName string `json:"display_name"`
	Grade       string `json:"grade"`
	Email       string `json:"email" gorm:"unique"`
	Password    []byte `json:"-"` // 非表示
}
