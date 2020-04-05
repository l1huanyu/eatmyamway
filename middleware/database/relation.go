package database

import (
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

func UpdateOrCreateFocusRelation(userID, amwayID uint) error {
	r := &model.Relation{
		UserID:          userID,
		Focus:           true,
		AmwayID:         amwayID,
		InteractionType: model.InteractionNull,
	}
	tx := Conn().Begin()
	err := tx.FirstOrCreate(r, r.UserIDColumnName()+" = ? and "+r.AmwayIDColumnName()+" = ?", userID, amwayID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(r).Update(r.FocusColumnName, true).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func QueryUserFocusAmwayID(userID uint) (uint, error) {
	r := new(model.Relation)
	err := Conn().Select(r.AmwayIDColumnName()).First(r, r.UserIDColumnName()+" = ? and "+r.FocusColumnName()+" = ?", userID, true).Error
	return r.AmwayID, err
}

func QueryFocusRelation(userID uint) (*model.Relation, error) {
	r := new(model.Relation)
	err := Conn().First(r, r.UserIDColumnName()+" = ? and "+r.FocusColumnName()+" = ?", userID, true).Error
	return r, err
}

func UpdateFocusRelation(r *model.Relation) error {
	return Conn().Model(r).Update(r.FocusColumnName, r.Focus).Error
}

func UpdateInteractionRelation(r *model.Relation) error {
	tx := Conn().Begin()
	err := tx.Model(r).Update(r.InteractionTypeColumnName(), r.InteractionType).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	u := new(model.User)
	err = tx.First(u, r.UserID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	a := new(model.Amway)
	err = tx.First(a, r.AmwayID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 用户获得互动经验值奖励
	u.GainEXP(viper.GetUint("exp_interact_amway"))

	switch r.InteractionType {
	case model.InteractionLike:
		// 当前用户不是聚焦安利的推送者
		if a.UserID != u.ID {
			au := new(model.User)
			err = tx.First(au, a.UserID).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			// 获取被点赞经验值奖励
			au.GainEXP(viper.GetUint("exp_somebody_like_my_amway"))
			// 更新数据
			err = tx.Model(au).Updates(map[string]interface{}{
				au.EXPColumnName():   au.EXP,
				au.LevelColumnName(): au.Level,
			}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			// 可以点赞自己的安利
			u.GainEXP(viper.GetUint("exp_somebody_like_my_amway"))
		}

		// 点赞增加安利的生命值
		a.HP += int(u.Level) * viper.GetInt("user_level_weight")
	case model.InteractionDislike:
		// 点踩减少安利的生命值
		a.HP -= int(u.Level) * viper.GetInt("user_level_weight")
		if a.HP <= 0 {
			// 生命值归0直接删除
			err = tx.Delete(a).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 更新用户
	err = tx.Model(u).Updates(map[string]interface{}{
		u.EXPColumnName():   u.EXP,
		u.LevelColumnName(): u.Level,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新安利
	err = tx.Model(a).Update(a.HPColumnName, a.HP).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
