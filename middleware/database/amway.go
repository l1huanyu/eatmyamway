package database

import (
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

// 按照生命值优先级随机读取N条记录
func QueryAmwayRand(limit int) ([]*model.Amway, error) {
	amways := []*model.Amway{}
	err := Conn().Where("valid = ?", 1).Order("rand()").Limit(limit).Find(&amways).Error
	return amways, err
}

func CreateAmway(userID, userLevel uint) error {
	a := &model.Amway{
		Valid:   false,
		UserID:  userID,
		HP:      viper.GetInt("default_amway_hp") + viper.GetInt("user_level_weight")*int(userLevel),
		Version: 1,
	}

	return Conn().Create(a).Error
}

func QueryInvalidAmway(userID uint) (*model.Amway, error) {
	a := new(model.Amway)
	err := Conn().Where("user_id = ?", userID).First(a, "valid = ?", false).Error
	return a, err
}

func UpdatesAmway(a *model.Amway) error {
	return Conn().Where(a.VersionColumnName()+" = ?", a.Version).Model(a).Updates(map[string]interface{}{
		a.ValidColumnName():         a.Valid,
		a.NameColumnName():          a.Name,
		a.TypeColumnName():          a.Type,
		a.MarketingCopyColumnName(): a.MarketingCopy,
		a.FakePortalColumnName():    a.FakePortal,
		a.VersionColumnName():       a.Version + 1,
	}).Error
}

func DeleteAmway(a *model.Amway) error {
	return Conn().Delete(a).Error
}
