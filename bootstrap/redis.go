package bootstrap

import (
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
)

// SetupRedis connect to redis
func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%s:%s", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
	if err := redis.Redis.Ping(); err != nil {
		logger.ErrorString("Redis", "Connect error", err.Error())
	} else {
		logger.InfoString("Redis", "Connect success 222", "success")
	}
}
