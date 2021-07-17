//
// user.go
// ユーザstruct
//

package models

import "gorm.io/gorm"

// ユーザ情報
type User struct {
	gorm.Model
	DisplayName string `json:"display_name"`
	Email       string `json:"email" gorm:"unique"`
	Password    []byte `json:"-"`		// 非表示
}