package wechat

import (
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/interface/wechat", checkSignature)
	e.POST("/interface/wechat", receiveMessages)

	log.Error("Start", e.Start(viper.GetString("address")).Error(), nil)
}
