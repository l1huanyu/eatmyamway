package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	OpenID   string
	NickName string
	Level    uint
	EXP      uint // 经验值
	Version  uint // 版本号
}
