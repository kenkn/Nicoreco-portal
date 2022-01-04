//
// lab_reply.go
// 研究室リプライstruct
//

package models

import "gorm.io/gorm"

type LabReply struct {
	gorm.Model
	LabReviewID string    `json:"lab_review_id"`
	LabReview   LabReview `gorm:"foreignKey:LabReviewID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReplyerID   string    `json:"replyer_id" gorm:"size:256"`
	Replyer     User      `gorm:"foreignKey:ReplyerID;references:UserID"`
	Body        string    `json:"body"`
}
