package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/l1huanyu/eatmyamway/model"
	"github.com/spf13/viper"
)

var gDB *gorm.DB

func Open() {
	gDB, err := gorm.Open(viper.GetString("database_driver"), viper.GetString("database_source"))
	if err != nil {
		panic(err)
	}

	gDB.DB().SetMaxIdleConns(viper.GetInt("database_max_idle_conns"))
	gDB.DB().SetMaxOpenConns(viper.GetInt("database_max_open_conns"))
	gDB.DB().SetConnMaxLifetime(viper.GetDuration("database_conn_max_life_time"))

	gDB.AutoMigrate(model.Models()...)
}

func Close() error {
	return gDB.Close()
}

func Conn() *gorm.DB {
	return gDB
}
