package wechat

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/l1huanyu/eatmyamway/interface/ui"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
	content := ""

	request := new(requestMsg)
	err := c.Bind(request)
	if err != nil {
		logrus.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err = gValidator.Struct(request)
	if err != nil {
		logrus.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	switch request.msgType {
	case _TEXT:
		if len(request.content) == 0 {
			return c.NoContent(http.StatusBadRequest)
		}
		// 开始调度
		content = ui.Schedule(request.fromUserName, request.content)

	case _EVENT:
		if request.event == _SUBSCRIBE {
			content = ui.Prologue(request.fromUserName)
		} else {
			if request.event == _UNSUBSCRIBE {
				ui.Realese(request.fromUserName)
			}
			return c.NoContent(http.StatusOK)
		}

	default:
		content = ui.NotSuport()
	}

	response := &responseMsg{
		toUserName:   request.fromUserName,
		fromUserName: request.toUserName,
		createTime:   int(time.Now().Unix()),
		msgType:      _TEXT,
		content:      content,
	}

	return c.String(http.StatusOK, fmt.Sprintf(_RESPONSE_XML, response.toUserName, response.fromUserName, response.createTime, response.msgType, response.content))
}
