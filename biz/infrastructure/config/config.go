package config

import (
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf
	ListenOn string
	Mongo    struct {
		URL string
		DB  string
	}
	CacheConf cache.CacheConf
	RedisConf redis.RedisConf
}

func NewConfig() (*Config, error) {
	c := new(Config)
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "etc/config.yaml"
	}
	err := conf.Load(path, c)
	if err != nil {
		return nil, err
	}
	err = c.SetUp()
	if err != nil {
		return nil, err
	}
	return c, nil
}
