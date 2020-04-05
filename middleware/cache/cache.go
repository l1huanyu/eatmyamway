package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
)

var gCache *cache.Cache

func Init() {
	gCache = cache.New(time.Duration(viper.GetInt("cache_default_expiration_minute"))*time.Minute, time.Duration(viper.GetInt("cache_cleanup_interval_minute"))*time.Minute)
}

func Set(key string, value interface{}) {
	gCache.Set(key, value, cache.DefaultExpiration)
}

func Get(key string) (interface{}, bool) {
	return gCache.Get(key)
}
