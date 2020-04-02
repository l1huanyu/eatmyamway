package scheduler

import "github.com/spf13/viper"

const (
	_NodeNull = iota
	NodeSubscribe
	NodeUnsubscribe
)

var gScheduler = map[uint]func(node *SchedulerNode){
	_NodeNull: nil,
}

// SchedulerNode 调度节点
type SchedulerNode struct {
	NextHop uint
	OpenID  string
	Msg     string
	Content string
}

// Schedule 开始调度
func (node *SchedulerNode) Schedule() {
	if f, ok := gScheduler[node.NextHop]; ok && f != nil {
		f(node)
	}

	if len(node.Content) == 0 {
		node.Content = viper.GetString("internal_error")
	}
}
