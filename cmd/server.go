package main

import (
	"fmt"

	"github.com/l1huanyu/eatmyamway/config"
	"github.com/l1huanyu/eatmyamway/log"
	"github.com/spf13/pflag"
)

var gCfg = pflag.StringP("config", "c", "", "eatmyamway config file path. ")

func main() {
	pflag.Parse()

	if err := config.Init(*gCfg); err != nil {
		panic(err)
	}

	log.Init()
	defer log.Close()

	log.Info("main", map[string]interface{}{
		"撤硕": "美汁汁儿",
	}, "LAST DANCE")

	fmt.Println("~ 所以暂时将你👀闭了起来...(๑˘ ˘๑)")
}
