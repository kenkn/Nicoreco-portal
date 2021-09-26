//
// LgtmAnswer.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type LgtmAnswer struct {
	gorm.Model
	AnswerID uint   `json:"answer_id"`
	UserID   string `json:"user_id"`
}
