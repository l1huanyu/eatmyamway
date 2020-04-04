package httpserver

const (
	_EVENT        = "event"
	_TEXT         = "text"
	_SUBSCRIBE    = "subscribe"
	_UNSUBSCRIBE  = "unsubscribe"
	_RESPONSE_XML = "<xml><ToUserName>%s</ToUserName><FromUserName>%s</FromUserName><CreateTime>%d</CreateTime><MsgType>%s</MsgType><Content>%s</Content></xml>"
)

type (
	requestMsg struct {
		ToUserName   string `xml:"ToUserName" validate:"required"`   // 开发者微信号
		FromUserName string `xml:"FromUserName" validate:"required"` // 发送方账号（一个OpenID）
		CreateTime   int    `xml:"CreateTime" validate:"required"`   // 消息创建时间
		MsgType      string `xml:"MsgType" validate:"required"`      // text
		Content      string `xml:"Content"`                          // 文本消息内容
		MsgID        int64  `xml:"MsgId"`                            // 消息id
		Event        string `xml:"Event"`                            // 事件类型
	}

	responseMsg struct {
		ToUserName   string // 接收方账号（收到的OpenID）
		FromUserName string // 开发者微信号
		CreateTime   int    // 消息创建时间
		MsgType      string // text
		Content      string // 回复的消息内容（可换行）
	}
)
