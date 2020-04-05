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
	_NodeNull                                 = iota // 空节点，该节点代表服务内部出现错误，下一跳不变，期望用户再试一次
	_NodeUnsubscribe                                 // 取消订阅
	_NodeSubscribe                                   // 订阅，下一跳：选择更新昵称或返回主界面
	_NodeSelectUpdateNickNameOrDashboard             // 选择更新昵称或返回主界面，下一跳：更新昵称/返回主界面
	_NodeUpdateNickName                              // 更新昵称，下一跳：主界面
	_NodeDashboard                                   // 主界面，下一跳：选择互动或查找安利/创建新的安利/个人信息界面
	_NodeSelectInteractOrQueryAmway                  // 选择互动或查找安利，下一跳：喜欢/不喜欢/查看下一个/返回主界面
	_NodeCreateAndUpdateAmwayName                    // 创建并输入安利主体名字
	_NodeUpdateAmwayType                             // 输入安利类型
	_NodeSelectUpdateAmwayMarketingCopyOrSkip        // 输入安利理由/跳过
	_NodeSelectUpdateAmwayFakePortalOrSkip           // 输入【伪】传送门/跳过，下一跳返回主界面
	_NodePersonalInterface                           // 个人界面，下一跳：
)

const (
	amway_information                               = "招待不周~ 来自用户\n【%s】LV.%d UID:%d\n的安利！\n=͟͟͞͞ ٩( ๑╹ ꇴ╹)۶\n----------------------------------------\n名字：%s\n类型：%s\n生命值：%d\n安利理由：%s\n【伪】传送门：%s\n----------------------------------------"
	node_subscribe                                  = "久等了~\n感谢订阅【食我安利】！\no(*￣▽￣*)ブ\n\n鲁迅说过：愿乐于分享的人都摆脱冷气，只是向人安利，不必听冷漠不屑者流的话。能安利的不吝啬，能接受的不拒绝。有一分热爱，分享一分喜欢，就令萤火一般，也可以在黑暗里发一点光，不必等候炬火。此后如竟没有炬火：我便是唯一的光。\nPS：和别人分享自己喜欢的事物，一份的快乐就能变成多份哟~\n\n"
	node_select_update_nick_name_or_dashboard       = "那个~ 现在的昵称是：\n【%s】\n要换个昵称吗？\n(๑• . •๑)\n1：好鸭...\n2：不了..."
	node_update_nick_name                           = "那么~ 请输入新的昵称吧，不要超过8个字符哦~ \n(๑✦ˑ̫✦)"
	node_dashboard                                  = "所以~ \n【%s】LV.%d UID:%d大大\n接下来做什么呢？\n(๑• . •๑)\n1：想被安利...\n2：我想安利...\n3：查看我的信息"
	node_select_interact_or_query_amway             = "呐~ 喜欢这份安利吗？\n⁄(⁄ ⁄•⁄ω⁄•⁄ ⁄)⁄\n1：喜欢...\n2：你是个好安利，可是...\n3：康康其他的...\n4：返回主界面"
	node_like_amway                                 = "哇~ 你能喜欢我真的很开心~ \n(*/ω＼*)\n再康康其他的吧~\n\n"
	node_liked_amway                                = "可是~ 你已经喜欢过它了~\n(´◔ω◔)"
	node_dislike_amway                              = "吓~ Σ( ° △ °|||) 没关系，再康康其他的吧~\n\n"
	node_disliked_amway                             = "对不起...没想到你这么讨厌...这已经不是第一次了...(。_。)"
	node_create_and_update_amway_name               = "咳咳~ 想安利什么呢？不要超过20个字符哦~ \n(๑✦ˑ̫✦)\n1：返回"
	node_update_amway_type                          = "嗯嗯~ 这是番剧、游戏、视频、日剧美剧还是...呢？\n(｡･ω･｡)\n1：返回"
	node_select_update_amway_marketing_copy_or_skip = "好的知道了~ 宁为什么想安利它呢？\n(っ•̀ω•́)っ✎⁾⁾\n1：返回\n2：跳过"
	node_select_update_amway_fake_portal_or_skip    = "这样啊~ 能不能给大家一个神奇链接呢~\no(*≧▽≦)ツ\n1：返回\n2：跳过"
)

// 目前版本并没有对全局调度器的动态修改，故不用加锁
var gScheduler = map[uint]nodeFunc{
	_NodeNull:                                 nodeFunc{},
	_NodeUnsubscribe:                          nodeFunc{},
	_NodeSubscribe:                            nodeFunc{queryOrCreateUser, updateUserNextHop},
	_NodeSelectUpdateNickNameOrDashboard:      nodeFunc{queryUser, selectUpdateNickNameOrDashboard, updateUserNextHop},
	_NodeUpdateNickName:                       nodeFunc{queryUser, updateUserNickName, updateUserNextHop},
	_NodeDashboard:                            nodeFunc{queryUser, dashboard, queryAmwayRand, createAmway, personalInterface, updateUserNextHop},
	_NodeSelectInteractOrQueryAmway:           nodeFunc{queryUser, queryFocusRelation, selectInteractOrQueryAmway, updatesUser},
	_NodeCreateAndUpdateAmwayName:             nodeFunc{queryUser, queryInvalidAmway, updateAmwayName, updateUserNextHop},
	_NodeUpdateAmwayType:                      nodeFunc{queryUser, queryInvalidAmway, updateAmwayType, updateUserNextHop},
	_NodeSelectUpdateAmwayMarketingCopyOrSkip: nodeFunc{queryUser, queryInvalidAmway, updateAmwayMarketingCopy, updateUserNextHop},
	_NodeSelectUpdateAmwayFakePortalOrSkip:    nodeFunc{queryUser, queryInvalidAmway, updateAmwayFakePortal, updatesUser},
	_NodePersonalInterface:                    nodeFunc{},
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
