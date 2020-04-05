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
	Level1 = iota + 1 // 获得 0 EXP
	Level2            // 获得 200 EXP
	Level3            // 获得 1500 EXP
	Level4            // 获得 4500 EXP
	Level5            // 获得 10800 EXP
	Level6            // 获得 28800 EXP
)

// 经验值设定：发布安利EXP+10，互动（喜欢，不喜欢）EXP+5，发布的安利被喜欢+1
var levelEXPList = [...]uint{0, 0, 200, 1500, 4500, 10800, 28800}

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

// 获得经验值
func (u *User) GainEXP(exp uint) {
	u.EXP += exp
	u.levelUp()
}

func (u *User) levelUp() {
	// 满级直接返回
	if u.Level >= Level6 {
		return
	}

	if u.EXP >= levelEXPList[u.Level+1] {
		u.Level++
	}
}
