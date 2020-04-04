package scheduler

import (
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
)

// 根据openID查找用户
func queryUser(node *Node) {
	u, err := database.QueryUserByOpenID(node.OpenID)
	if err != nil {
		log.Error("queryUser.database.QueryOrCreateUserByOpenID", err.Error(), map[string]interface{}{"node.OpenID": node.OpenID})
		return
	}
	node.curuser = u
}

// 更新下一跳
func updateUserNextHop(node *Node) {
	err := database.UpdateUserNextHop(node.curuser)
	if err != nil {
		log.Error("updateUserNextHop.database.UpdateUserNextHop", err.Error(), map[string]interface{}{"userID": node.curuser.ID})
		return
	}
}

// 更新用户信息
func updatesUser(node *Node) {
	err := database.UpdatesUser(node.curuser)
	if err != nil {
		log.Error("updatesUser.database.UpdatesUser", err.Error(), map[string]interface{}{"userID": node.curuser.ID})
		return
	}
}
