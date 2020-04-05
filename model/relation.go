package model

import "github.com/jinzhu/gorm"

type Relation struct {
	gorm.Model
	UserID          uint `gorm:"unique_index:unique_index_relation_user_id_focus_amway_id"`
	Focus           bool `gorm:"unique_index:unique_index_relation_user_id_focus_amway_id"` // 当前用户是否聚焦该关系
	AmwayID         uint `gorm:"unique_index:unique_index_relation_user_id_focus_amway_id"`
	InteractionType uint // 互动类型
}

const (
	InteractionNull = iota
	InteractionLike
	InteractionDislike
)

func RelationTableName() string {
	return "relations"
}

func (r *Relation) UserIDColumnName() string {
	return "user_id"
}

func (r *Relation) FocusColumnName() string {
	return "focus"
}

func (r *Relation) AmwayIDColumnName() string {
	return "amway_id"
}

func (r *Relation) InteractionTypeColumnName() string {
	return "interaction_type"
}
