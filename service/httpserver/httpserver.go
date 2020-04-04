package httpserver

import (
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func Start() {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET(viper.GetString("wechat_interface"), checkSignature)
	e.POST(viper.GetString("wechat_interface"), receiveMessages)

	log.Info("main", "EAT MY AMWAY! ", nil)
	log.Error("Start", e.Start(viper.GetString("httpserver_address")).Error(), nil)
}
