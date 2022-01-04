//
// question.go
// 回答struct
//

package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	QuestionID uint     `json:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AnswererID string   `json:"answerer_id" gorm:"size:256"`
	Answerer   User     `gorm:"foreignKey:AnswererID;references:UserID"`
	Body       string   `json:"body"`
	Lgtm       uint     `json:"lgtm"`
}
