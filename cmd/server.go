package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/l1huanyu/eatmyamway/config"
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/cache"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/l1huanyu/eatmyamway/service/httpserver"
	"github.com/spf13/pflag"
)

var gCfg = pflag.StringP("config", "c", "", "eatmyamway config file path. ")

func main() {
	fmt.Print("~ 所以暂时将你眼睛闭了起来...(๑˘ ˘๑)")

	c := make(chan os.Signal, 0)
	signal.Notify(c, os.Interrupt, os.Kill)
	go safeExit(c)

	// 初始化配置文件
	pflag.Parse()

	if err := config.Init(*gCfg); err != nil {
		panic(err)
	}

	// 初始化日志模块
	log.Init()
	defer log.Close()

	// 初始化缓存
	cache.Init()

	// 打开数据库连接
	database.Open()
	defer database.Close()

	// 起飞飞飞 ~
	httpserver.Start()
}

// 捕获ctrl c信号，在退出程序前关闭连接
func safeExit(c chan os.Signal) {
	<-c
	log.Close()
	if err := database.Close(); err != nil {
		log.Info("database.Close", err.Error(), nil)
	}
	log.Info("safeExit", "LAST DANCE ~", nil)
	os.Exit(0)
}
