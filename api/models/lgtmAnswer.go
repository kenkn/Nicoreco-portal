//
// LgtmAnswer.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type LgtmAnswer struct {
	gorm.Model
	ParentID uint   `json:"parent_id"`
	UserID   string `json:"user_id"`
}
