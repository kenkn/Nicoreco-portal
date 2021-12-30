//
// lab_review.go
// 研究室レビューstruct
//

package models

import "gorm.io/gorm"

type LabReview struct {
	gorm.Model
	Lab           string `json:"lab"`
	LabReviewerID string `json:"lab_reviewer_id"`
	Body          string `json:"body"`
	Lgtm          uint   `json:"lgtm"`
}
