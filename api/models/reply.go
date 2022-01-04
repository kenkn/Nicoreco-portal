//
// reply.go
// 返信struct
//

package models

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	QuestionID uint     `json:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AnswerID   uint     `json:"answer_id"`
	Answer     Answer   `gorm:"foreignKey:AnswerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReplyerID  string   `json:"replyer_id" gorm:"size:256"`
	Replyer    User     `gorm:"foreignKey:ReplyerID;references:UserID"`
	Body       string   `json:"body"`
	Lgtm       uint     `json:"lgtm"`
}
