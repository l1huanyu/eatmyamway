package wechat

const (
	_EVENT        = "event"
	_TEXT         = "text"
	_SUBSCRIBE    = "subscribe"
	_UNSUBSCRIBE  = "unsubscribe"
	_RESPONSE_XML = "<xml><ToUserName>%s</ToUserName><FromUserName>%s</FromUserName><CreateTime>%d</CreateTime><MsgType>%s</MsgType><Content>%s</Content></xml>"
)

type (
	requestMsg struct {
		toUserName   string `xml:"ToUserName" validate:"required"`   // 开发者微信号
		fromUserName string `xml:"FromUserName" validate:"required"` // 发送方账号（一个OpenID）
		createTime   int    `xml:"CreateTime" validate:"required"`   // 消息创建时间
		msgType      string `xml:"MsgType" validate:"required"`      // text
		content      string `xml:"Content"`                          // 文本消息内容
		msgID        int64  `xml:"MsgId"`                            // 消息id
		event        string `xml:"Event"`                            // 事件类型
	}

	responseMsg struct {
		toUserName   string // 接收方账号（收到的OpenID）
		fromUserName string // 开发者微信号
		createTime   int    // 消息创建时间
		msgType      string // text
		content      string // 回复的消息内容（可换行）
	}
)
