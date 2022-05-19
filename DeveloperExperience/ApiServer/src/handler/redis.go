package handler

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var ClientRedis *redis.Client

func InitRedisClient() {
	options := redis.Options{
		Network:            ConfigContent.Redis.Network,
		Addr:               fmt.Sprintf("%s:%s", ConfigContent.Redis.Host, ConfigContent.Redis.Port),
		Dialer:             nil,
		OnConnect:          nil,
		Password:           ConfigContent.Redis.Password,
		DB:                 ConfigContent.Redis.Db,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	ClientRedis = redis.NewClient(&options)
}

func RedisSetKey(Key string, Value string) {
	ClientRedis.Set(Key, Value, time.Second*time.Duration(ConfigContent.Cache.Second))
}

func RedisGetKey(Key string) string {
	Result := ClientRedis.Get(Key)
	Value := Result.Val()
	return Value
}
