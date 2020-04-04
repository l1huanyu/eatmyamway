package httpserver

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/l1huanyu/eatmyamway/service/scheduler"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var gValidator = validator.New()

func checkSignature(c echo.Context) error {
	signature := c.QueryParam("signature")
	timestamp := c.QueryParam("timestamp")
	nonce := c.QueryParam("nonce")
	token := viper.GetString("access_token")
	echoStr := c.QueryParam("echostr")

	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)
	tmpStr := strings.Join(tmpArr, "")
	tmpBytes := sha1.Sum([]byte(tmpStr))
	tmpStr = string(tmpBytes[:])

	if strings.Compare(signature, tmpStr) != 0 {
		return c.String(http.StatusBadRequest, "Check signature from wechat server failed. ")
	}

	return c.String(http.StatusOK, echoStr)
}

func receiveMessages(c echo.Context) error {
	request := new(requestMsg)
	err := c.Bind(request)
	if err != nil {
		log.Error("receiveMessages.echo.Context.Bind", err.Error(), nil)
		return c.NoContent(http.StatusBadRequest)
	}

	err = gValidator.Struct(request)
	if err != nil {
		log.Error("receiveMessages.gValidator.Struct", err.Error(), nil)
		return c.NoContent(http.StatusBadRequest)
	}

	node := &scheduler.Node{
		OpenID: request.fromUserName,
		Msg:    request.content,
	}

	switch request.msgType {
	case _TEXT:
		if len(request.content) == 0 {
			log.Error("receiveMessages", "len(request.content) == 0", nil)
			return c.NoContent(http.StatusBadRequest)
		}
		// 查找当前用户的下一跳
		node.NextHop, err = database.QueryUserNextHopByOpenID(node.OpenID)
		if err != nil && err == gorm.ErrRecordNotFound {
			node.NextHop = scheduler.NodeSubscribe()
		} else {
			log.Error("database.QueryUserNextHopByOpenID", err.Error(), map[string]interface{}{"node.OpenID": node.OpenID})
		}

	case _EVENT:
		if request.event == _SUBSCRIBE {
			node.NextHop = scheduler.NodeSubscribe()
		} else if request.event == _UNSUBSCRIBE {
			node.NextHop = scheduler.NodeUnsubscribe()
		}
	}

	// 调度执行下一跳动作
	node.Schedule()

	response := &responseMsg{
		toUserName:   request.fromUserName,
		fromUserName: request.toUserName,
		createTime:   int(time.Now().Unix()),
		msgType:      _TEXT,
		content:      node.Content,
	}

	return c.String(http.StatusOK, fmt.Sprintf(_RESPONSE_XML, response.toUserName, response.fromUserName, response.createTime, response.msgType, response.content))
}
