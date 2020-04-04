package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	OpenID    string `gorm:"type:varchar(64);unique;index:index_user_open_id_next_hop"`
	NextHop   uint   `gorm:"index:index_user_open_id_next_hop"` // 下一跳，与open_id组成复合索引
	NickName  string `gorm:"type:varchar(8)"`
	Level     uint
	EXP       uint // 经验值
	AdminRole bool // 是否有管理员权限
}

const (
	Level0 = iota
	Level1
	Level2
	Level3
	Level4
	Level5
	Level6
)

func UserTableName() string {
	return "users"
}

func (u *User) OpenIDColumnName() string {
	return "open_id"
}

func (u *User) NextHopColumnName() string {
	return "next_hop"
}

func (u *User) NickNameColumnName() string {
	return "nick_name"
}

func (u *User) LevelColumnName() string {
	return "level"
}

func (u *User) EXPColumnName() string {
	return "exp"
}
