package scheduler

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

func queryFocusRelation(node *Node) {
	r, err := database.QueryFocusRelation(node.curuser.ID)
	if err != nil {
		log.Error("queryFocusRelation.database.QueryFocusRelation", err.Error(), map[string]interface{}{"node.curuser.ID": node.curuser.ID})
		return
	}
	node.ctx = r
}

func selectInteractOrQueryAmway(node *Node) {
	option, err := strconv.Atoi(node.Msg)
	if err != nil {
		log.Error("selectInteractOrQueryAmway.strconv.Atoi", err.Error(), map[string]interface{}{"node.Msg": node.Msg})
	}

	r := new(model.Relation)
	switch option {
	case 1:
		if node.ctx != nil {
			r = node.ctx.(*model.Relation)
			// 喜欢
			if r.InteractionType == model.InteractionLike {
				node.Content = node_liked_amway
				return
			}

			r.InteractionType = model.InteractionLike
			err = database.UpdateInteractionRelation(r)
			if err == gorm.ErrRecordNotFound {
				node.Content = viper.GetString("not_found")
			} else if err != nil {
				log.Error("selectInteractOrQueryAmway.database.UpdateInteractionRelation", err.Error(), map[string]interface{}{"r.InteractionType": r.InteractionType})
				return
			}
		}
		node.Content = node_like_amway

	case 2:
		if node.ctx != nil {
			r = node.ctx.(*model.Relation)
			// 不喜欢
			if r.InteractionType == model.InteractionDislike {
				node.Content = node_disliked_amway
				return
			}
			r.InteractionType = model.InteractionDislike
			err = database.UpdateInteractionRelation(r)
			if err == gorm.ErrRecordNotFound {
				node.Content = viper.GetString("not_found")
			} else if err != nil {
				log.Error("selectInteractOrQueryAmway.database.UpdateInteractionRelation", err.Error(), map[string]interface{}{"r.InteractionType": r.InteractionType})
				return
			}
		}
		node.Content = node_dislike_amway
	case 3:
		// 无动作
	case 4:
		// 返回主界面
		node.curuser.NextHop = _NodeDashboard
		node.Content = fmt.Sprintf(node_dashboard, node.curuser.NickName, node.curuser.Level, node.curuser.ID)
		if node.ctx == nil {
			return
		}
		// 取消聚焦
		r.Focus = false
		err = database.UpdateFocusRelation(r)
		if err != nil {
			log.Error("selectInteractOrQueryAmway.database.UpdateFocusRelation", err.Error(), map[string]interface{}{"r.UserID": r.UserID, "r.AmwayID": r.AmwayID})
			return
		}
		return
	default:
		log.Error("selectInteractOrQueryAmway", "invalid option", map[string]interface{}{"option": option})
		return
	}

	// 取消聚焦
	if r != nil {
		r.Focus = false
		err = database.UpdateFocusRelation(r)
		if err != nil {
			log.Error("selectInteractOrQueryAmway.database.UpdateFocusRelation", err.Error(), map[string]interface{}{"r.UserID": r.UserID, "r.AmwayID": r.AmwayID})
			return
		}
	}

	// 查看下一个
	node.ctx = _DashboardQueryAmway
	queryAmwayRand(node)
}
