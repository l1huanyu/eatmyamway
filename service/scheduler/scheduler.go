package scheduler

import (
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

type (
	nodeAction func(node *Node)
	nodeFunc   []nodeAction

	// Node 调度节点
	Node struct {
		NextHop uint
		OpenID  string
		Msg     string
		Content string
		curuser *model.User // 当前用户
		ctx     interface{} // 缓存调度节点的上下文
	}
)

const (
	_NodeNull                            = iota // 空节点，该节点代表服务内部出现错误，下一跳不变，期望用户再试一次
	_NodeUnsubscribe                            // 取消订阅
	_NodeSubscribe                              // 订阅，下一跳：选择更新昵称或返回主界面
	_NodeSelectUpdateNickNameOrDashboard        // 选择更新昵称或返回主界面，下一跳：更新昵称/返回主界面
	_NodeUpdateNickName                         // 更新昵称，下一跳：主界面
	_NodeDashboard                              // 主界面，下一跳：选择互动或查找安利/创建新的安利/个人信息界面
	_NodeSelectInteractOrQueryAmway             // 选择互动或查找安利，下一跳：喜欢/不喜欢/查看下一个/返回主界面
	_NodeCreateAmway                            // 创建新的安利
	_NodePersonalInterface                      // 个人界面，下一跳：
	_NodeLikeAmway
	_NodeDislikeAmway
)

// 目前版本并没有对全局调度器的动态修改，故不用加锁
var gScheduler = map[uint]nodeFunc{
	_NodeNull:                            nodeFunc{},
	_NodeUnsubscribe:                     nodeFunc{},
	_NodeSubscribe:                       nodeFunc{queryOrCreateUser, updateUserNextHop},
	_NodeSelectUpdateNickNameOrDashboard: nodeFunc{queryUser, selectUpdateNickNameOrDashboard, updateUserNextHop},
	_NodeUpdateNickName:                  nodeFunc{queryUser, updateUserNickName, updateUserNextHop},
	_NodeDashboard:                       nodeFunc{queryUser, dashboard, queryAmwayWithMaxHP, createAmway, personalInterface, updateUserNextHop},
	_NodeSelectInteractOrQueryAmway:      nodeFunc{queryUser, updateUserNextHop},
}

// Schedule 开始调度
func (node *Node) Schedule() {
	if funcs, ok := gScheduler[node.NextHop]; ok {
		for _, f := range funcs {
			f(node)
		}
	}

	if len(node.Content) == 0 {
		node.Content = viper.GetString("internal_error")
	}
}

// NodeSubscribe return _NodeSubscribe
func NodeSubscribe() uint {
	return _NodeSubscribe
}

// NodeUnsubscribe return _NodeUnsubscribe
func NodeUnsubscribe() uint {
	return _NodeUnsubscribe
}
