package model

import "github.com/jinzhu/gorm"

type Record struct {
	gorm.Model
	UserID  uint
	AmwayID uint
	Like    bool
}
