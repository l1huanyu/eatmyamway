package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/l1huanyu/eatmyamway/config"
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/l1huanyu/eatmyamway/service/httpserver"
	"github.com/spf13/pflag"
)

var gCfg = pflag.StringP("config", "c", "", "eatmyamway config file path. ")

func main() {
	fmt.Print("~ 所以暂时将你👀闭了起来...(๑˘ ˘๑)")

	// 初始化配置文件
	pflag.Parse()

	if err := config.Init(*gCfg); err != nil {
		panic(err)
	}

	// 初始化日志模块
	log.Init()
	defer log.Close()

	log.Info("main", "LAST DANCE ~", nil)

	// 打开数据库连接
	database.Open()
	defer database.Close()

	// 前进四...
	httpserver.Start()
}
