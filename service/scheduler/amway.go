package scheduler

import (
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
		node.Content = node_dashboard
		node.ctx = nil
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
	a := node.ctx.(*model.Amway)
	switch node.Msg {
	case "1":
		// 返回
		node.curuser.NextHop = _NodeCreateAndUpdateAmwayName
		node.Content = node_create_and_update_amway_name
		node.ctx = nil
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
	a := node.ctx.(*model.Amway)
	switch node.Msg {
	case "1":
		// 返回
		node.curuser.NextHop = _NodeUpdateAmwayType
		node.Content = node_update_amway_type
		node.ctx = nil
	case "2":
		// 跳过
		node.curuser.NextHop = _NodeSelectUpdateAmwayFakePortalOrSkip
		node.Content = node_select_update_amway_fake_portal_or_skip
		node.ctx = nil
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
	a := node.ctx.(*model.Amway)
	switch node.Msg {
	case "1":
		// 返回
		node.curuser.NextHop = _NodeSelectUpdateAmwayMarketingCopyOrSkip
		node.Content = node_select_update_amway_marketing_copy_or_skip
		node.ctx = nil
	case "2":
		// 跳过
		node.curuser.NextHop = _NodeDashboard
		node.Content = node_dashboard
		node.ctx = nil
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
		node.Content = node_dashboard
	}
}
