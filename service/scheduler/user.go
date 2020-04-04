package scheduler

import (
	"strconv"

	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/spf13/viper"
)

// 选择更新昵称或返回主界面，下一跳 _NodeUpdateNickName/_NodeDashboard
func selectUpdateNickNameOrDashboard(node *Node) {
	option, err := strconv.Atoi(node.Msg)
	if err != nil {
		log.Error("selectUpdateNickNameOrDashboard.strconv.Atoi", err.Error(), map[string]interface{}{"node.Msg": node.Msg})
		return
	}

	switch option {
	case 1:
		node.curuser.NextHop = _NodeUpdateNickName
		node.Content = viper.GetString("node_update_nick_name")
	case 2:
		node.curuser.NextHop = _NodeDashboard
		node.Content = viper.GetString("node_dashboard")
	default:
		log.Error("selectUpdateNickNameOrDashboard", "invalid option", map[string]interface{}{"option": option})
		return
	}
}

// 更新昵称，下一跳 _NodeDashboard
func updateUserNickName(node *Node) {
	node.curuser.NickName = node.Content
	err := database.UpdateUserNickName(node.curuser)
	if err != nil {
		log.Error("updateUserNickName.database.UpdateUserNickName", err.Error(), map[string]interface{}{"node.Content": node.Content})
		return
	}

	node.curuser.NextHop = _NodeDashboard
	node.Content = viper.GetString("node_dashboard")
}
