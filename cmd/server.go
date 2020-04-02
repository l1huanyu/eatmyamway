package main

import (
	"fmt"

	"github.com/l1huanyu/eatmyamway/config"
	"github.com/l1huanyu/eatmyamway/interface/wechat"
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/l1huanyu/eatmyamway/middleware/database"
	"github.com/spf13/pflag"
)

var gCfg = pflag.StringP("config", "c", "", "eatmyamway config file path. ")

func main() {
	fmt.Print("~ 所以暂时将你👀闭了起来...(๑˘ ˘๑)")

	pflag.Parse()

	if err := config.Init(*gCfg); err != nil {
		panic(err)
	}

	log.Init()
	defer log.Close()

	log.Info("main", "LAST DANCE", map[string]interface{}{
		"撤硕": "美汁汁儿",
	})

	database.Open()
	defer database.Close()

	wechat.Run()
}
