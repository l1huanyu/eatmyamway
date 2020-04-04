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

// 查询生命值最高的有效amway记录，下一跳_NodeSelectInteractOrQueryAmway
func queryAmwayWithMaxHP(node *Node) {
	if node.ctx.(int) != _DashboardQueryAmway {
		return
	}

	a, err := database.QueryAmwayRand()
	if err == gorm.ErrRecordNotFound {
		node.curuser.NextHop = _NodeDashboard
		node.Content = viper.GetString("not_found")
	} else {
		log.Error("queryAmwayWithMaxHP.database.QueryAmwayRand", err.Error(), nil)
		return
	}

	au, err := database.QueryUserByID(a.UserID)
	if err != nil {
		log.Error("queryAmwayWithMaxHP.database.QueryUserByID", err.Error(), map[string]interface{}{"a.UserID": a.UserID})
		return
	}

	node.curuser.NextHop = _NodeSelectInteractOrQueryAmway
	node.Content = fmt.Sprintf(viper.GetString("amway_information")+"\n\n"+viper.GetString("node_select_interact_or_query_amway"),
		au.Level, au.NickName, au.ID, a.Name, a.HP, a.Type, a.MarketingCopy, a.FakePortal)
}

func createAmway(node *Node) {
	if node.ctx.(int) != _DashboardCreateAmway {
		return
	}
}

func personalInterface(node *Node) {
	if node.ctx.(int) != _DashboardPersonalInterface {
		return
	}
	node.Content = viper.GetString("not_support")
}
