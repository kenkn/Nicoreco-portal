//
// lgtmQuestion.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type LgtmQuestion struct {
	gorm.Model
	QuestionID uint   `json:"question_id"`
	UserID     string `json:"user_id"`
}
