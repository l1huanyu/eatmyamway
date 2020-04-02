package database

import (
	"github.com/jinzhu/gorm"
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

// 数据库连接池
var gDB *gorm.DB

// Open 打开数据库连接并设置连接属性
func Open() {
	gDB, err := gorm.Open(viper.GetString("database_driver"), viper.GetString("database_source"))
	if err != nil {
		panic(err)
	}

	// 设置最大闲置连接
	gDB.DB().SetMaxIdleConns(viper.GetInt("database_max_idle_conns"))
	// 设置最大连接数量
	gDB.DB().SetMaxOpenConns(viper.GetInt("database_max_open_conns"))
	// 设置连接的最大可复用时间
	gDB.DB().SetConnMaxLifetime(viper.GetDuration("database_conn_max_life_time"))

	gDB.AutoMigrate(model.Models()...)
}

// Close 关闭连接池
func Close() error {
	return gDB.Close()
}

// Conn 从连接池中取出一个空闲连接
func Conn() *gorm.DB {
	return gDB
}
