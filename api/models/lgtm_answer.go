//
// LgtmAnswer.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type LgtmAnswer struct {
	gorm.Model
	AnswerID uint   `json:"answer_id"`
	Answer   Answer `gorm:"foreignKey:AnswerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LgtmerID string `json:"lgtmer_id" gorm:"size:256"`
	Lgtmer   User   `gorm:"foreignKey:LgtmerID;references:UserID;"`
	IsLgtmed bool   `json:"is_lgtmed"`
}
