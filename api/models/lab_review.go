//
// lab_review.go
// 研究室レビューstruct
//

package models

import "gorm.io/gorm"

type LabReview struct {
	gorm.Model
	Lab           string `json:"lab"`
	LabReviewerID string `json:"labreviewer_id" gorm:"size:256"`
	LabReviewer   User   `gorm:"foreignKey:LabReviewerID;references:UserID"`
	Body          string `json:"body"`
	Lgtm          uint   `json:"lgtm"`
}
