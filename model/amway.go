package model

import "github.com/jinzhu/gorm"

type Amway struct {
	gorm.Model
	UserID        uint   // 主动营业用户ID
	HP            uint   // 生命值
	Valid         bool   // 生命值 > 0 时有效
	Name          string // 主体名字
	Type          string // 主要类型
	FakePortal    string // 【伪】传送门
	MarketingCopy string // 营销文案
	Version       uint   // 版本号
}
