package redis

import (
	"gogin/example/pkg/logging"
	"gogin/example/pkg/setting"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func init() {
	sec, err := setting.Cfg.GetSection("redis")

	if err != nil {
		logging.Error(err)
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr: sec.Key("Host").String(),
		Password: sec.Key("Password").String(),
		DB: sec.Key("DB").MustInt(),
	})
}