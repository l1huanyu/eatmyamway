package scheduler

import (
	"fmt"

	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

func queryInvalidAmway(node *Node) {
	a, err := database.QueryInvalidAmway(node.curuser.ID)
	if err != nil {
		log.Error("queryInvalidAmway.database.QueryInvalidAmway", err.Error(), map[string]interface{}{"node.curuser.ID": node.curuser.ID})
	}

	node.ctx = a
}

func updateAmwayName(node *Node) {
	if node.ctx == nil {
		return
	}

	a := node.ctx.(*model.Amway)
	switch node.Msg {
	case "1":
		// 返回
		err := database.DeleteAmway(a)
		if err != nil {
			log.Error("updateAmwayName.database.DeleteAmway", err.Error(), map[string]interface{}{"a.ID": a.ID})
			return
		}
		node.curuser.NextHop = _NodeDashboard
		node.Content = fmt.Sprintf(node_dashboard, node.curuser.NickName, node.curuser.Level, node.curuser.ID)
	default:
		a.Name = node.Msg
		err := database.UpdatesAmway(a)
		if err != nil {
			log.Error("updateAmwayName.database.UpdatesAmway", err.Error(), map[string]interface{}{"a.Name": a.Name})
			return
		}
		node.curuser.NextHop = _NodeUpdateAmwayType
		node.Content = node_update_amway_type
	}
}

func updateAmwayType(node *Node) {
	if node.ctx == nil {
		return
	}

	a := node.ctx.(*model.Amway)
	switch node.Msg {
	case "1":
		// 返回
		node.curuser.NextHop = _NodeCreateAndUpdateAmwayName
		node.Content = node_create_and_update_amway_name
	default:
		a.Type = node.Msg
		err := database.UpdatesAmway(a)
		if err != nil {
			log.Error("updateAmwayType.database.UpdatesAmway", err.Error(), map[string]interface{}{"a.Type": a.Type})
			return
		}
		node.curuser.NextHop = _NodeSelectUpdateAmwayMarketingCopyOrSkip
		node.Content = node_select_update_amway_marketing_copy_or_skip
	}
}

func updateAmwayMarketingCopy(node *Node) {
	if node.ctx == nil {
		return
	}

	a := node.ctx.(*model.Amway)
	switch node.Msg {
	case "1":
		// 返回
		node.curuser.NextHop = _NodeUpdateAmwayType
		node.Content = node_update_amway_type
	case "2":
		// 跳过
		node.curuser.NextHop = _NodeSelectUpdateAmwayFakePortalOrSkip
		node.Content = node_select_update_amway_fake_portal_or_skip
	default:
		a.MarketingCopy = node.Msg
		err := database.UpdatesAmway(a)
		if err != nil {
			log.Error("updateAmwayMarketingCopy.database.UpdatesAmway", err.Error(), map[string]interface{}{"a.MarketingCopy": a.MarketingCopy})
			return
		}
		node.curuser.NextHop = _NodeSelectUpdateAmwayFakePortalOrSkip
		node.Content = node_select_update_amway_fake_portal_or_skip
	}
}

func updateAmwayFakePortal(node *Node) {
	if node.ctx == nil {
		return
	}

	a := node.ctx.(*model.Amway)
	switch node.Msg {
	case "1":
		// 返回
		node.curuser.NextHop = _NodeSelectUpdateAmwayMarketingCopyOrSkip
		node.Content = node_select_update_amway_marketing_copy_or_skip
	case "2":
		// 跳过
		a.Valid = true
		err := database.UpdatesAmway(a)
		if err != nil {
			log.Error("updateAmwayFakePortal.database.UpdatesAmway", err.Error(), map[string]interface{}{"a.FakePortal": a.FakePortal})
			return
		}
		node.curuser.GainEXP(viper.GetUint("exp_publish_amway"))
		node.curuser.NextHop = _NodeDashboard
		node.Content = fmt.Sprintf(node_dashboard, node.curuser.NickName, node.curuser.Level, node.curuser.ID)
	default:
		a.FakePortal = node.Msg
		a.Valid = true
		err := database.UpdatesAmway(a)
		if err != nil {
			log.Error("updateAmwayFakePortal.database.UpdatesAmway", err.Error(), map[string]interface{}{"a.FakePortal": a.FakePortal})
			return
		}
		node.curuser.GainEXP(viper.GetUint("exp_publish_amway"))
		node.curuser.NextHop = _NodeDashboard
		node.Content = fmt.Sprintf(viper.GetString("publish_success")+"\n\n"+node_dashboard, node.curuser.NickName, node.curuser.Level, node.curuser.ID)
	}
}
