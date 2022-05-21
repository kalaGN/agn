package bootstrap

import (
	"agn/pkg/config"
	"agn/pkg/redis"
	"fmt"
)

// SetupRedis 初始化 Redis
func SetupRedis() {
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)

	//
	redis.ConnecLimittRedis(
		fmt.Sprintf("%v:%v", config.GetString("limitredis.host"), config.GetString("limitredis.port")),
		config.GetString("limitredis.username"),
		config.GetString("limitredis.password"),
		config.GetInt("limitredis.database"),
	)

}
