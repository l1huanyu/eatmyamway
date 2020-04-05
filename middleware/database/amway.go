package database

import (
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

func QueryAmwayRand() (*model.Amway, error) {
	a := new(model.Amway)
	err := Conn().Where(a.ValidColumnName()+" = ?", true).First(a, "id >= ((SELECT MAX(id) FROM amways)-(SELECT MIN(id) FROM amways)) * RAND() + (SELECT MIN(id) FROM amways)").Error
	return a, err
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
