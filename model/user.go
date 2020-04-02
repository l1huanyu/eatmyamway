package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	OpenID    string `gorm:"type:varchar(100);index:index_user_open_id"`
	NickName  string `gorm:"type:varchar(8)"`
	Level     uint
	EXP       uint // 经验值
	AdminRole bool // 是否有管理员权限
	Version   uint // 版本号
}
