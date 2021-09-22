//
// lgtmQuestion.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type LgtmQuestion struct {
	gorm.Model
	ParentID uint   `json:"parent_id"`
	UserID   string `json:"user_id"`
}
