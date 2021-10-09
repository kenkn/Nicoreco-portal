//
// lgtmLabReview.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type LgtmLabReview struct {
	gorm.Model
	LabReviewID uint   `json:"lab_review_id"`
	UserID      string `json:"user_id"`
}
