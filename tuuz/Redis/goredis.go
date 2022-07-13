package Redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"main.go/config/app_conf"
	"time"
)

var goredis_ctx = context.Background()

var goredis *redis.Client

func init() {
	if !app_conf.Recicon_on {
		return
	}
	options := redis.Options{
		Addr:         app_conf.Redicon_address + ":" + app_conf.Redicon_port,
		Username:     app_conf.Redicon_username,
		Password:     app_conf.Redicon_password, // no password set
		Network:      app_conf.Redicon_proto,
		DialTimeout:  app_conf.Redicon_timeout,
		PoolTimeout:  30,
		ReadTimeout:  app_conf.Redicon_readtimeout,
		WriteTimeout: app_conf.Redicon_writetimeout,
		DB:           0, // use default DB
		MinIdleConns: app_conf.Redicon_MinIdleConn,
	}
	if app_conf.Redicon_poolsize > 0 {
		options.PoolSize = app_conf.Redicon_poolsize
	}
	goredis = redis.NewClient(&options)
	go func() {
		for {
			time.Sleep(3 * time.Second)
			_, err := goredis.Ping(goredis_ctx).Result()
			if err != nil {
				panic("redis_out")
			}
		}
	}()
}
