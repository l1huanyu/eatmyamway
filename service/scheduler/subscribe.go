package scheduler

import (
	"fmt"

	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/cache"
	"github.com/l1huanyu/eatmyamway/middleware/database"
)

// 仅在用户订阅时使用该action，下一跳为 _NodeSelectUpdateNickNameOrDashboard
func queryOrCreateUser(node *Node) {
	u, err := database.QueryOrCreateUserByOpenID(node.OpenID)
	if err != nil {
		log.Error("queryOrCreateUser.database.QueryOrCreateUserByOpenID", err.Error(), map[string]interface{}{"node.OpenID": node.OpenID})
		return
	}
	node.curuser = u

	// 为活跃用户新建缓存
	key := fmt.Sprintf("curuserid%d", u.ID)
	if _, found := cache.Get(key); !found {
		cache.Set(key, map[uint]struct{}{})
	}

	u.NextHop = _NodeSelectUpdateNickNameOrDashboard
	node.Content = fmt.Sprintf(node_subscribe+node_select_update_nick_name_or_dashboard, node.curuser.NickName)
}
