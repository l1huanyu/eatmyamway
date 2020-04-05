package model

import "github.com/jinzhu/gorm"

type Amway struct {
	gorm.Model
	UserID        uint   `gorm:"index:index_amway_user_id_valid;unique_index:unique_index_amway_user_id_name"` // 主动营业用户ID
	HP            int    // 生命值
	Valid         bool   `gorm:"index:index_amway_user_id_valid"`                               // 生命值 > 0 时有效
	Name          string `gorm:"type:varchar(20);unique_index:unique_index_amway_user_id_name"` // 主体名字
	Type          string `gorm:"type:varchar(20)"`                                              // 主要类型
	MarketingCopy string `gorm:"type:varchar(140)"`                                             // 营销文案
	FakePortal    string `gorm:"type:varchar(100)"`                                             // 【伪】传送门
	Version       uint   // 版本号
}

func AmwayTableName() string {
	return "amways"
}

func (a *Amway) HPColumnName() string {
	return "hp"
}

func (a *Amway) ValidColumnName() string {
	return "valid"
}

func (a *Amway) NameColumnName() string {
	return "name"
}

func (a *Amway) TypeColumnName() string {
	return "type"
}

func (a *Amway) MarketingCopyColumnName() string {
	return "marketing_copy"
}

func (a *Amway) FakePortalColumnName() string {
	return "fake_portal"
}

func (a *Amway) VersionColumnName() string {
	return "version"
}
