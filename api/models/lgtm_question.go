//
// lgtmQuestion.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type LgtmQuestion struct {
	gorm.Model
	QuestionID uint     `json:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LgtmerID   string   `json:"lgtmer_id" gorm:"size:256"`
	Lgtmer     User     `gorm:"foreignKey:LgtmerID;references:UserID;"`
	IsLgtmed   bool     `json:"is_lgtmed"`
}
