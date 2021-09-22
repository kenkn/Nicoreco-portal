//
// lgtm.go
// LGTMstruct
//

package models

import "gorm.io/gorm"

type Lgtm struct {
	gorm.Model
	ParentID uint   `json:"parent_id"`
	UserID   string `json:"user_id"`
}
