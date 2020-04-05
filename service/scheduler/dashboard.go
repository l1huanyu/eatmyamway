package scheduler

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
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

	a, err := database.QueryAmwayRand()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			node.curuser.NextHop = _NodeDashboard
			node.Content = viper.GetString("not_found")
		} else {
			log.Error("queryAmwayRand.database.QueryAmwayRand", err.Error(), nil)
		}
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
