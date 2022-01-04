//
// lgtmLabReview.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type LgtmLabReview struct {
	gorm.Model
	LabReviewID uint      `json:"lab_review_id"`
	LabReview   LabReview `gorm:"foreignKey:LabReviewID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LgtmerID    string    `json:"lgtmer_id" gorm:"size:256"`
	Lgtmer      User      `gorm:"foreignKey:LgtmerID;references:UserID;"`
}
