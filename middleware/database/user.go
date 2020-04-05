package database

import (
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

func QueryUserNextHopByOpenID(openID string) (uint, error) {
	u := new(model.User)
	err := Conn().Select(u.NextHopColumnName()).First(u, u.OpenIDColumnName()+" = ?", openID).Error
	return u.NextHop, err
}

func QueryUserByID(id uint) (*model.User, error) {
	u := new(model.User)
	u.ID = id
	err := Conn().First(u, u.ID).Error
	return u, err
}

func QueryOrCreateUserByOpenID(openID string) (*model.User, error) {
	u := &model.User{
		OpenID:    openID,
		NextHop:   0,
		NickName:  viper.GetString("default_nick_name"),
		Level:     model.Level1,
		EXP:       0,
		AdminRole: false,
	}
	err := Conn().FirstOrCreate(u, u.OpenIDColumnName()+" = ?", openID).Error
	return u, err
}

func QueryUserByOpenID(openID string) (*model.User, error) {
	u := new(model.User)
	err := Conn().First(u, u.OpenIDColumnName()+" = ?", openID).Error
	return u, err
}

func UpdateUserNextHop(u *model.User) error {
	return Conn().Model(u).Update(u.NextHopColumnName(), u.NextHop).Error
}

func UpdateUserNickName(u *model.User) error {
	return Conn().Model(u).Update(u.NickNameColumnName(), u.NickName).Error
}

func UpdatesUser(u *model.User) error {
	return Conn().Model(u).Updates(map[string]interface{}{
		u.NextHopColumnName():  u.NextHop,
		u.NickNameColumnName(): u.NickName,
		u.LevelColumnName():    u.Level,
		u.EXPColumnName():      u.EXP,
	}).Error
}
