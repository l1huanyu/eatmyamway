package model

import "github.com/jinzhu/gorm"

type Record struct {
	gorm.Model
	UserID  uint `gorm:"index:index_record_user_id"`
	AmwayID uint `gorm:"index:index_record_amway_id"`
	Like    bool
}
