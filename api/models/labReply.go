//
// lab_reply.go
// 研究室リプライstruct
//

package models

import "gorm.io/gorm"

type LabReply struct {
	gorm.Model
	LabReviewID string `json:"lab_review_id"`
	UserID      string `json:"user_id"`
	Body        string `json:"body"`
}
