//
// question.go
// 質問struct
//

package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuestionerID string `json:"questioner_id" gorm:"size:256"`
	Questioner   User   `gorm:"foreignKey:QuestionerID;references:UserID"`
	Subject      string `json:"subject"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	Lgtm         uint   `json:"lgtm"`
	AnswerCount  uint   `json:"answer_count"`
}
