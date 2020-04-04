package scheduler

import (
	"fmt"

	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/spf13/viper"
)

// 仅在用户订阅时使用该action，下一跳为 _NodeSelectUpdateNickNameOrDashboard
func queryOrCreateUser(node *Node) {
	u, err := database.QueryOrCreateUserByOpenID(node.OpenID)
	if err != nil {
		log.Error("queryOrCreateUser.database.QueryOrCreateUserByOpenID", err.Error(), map[string]interface{}{"node.OpenID": node.OpenID})
		return
	}
	node.curuser = u

	u.NextHop = _NodeSelectUpdateNickNameOrDashboard
	node.Content = fmt.Sprintf("%s\n\n%s", viper.GetString("node_subscribe"), viper.GetString("node_select_update_nick_name_or_dashboard"), node.curuser.NickName)
}
