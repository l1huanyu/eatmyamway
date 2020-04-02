package model

import "github.com/jinzhu/gorm"

type Amway struct {
	gorm.Model
	UserID        uint   `gorm:"index:index_amway_user_id"` // 主动营业用户ID
	HP            uint   `gorm:"index:index_amway_hp"`      // 生命值
	Valid         bool   // 生命值 > 0 时有效
	Name          string `gorm:"type:varchar(16)"` // 主体名字
	Type          uint   // 主要类型
	FakePortal    string `gorm:"type:varchar(255)"` // 【伪】传送门
	MarketingCopy string `gorm:"type:varchar(255)"` // 营销文案
	Version       uint   // 版本号
}
