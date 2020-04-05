package scheduler

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/cache"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

const (
	_DashboardQueryAmway = iota + 1
	_DashboardCreateAmway
	_DashboardPersonalInterface
)

// 主界面
func dashboard(node *Node) {
	option, err := strconv.Atoi(node.Msg)
	if err != nil {
		log.Error("dashboard.strconv.Atoi", err.Error(), map[string]interface{}{"node.Msg": node.Msg})
		return
	}

	node.ctx = option
}

// 随机查找效amway记录，下一跳_NodeSelectInteractOrQueryAmway
func queryAmwayRand(node *Node) {
	if node.ctx.(int) != _DashboardQueryAmway {
		return
	}

	amways, err := database.QueryAmwayRand(viper.GetInt("query_amway_rand_limit"))
	if len(amways) == 0 {
		node.curuser.NextHop = _NodeDashboard
		node.Content = viper.GetString("not_found")
		return
	}
	if err != nil {
		log.Error("queryAmwayRand.database.QueryAmwayRand", err.Error(), nil)
		return
	}

	var a *model.Amway
	key := fmt.Sprintf("curuserid%d", node.curuser.ID)
	if read, found := cache.Get(key); found {
		for i := range amways {
			if _, ok := read.(map[uint]struct{})[amways[i].ID]; !ok {
				a = amways[i]
				break
			}
		}
	} else {
		// 缓存过期随机返回
		a = amways[rand.Intn(len(amways))]
	}

	// 这一批全部已阅
	if a.Valid == false {
		node.curuser.NextHop = _NodeDashboard
		node.Content = viper.GetString("not_found")
		return
	}

	// 更新或创建聚焦关系
	err = database.UpdateOrCreateFocusRelation(node.curuser.ID, a.ID)
	if err != nil {
		log.Error("queryAmwayRand.database.UpdateOrCreateFocusRelation", err.Error(), map[string]interface{}{"node.curuser.ID": node.curuser.ID, "a.ID": a.UserID})
		return
	}

	au, err := database.QueryUserByID(a.UserID)
	if err != nil {
		log.Error("queryAmwayRand.database.QueryUserByID", err.Error(), map[string]interface{}{"a.UserID": a.UserID})
		return
	}

	node.curuser.NextHop = _NodeSelectInteractOrQueryAmway
	node.Content += fmt.Sprintf(amway_information+"\n"+node_select_interact_or_query_amway,
		au.NickName, au.Level, au.ID, a.Name, a.Type, a.HP, a.MarketingCopy, a.FakePortal)
}

func createAmway(node *Node) {
	if node.ctx.(int) != _DashboardCreateAmway {
		return
	}

	err := database.CreateAmway(node.curuser.ID, node.curuser.Level)
	if err != nil {
		log.Error("createAmway.database.CreateAmway", err.Error(), map[string]interface{}{"node.curuser.ID": node.curuser.ID})
		return
	}

	node.curuser.NextHop = _NodeCreateAndUpdateAmwayName
	node.Content = node_create_and_update_amway_name
}

func personalInterface(node *Node) {
	if node.ctx.(int) != _DashboardPersonalInterface {
		return
	}
	node.Content = viper.GetString("not_support")
}
