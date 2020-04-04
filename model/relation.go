package model

import "github.com/jinzhu/gorm"

type Relation struct {
	gorm.Model
	UserID  uint `gorm:"index:index_relation_user_id"`
	AmwayID uint `gorm:"index:index_relation_amway_id"`
	Like    bool
}

func RelationTableName() string {
	return "relations"
}
