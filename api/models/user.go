//
// user.go
// ユーザstruct
//

package models

import "gorm.io/gorm"

// ユーザ情報
type User struct {
	gorm.Model
	DisplayName string `json:"display_name" sql:"type:VARCHAR(30) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	UserID      string `json:"user_id" gorm"unique"`
	Grade       string `json:"grade sql:"type:VARCHAR(30) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Email       string `json:"email" gorm:"unique"`
	Password    []byte `json:"-"` // 非表示
}
