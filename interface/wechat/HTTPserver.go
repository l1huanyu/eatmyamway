package wechat

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/interface/wechat", checkSignature)
	e.POST("/interface/wechat", receiveMessages)

	e.Start(viper.GetString("address"))
}
