//
// reply.go
// 返信struct
//

package models

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	QuestionID uint   `json:"question_id"`
	ParentID   uint   `json:"parent_id"`
	UserID     string `json:"user_id"`
	Body       string `json:"body"`
	Lgtm       uint   `json:"lgtm"`
}
